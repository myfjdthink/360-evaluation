package response

import "360-evaluation/models"

type UserResponse struct {
	User models.User `json:"user"`
}
