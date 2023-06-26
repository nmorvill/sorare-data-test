package main

import (
	"errors"
)

// mock implementation for testing purposes

var ManagersDb map[string]*Manager
var CardValuesDb map[string]float64

// returns card value from card tokenId
func GetCardValue(tokenId string) float64 {
	if val, ok := CardValuesDb[tokenId]; ok {
		return val
	}
	return 0
}

// returns a manager based on his id
func GetManager(sorareAddress string) (*Manager, error) {
	if val, ok := ManagersDb[sorareAddress]; ok {
		return val, nil
	}
	return &Manager{}, errors.New("Manager not found in DB")
}
