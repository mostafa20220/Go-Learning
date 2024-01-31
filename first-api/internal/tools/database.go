package tools

import (
	"github.com/sirupsen/logrus"
)

// database collections
type LoginDetails struct{
	Username string
	Token string
}
type CoinsDetails struct{
	Username string
	Coins int64
}

type DatabaseInterface interface{
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoinsDetails(username string) *CoinsDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){

	var database DatabaseInterface = &mockDB{}
	var err = database.SetupDatabase()

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &database, err
}
