package services

import (
	"errors"
	"time"
	"example.com/poc/config"
	"example.com/poc/models"
	"example.com/poc/repositories"
)

func SubmitCreditCardUserRole(creditCard *models.CreditCard) error {
	// Check if the card has already been processed
	exists := repositories.CheckCreditCardExists(creditCard.CardNumber)

	if exists {
		return errors.New("card already processed")
	}

	// Check if the issuing country is in the banned list
	isBanned := IsValidCreditCard(creditCard.IssuingCountry)

	if isBanned {
		return errors.New("card issuing country is banned")
	}

	// Store the card in the repository
	var err error
	err = repositories.SubmitCreditCardUserRole(creditCard)
	if err != nil {
		return err
	}

	return nil
}

func SubmitCreditCardAdminRole(creditCard []*models.CreditCard) (bool , error) {
	// Initialize a temporary array to store the unprocessed credit cards
	var tempArray []*models.CreditCard

	for _, card := range creditCard {
		exists := repositories.CheckCreditCardExists(card.CardNumber)
		if exists {
			// Skip the card if it has already been processed
			continue
		}

		isBanned := IsValidCreditCard(card.IssuingCountry)
		if isBanned {
			// Skip the card if the issuing country is banned
			continue
		}

		// Add the card to the temporary array
		tempArray = append(tempArray, card)
	}

	if len(tempArray) == 0{
		return false , errors.New("no new records to insert")
	}

	// Determine the number of arrays needed for bulk processing
	arrayCount := len(tempArray) / 5
	if len(tempArray)%5 != 0 {
		arrayCount++
	}

	// Perform bulk processing
	for i := 0; i < arrayCount; i++ {
		startIndex := i * 5
		endIndex := (i + 1) * 5
		if endIndex > len(tempArray) {
			endIndex = len(tempArray)
		}

		// Pass each batch of 5 records to SubmitCreditCardAdminRole function
		err := SubmitCreditCardAdminRoleBatch(tempArray[startIndex:endIndex])
		if err != nil {
			return false , err
		}

		// Wait for 30 seconds before processing the next batch
		if i != arrayCount-1 {
			time.Sleep(30 * time.Second)
		}
	}

	return true ,nil
}

func SubmitCreditCardAdminRoleBatch(batch []*models.CreditCard) error {
	// Store the cards in the repository
	err := repositories.SubmitCreditCardAdminRole(batch)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCreditCards() ([]models.CreditCard, error) {
	return repositories.GetAllCreditCards()
}

func GetCreditCard(id string) (models.CreditCard, error) {
	return repositories.GetCreditCardByID(id)
}

func IsValidCreditCard(issuingCountry string) bool {

	if isCardBanned(issuingCountry) {
		return true
	}
	return false
}

func isCardBanned(country string) bool {
	config := config.GetConfig()
	bannedCountries := config.BannedCountries
	for _, bannedCountry := range bannedCountries {
		if bannedCountry == country {
			return true
		}
	}

	return false
}