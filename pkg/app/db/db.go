package db

import "github.com/eleztian/apipm/pkg/app/db/mongodb"

func DBUserInfo() *mongodb.DB {
	return mongodb.NewCollection(COL_USER_INFO)
}