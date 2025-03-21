package models

type Password struct {
	NewPassword    string `json:"newPassword"`
	ActualPassword string `json:"actualPassword"`
}
