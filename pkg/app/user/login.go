package user

import (
	"bs-2018/hpfs-client/errors"
	"context"
	"github.com/eleztian/apipm/pkg/app"
	"github.com/eleztian/apipm/pkg/app/db/mongodb"
)

func Login(ctx context.Context, account string, password string) (*app.Response, error) {
	if account == "" || password == "" {
		return nil, errors.New("account or password is empty")
	}

	mDb := mongodb.NewCollection("userInfo")

	return &app.Response{
		Code: 0,
		Data: "tkpppppppppppppppppppp",
	}, nil
}

func Logout(ctx context.Context) error {
	return nil
}


func GenJWT() (string, error) {

	return "", nil
}