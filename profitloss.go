package main

import (
	"fmt"
	"math"
)

type ProfitLossCard struct {
	TokenId       string  //id carte
	SorareAddress string  //addr manager
	Cost          float64 //cout d'acquisition
	SoldFor       float64 //valeur de vente
	Realized      float64 //profit/perte réalisée
	Unrealized    float64 //profit/perte non réalisée
}

type Trade struct {
	AddressManagerA string   //address of manager A
	AddressManagerB string   //address of manager B
	EthManagerA     float64  //eth sent by manager A
	EthManagerB     float64  //eth sent by manager B
	TokensManagerA  []string //ids of cards sent by manager A
	TokensManagerB  []string //ids of card sent by manager B
}

type Manager struct {
	SorareAddress   string                    //address of manager
	ProfitLossCards map[string]ProfitLossCard //profit & losses of  based on tokenid
}

func MakeTrade(trade Trade) {
	managerA, errA := GetManager(trade.AddressManagerA)
	managerB, errB := GetManager(trade.AddressManagerB)

	if len(trade.TokensManagerA) == 0 { // Manager A only send ETH
		var totalBuyCost float64
		for _, t := range trade.TokensManagerB {
			totalBuyCost += managerB.ProfitLossCards[t].Cost
		}
		for _, t := range trade.TokensManagerB {
			cardPrice := ((trade.EthManagerA - trade.EthManagerB) * managerB.ProfitLossCards[t].Cost) / totalBuyCost
			if errA == nil {
				managerA.BuyCardForEth(t, cardPrice)
			}
			if errB == nil {
				managerB.sellCardForEth(t, cardPrice)
			}
		}
	} else if len(trade.TokensManagerB) == 0 { // Manager B only send ETH
		var totalBuyCost float64
		for _, t := range trade.TokensManagerA {
			totalBuyCost += managerA.ProfitLossCards[t].Cost
		}
		for _, t := range trade.TokensManagerA {
			cardPrice := ((trade.EthManagerB - trade.EthManagerA) * managerA.ProfitLossCards[t].Cost) / totalBuyCost
			if errB == nil {
				managerB.BuyCardForEth(t, cardPrice)
			}
			if errA == nil {
				managerA.sellCardForEth(t, cardPrice)
			}
		}
	} else {
		totalCostSentByManagerA, totalCurrentValueSentByManagerA := trade.EthManagerA, trade.EthManagerA
		for _, t := range trade.TokensManagerA {
			if errA == nil {
				totalCostSentByManagerA += managerA.ProfitLossCards[t].Cost
			}
			totalCurrentValueSentByManagerA += GetCardValue(t)
		}
		totalCostSentByManagerB, totalCurrentValueSentByManagerB := trade.EthManagerB, trade.EthManagerB
		for _, t := range trade.TokensManagerB {
			if errB == nil {
				totalCostSentByManagerB += managerB.ProfitLossCards[t].Cost
			}
			totalCurrentValueSentByManagerB += GetCardValue(t)
		}

		fmt.Println(totalCostSentByManagerA, totalCostSentByManagerB, totalCurrentValueSentByManagerA, totalCurrentValueSentByManagerB)

		for _, t := range trade.TokensManagerA {
			if errA == nil {
				managerA.sellCardForCardDeal(t)
			}
			if errB == nil {
				managerB.BuyCardForEth(t, totalCostSentByManagerB*GetCardValue(t)/totalCurrentValueSentByManagerA)
			}
		}
		for _, t := range trade.TokensManagerB {
			if errB == nil {
				managerB.sellCardForCardDeal(t)
			}
			if errA == nil {
				managerA.BuyCardForEth(t, totalCostSentByManagerA*GetCardValue(t)/totalCurrentValueSentByManagerB)
			}
		}
	}
}

func (m *Manager) sellCardForEth(tokenId string, soldFor float64) {
	pl := m.ProfitLossCards[tokenId]
	pl.SoldFor = soldFor
	pl.Unrealized = 0
	pl.Realized = roundTo3Dec(soldFor - pl.Cost)
	m.ProfitLossCards[tokenId] = pl
}

func (m *Manager) sellCardForCardDeal(tokenId string) {
	pl := m.ProfitLossCards[tokenId]
	pl.Unrealized = 0
	pl.SoldFor = pl.Cost
	m.ProfitLossCards[tokenId] = pl
}

func (m *Manager) GetCardAsReward(tokenId string) {
	m.ProfitLossCards[tokenId] = ProfitLossCard{TokenId: tokenId, SorareAddress: m.SorareAddress, Cost: 0, SoldFor: -1, Realized: 0, Unrealized: GetCardValue(tokenId)}
}

func (m *Manager) BuyCardForEth(tokenId string, price float64) {
	m.ProfitLossCards[tokenId] = ProfitLossCard{TokenId: tokenId, SorareAddress: m.SorareAddress, Cost: roundTo3Dec(price), SoldFor: -1, Realized: 0, Unrealized: roundTo3Dec(price - GetCardValue(tokenId))}
}

type TradeBuilder struct {
	Trade Trade
}

func NewBuilder() *TradeBuilder {
	return &TradeBuilder{
		Trade: Trade{},
	}
}

func (tb *TradeBuilder) SetAddressManagerA(address string) *TradeBuilder {
	tb.Trade.AddressManagerA = address
	return tb
}
func (tb *TradeBuilder) SetAddressManagerB(address string) *TradeBuilder {
	tb.Trade.AddressManagerB = address
	return tb
}
func (tb *TradeBuilder) SetEthManagerA(eth float64) *TradeBuilder {
	tb.Trade.EthManagerA = eth
	return tb
}
func (tb *TradeBuilder) SetEthManagerB(eth float64) *TradeBuilder {
	tb.Trade.EthManagerB = eth
	return tb
}
func (tb *TradeBuilder) SetTokensManagerA(tokens []string) *TradeBuilder {
	tb.Trade.TokensManagerA = tokens
	return tb
}
func (tb *TradeBuilder) SetTokensManagerB(tokens []string) *TradeBuilder {
	tb.Trade.TokensManagerB = tokens
	return tb
}

func roundTo3Dec(n float64) float64 {
	return math.Round(n*1000) / 1000
}
