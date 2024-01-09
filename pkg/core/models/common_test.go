package models

import "github.com/DevDrift/investment-game/pkg/core"

var (
	testId      = "test"
	testName    = "tName"
	testNewName = "newName"
	testEmail   = "tEmail"
	inputMoney  = 100.0
	testPlayer1 = &core.Player{
		Id:    testId,
		Name:  testName,
		Email: testEmail,
	}
	testBalance1 = &core.Balance{
		Id:      testId,
		Account: inputMoney,
	}
	testPortfolio = &core.Portfolio{
		Id: testId,
	}
	auctionTestStep = 0.10
)
