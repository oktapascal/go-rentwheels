package exceptions

import (
	"encoding/json"
	"go-rental/config"
	"go-rental/response"
	"net/http"
)

type DuplicateError struct {
	Error string
}

// NewDuplicateError creates a new DuplicateError with the given error message.
//
// error: The error message to be associated with the DuplicateError.
//
// Returns a DuplicateError with the given error message.
func NewDuplicateError(error string) DuplicateError {
	return DuplicateError{
		Error: error,
	}
}

// DuplicateHandler handles HTTP 400 Bad Request responses for duplicate errors.
// It writes a JSON response with the appropriate status code and error details.
// If an error occurs while encoding the response, it logs the error.
//
// Parameters:
// - writer: The http.ResponseWriter to write the response to.
// - err: The error interface containing the details of the error.
func DuplicateHandler(writer http.ResponseWriter, err any) {
	// Create a logger for error logging
	log := config.CreateLoggers(nil)

	// Set the content type of the response to JSON
	writer.Header().Set("Content-Type", "application/json")

	// Set the status code of the response to Bad Request
	writer.WriteHeader(http.StatusBadRequest)

	// Create an error response with the status code and error details
	errorResponse := response.ErrorResponse{

		Code:   http.StatusBadRequest,                  // Set the status code to Bad Request
		Status: http.StatusText(http.StatusBadRequest), // Set the status text to the corresponding HTTP status text
		Errors: err,                                    // Set the error details to the provided error
	}

	// Encode the error response into JSON
	encoder := json.NewEncoder(writer)

	// Check if there was an error encoding the response
	if errEncoder := encoder.Encode(errorResponse); errEncoder != nil {
		// Log the error if there was an error encoding the response
		log.Error(errEncoder)
	}
}
