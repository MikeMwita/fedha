package validator

import (
	"errors"
	"github.com/MikeMwita/fedha.git/services/app-payment/models"
)

func ValidateSTKPushRequestBody(req *models.STKPushRequestBody) error {
	if req.BusinessShortCode == "" {
		return errors.New("business shortcode is required")
	}

	if req.Password == "" {
		return errors.New("password is required")
	}

	if req.Timestamp == "" {
		return errors.New("timestamp is required")
	}

	if req.TransactionType == "" {
		return errors.New("transaction type is required")
	}

	if req.Amount == "" {
		return errors.New("amount is required")
	}

	if req.PartyA == "" {
		return errors.New("party A is required")
	}

	if req.PartyB == "" {
		return errors.New("party B is required")
	}

	if req.PhoneNumber == "" {
		return errors.New("phone number is required")
	}

	if req.CallbackURL == "" {
		return errors.New("callback URL is required")
	}

	if req.AccountReference == "" {
		return errors.New("account reference is required")
	}

	if req.TransactionDesc == "" {
		return errors.New("transaction description is required")
	}

	return nil
}
