package handler

import (
	"encoding/json"
	"net/http"

	"example.com/poc/models"
	"example.com/poc/services"
	"github.com/gorilla/mux"
)

type ServerResponse struct {
	Response   bool        `json:"response"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}

func SubmitCreditCardUserRole(w http.ResponseWriter, r *http.Request) {
	var creditCard models.CreditCard

	// Decode the request body into the credit card object
	err := json.NewDecoder(r.Body).Decode(&creditCard)
	if err != nil {
		sendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// Call the service to validate and store the credit card
	err = services.SubmitCreditCardUserRole(&creditCard)
	if err != nil {
		sendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	response := ServerResponse{
		Response:   true,
		StatusCode: http.StatusCreated,
		Data:       "Credit card submitted successfully",
	}

	sendResponse(w, response)
}

func SubmitCreditCardAdminRole(w http.ResponseWriter, r *http.Request) {
	var creditCards []*models.CreditCard

	// Decode the request body into an array of credit card objects
	err := json.NewDecoder(r.Body).Decode(&creditCards)
	if err != nil {
		sendErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	// Call the service to validate and store the credit cards in batches
	success, err := services.SubmitCreditCardAdminRole(creditCards)
	if err != nil {
		sendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if success {
		response := ServerResponse{
			Response:   true,
			StatusCode: http.StatusCreated,
			Data:       "Credit cards submitted successfully",
		}
		sendResponse(w, response)
	} else {
		response := ServerResponse{
			Response:   false,
			StatusCode: http.StatusBadRequest,
			Data:       "No new records to insert",
		}
		sendResponse(w, response)
	}
}

func GetAllCreditCards(w http.ResponseWriter, r *http.Request) {
	creditCards, err := services.GetAllCreditCards()

	if err != nil {
		sendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	response := ServerResponse{
		Response:   true,
		StatusCode: http.StatusCreated,
		Data:       creditCards,
	}

	sendResponse(w, response)
}

func GetCreditCard(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	creditCardID := params["id"]

	creditCard, err := services.GetCreditCard(creditCardID)
	if err != nil {
		sendErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	response := ServerResponse{
		Response:   true,
		StatusCode: http.StatusCreated,
		Data:       creditCard,
	}
	sendResponse(w, response)

}

func sendErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	response := ServerResponse{
		Response:   false,
		StatusCode: statusCode,
		Data:       err.Error(),
	}

	sendResponse(w, response)
}

func sendResponse(w http.ResponseWriter, response ServerResponse) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(jsonResponse)
}