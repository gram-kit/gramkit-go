package params

import "github.com/gram-kit/gramkit-go/models"

type SetPassportDataErrors struct {
	UserID int64                         `json:"user_id"`
	Errors []models.PassportElementError `json:"errors"`
}
