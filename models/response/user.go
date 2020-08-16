package response

import (
	"time"
)

type UserResponse struct {
	Name string `json:"name"`
}

type LoginByPasswordRes struct {
	Token  string        `json:"token"`
	Expire time.Duration `json:"expire"`
}
