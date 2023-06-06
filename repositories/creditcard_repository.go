package repositories

import (
	"context"
	"fmt"
	"log"

	"example.com/poc/models"
	"example.com/poc/utils"
	"github.com/go-kivik/kivik/v4"
)

var db *kivik.DB

func init() {
	// Connect to CouchDB on application startup
	var err error
	db, err = utils.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to CouchDB: %v", err)
	}
	db = utils.GetDB()
}

func SubmitCreditCardUserRole(creditCard *models.CreditCard) error {
	// Store the credit card in CouchDB
	ctx := context.Background()
	rev, err := db.Put(ctx, creditCard.CardNumber, creditCard, nil)
	if err != nil {
		return err
	}
	log.Println("document inserted at id", rev)
	return nil
}

func SubmitCreditCardAdminRole(creditCard []*models.CreditCard) error {
	// Store the credit cards in CouchDB
	ctx := context.Background()
	batchSize := 5

	for i := 0; i < len(creditCard); i += batchSize {
		endIndex := i + batchSize
		if endIndex > len(creditCard) {
			endIndex = len(creditCard)
		}

		batch := creditCard[i:endIndex]

		for _, card := range batch {
			rev, err := db.Put(ctx, card.CardNumber, card, nil)
			if err != nil {
				return err
			}
			log.Println("document inserted at id", rev)
		}

		log.Printf("Batch of %d cards inserted", len(batch))
	}

	return nil
}

func CheckCreditCardExists(id string) bool {
	// Create a selector to match the credit card with the given ID
	options := kivik.Options{
		"include_docs": true,
	}
	row := db.Get(context.Background(), id, options)

	doc := models.CreditCard{}
	if err := row.ScanDoc(&doc); err != nil {
		// An error occurred while scanning the document
		return false
	}

	if doc.CardNumber == id {
		return true
	}

	return true
}


func GetAllCreditCards() ([]models.CreditCard, error) {
	// Retrieve all credit cards from CouchDB
	result := []models.CreditCard{}
	options := map[string]interface{}{
		"include_docs": true,
	}
	rows := db.AllDocs(context.Background(), options)
	defer rows.Close()
	for rows.Next() {
		var doc models.CreditCard
		if err := rows.ScanDoc(&doc); err != nil {
			return nil, err
		}
		result = append(result, doc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func GetCreditCardByID(id string) (models.CreditCard, error) {
	// Retrieve the credit card from CouchDB by ID
	options := kivik.Options{
		"include_docs": true,
	}
	row := db.Get(context.Background(), id, options)
	doc := models.CreditCard{}
	if row.Err() != nil {
		return doc, fmt.Errorf("failed to retrieve credit card: %w", row.Err())
	}
	err := row.ScanDoc(&doc) //  row.ScanDoc to scan the document into the models.CreditCard struct.
	if err != nil {
		return doc, fmt.Errorf("failed to scan credit card document: %w", err)
	}
	return doc, nil
}
