package request

import "time"

type CredentialChangeRequest struct {
	OldCredential string `json:"old_credential" binding:"required"`
	NewCredential string `json:"new_credential" binding:"required"`
	Time          time.Time
}
