package tools

import (
	"time"
)


type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"mostafa": {
		Username: "mostafa",
		Token: "id1",
	},
	"eslam": {
		Username: "eslam",
		Token: "id2",
	},
}

var mockCoinsDetails = map[string]CoinsDetails{
	"mostafa": {
		Username: "mostafa",
		Coins: 1_000,
	},
	"eslam":{
		Username: "eslam",
		Coins: 100,
	},
}

func (m mockDB) GetUserLoginDetails(username string) *LoginDetails{
	
	// simulate db call
	time.Sleep(time.Second * 1)

	res, ok := mockLoginDetails[username]
	if !ok {

		return nil
	}
	
	return &res
}

func (m mockDB) GetUserCoinsDetails(username string) *CoinsDetails{
	// simulate db call
	time.Sleep(time.Second * 1)

	res, ok := mockCoinsDetails[username]
	if !ok {
		return nil
	}

	return &res
}

func (m mockDB) SetupDatabase() error {
	return nil
}