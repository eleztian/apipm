package db

import "github.com/eleztian/apipm/pkg/app/db/mongodb"

// MongoDB collection names
const (
	COL_USER_INFO = "user_info"
)

type UserInfo struct {
	Email    string
	Password string
}

