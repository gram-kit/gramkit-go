package models

import "github.com/gram-kit/gramkit-go/enums"

type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        enums.PassportElementType `json:"type"`
	Data        string                    `json:"data,omitempty"`
	PhoneNumber string                    `json:"phone_number,omitempty"`
	Email       string                    `json:"email,omitempty"`
	Files       []PassportFile            `json:"files,omitempty"`
	FrontSide   PassportFile              `json:"front_side,omitempty"`
	ReverseSide PassportFile              `json:"reverse_side,omitempty"`
	Selfie      PassportFile              `json:"selfie,omitempty"`
	Translation PassportFile              `json:"translation,omitempty"`
	Hash        string                    `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportElementError interface {
	isPassportElementError()
}

func (PassportElementErrorDataField) isPassportElementError()        {}
func (PassportElementErrorFrontSide) isPassportElementError()        {}
func (PassportElementErrorReverseSide) isPassportElementError()      {}
func (PassportElementErrorSelfie) isPassportElementError()           {}
func (PassportElementErrorFile) isPassportElementError()             {}
func (PassportElementErrorFiles) isPassportElementError()            {}
func (PassportElementErrorTranslationFile) isPassportElementError()  {}
func (PassportElementErrorTranslationFiles) isPassportElementError() {}
func (PassportElementErrorUnspecified) isPassportElementError()      {}

type PassportElementErrorDataField struct {
	Source    string                    `json:"source"`
	Type      enums.PassportElementType `json:"type"`
	FieldName string                    `json:"field_name"`
	DataHash  string                    `json:"data_hash"`
	Message   string                    `json:"message"`
}

type PassportElementErrorFrontSide struct {
	Source   string                    `json:"source"`
	Type     enums.PassportElementType `json:"type"`
	FileHash string                    `json:"file_hash"`
	Message  string                    `json:"message"`
}

type PassportElementErrorReverseSide struct {
	Source   string                    `json:"source"`
	Type     enums.PassportElementType `json:"type"`
	FileHash string                    `json:"file_hash"`
	Message  string                    `json:"message"`
}

type PassportElementErrorSelfie struct {
	Source   string                    `json:"source"`
	Type     enums.PassportElementType `json:"type"`
	FileHash string                    `json:"file_hash"`
	Message  string                    `json:"message"`
}

type PassportElementErrorFile struct {
	Source   string                    `json:"source"`
	Type     enums.PassportElementType `json:"type"`
	FileHash string                    `json:"file_hash"`
	Message  string                    `json:"message"`
}

type PassportElementErrorFiles struct {
	Source     string                    `json:"source"`
	Type       enums.PassportElementType `json:"type"`
	FileHashes []string                  `json:"file_hashes"`
	Message    string                    `json:"message"`
}

type PassportElementErrorTranslationFile struct {
	Source   string                    `json:"source"`
	Type     enums.PassportElementType `json:"type"`
	FileHash string                    `json:"file_hash"`
	Message  string                    `json:"message"`
}

type PassportElementErrorTranslationFiles struct {
	Source     string                    `json:"source"`
	Type       enums.PassportElementType `json:"type"`
	FileHashes string                    `json:"file_hashes"`
	Message    string                    `json:"message"`
}

type PassportElementErrorUnspecified struct {
	Source      string                    `json:"source"`
	Type        enums.PassportElementType `json:"type"`
	ElementHash string                    `json:"element_hash"`
	Message     string                    `json:"message"`
}
