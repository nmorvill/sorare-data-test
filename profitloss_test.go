package main

import "testing"

func TestTradeKylianMbappe(t *testing.T) {
	ManagersDb = map[string]*Manager{
		"mhg":      {SorareAddress: "mhg", ProfitLossCards: make(map[string]ProfitLossCard)},
		"sounders": {SorareAddress: "sounders", ProfitLossCards: make(map[string]ProfitLossCard)},
		"ymc":      {SorareAddress: "ymc", ProfitLossCards: make(map[string]ProfitLossCard)},
	}
	CardValuesDb = map[string]float64{
		"mbappe": 2.849,
	}

	mhg, _ := GetManager("mhg")
	sounders, _ := GetManager("sounders")
	ymc, _ := GetManager("ymc")

	mhg.BuyCardForEth("mbappe", 3.497)
	trade1 := NewBuilder().
		SetAddressManagerA("mhg").
		SetAddressManagerB("sounders").
		SetTokensManagerA([]string{"mbappe"}).
		SetEthManagerB(9.999).Trade
	trade2 := NewBuilder().
		SetAddressManagerA("sounders").
		SetAddressManagerB("ymc").
		SetTokensManagerA([]string{"mbappe"}).
		SetEthManagerB(49.99).Trade
	MakeTrade(trade1)
	MakeTrade(trade2)

	mhgPL := mhg.ProfitLossCards["mbappe"]
	expectedMhgPL := ProfitLossCard{TokenId: "mbappe", SorareAddress: "mhg", Cost: 3.497, SoldFor: 9.999, Realized: 6.502, Unrealized: 0}
	if mhgPL != expectedMhgPL {
		t.Errorf("Error with MHG Mbappe, got %v want %v", mhgPL, expectedMhgPL)
	}

	soundersPL := sounders.ProfitLossCards["mbappe"]
	expectedSoundersPL := ProfitLossCard{TokenId: "mbappe", SorareAddress: "sounders", Cost: 9.999, SoldFor: 49.99, Realized: 39.991, Unrealized: 0}
	if soundersPL != expectedSoundersPL {
		t.Errorf("Error with Sounders Mbappe, got %v want %v", soundersPL, expectedSoundersPL)
	}

	ymcPL := ymc.ProfitLossCards["mbappe"]
	expectedYmcPL := ProfitLossCard{TokenId: "mbappe", SorareAddress: "ymc", Cost: 49.99, SoldFor: -1, Realized: 0, Unrealized: 47.141}
	if ymcPL != expectedYmcPL {
		t.Errorf("Error with YMC Mbappe, got %v want %v", ymcPL, expectedYmcPL)
	}
}

func TestTradeIgorPlastun(t *testing.T) {
	ManagersDb = map[string]*Manager{
		"temujin":   {SorareAddress: "temujin", ProfitLossCards: make(map[string]ProfitLossCard)},
		"blackflag": {SorareAddress: "blackflag", ProfitLossCards: make(map[string]ProfitLossCard)},
		"alex19":    {SorareAddress: "alex19", ProfitLossCards: make(map[string]ProfitLossCard)},
	}
	CardValuesDb = map[string]float64{
		"pellegrini":  0.180,
		"tiquinho":    0.251,
		"vladochimos": 0.408,
		"dzeko":       0.053,
		"carlos":      0.021,
		"plastun":     0.005,
	}

	tem, _ := GetManager("temujin")
	bf, _ := GetManager("blackflag")
	alex, _ := GetManager("alex19")

	tem.BuyCardForEth("plastun", 0.021)

	alex.BuyCardForEth("pellegrini", 0.07)
	alex.BuyCardForEth("tiquinho", 0.037)

	trade1 := NewBuilder().
		SetAddressManagerA("temujin").
		SetAddressManagerB("blackflag").
		SetTokensManagerA([]string{"plastun"}).
		SetEthManagerB(0.03).Trade
	trade2 := NewBuilder().
		SetAddressManagerA("blackflag").
		SetTokensManagerA([]string{"plastun"}).
		SetEthManagerB(0.19).Trade
	trade3 := NewBuilder().
		SetAddressManagerA("alex19").
		SetTokensManagerA([]string{"pellegrini", "tiquinho"}).
		SetEthManagerA(0.9).
		SetTokensManagerB([]string{"vladochimos", "dzeko", "carlos", "plastun"}).Trade

	MakeTrade(trade1)
	MakeTrade(trade2)
	MakeTrade(trade3)

	temPL := tem.ProfitLossCards["plastun"]
	expectedTemPL := ProfitLossCard{TokenId: "plastun", SorareAddress: "temujin", Cost: 0.021, SoldFor: 0.03, Realized: 0.009, Unrealized: 0}
	if temPL != expectedTemPL {
		t.Errorf("Error with Temujin Plastun, got %v want %v", temPL, expectedTemPL)
	}

	bfPL := bf.ProfitLossCards["plastun"]
	expectedBfPL := ProfitLossCard{TokenId: "plastun", SorareAddress: "blackflag", Cost: 0.03, SoldFor: 0.19, Realized: 0.16, Unrealized: 0}
	if bfPL != expectedBfPL {
		t.Errorf("Error with BlackFlag Plastun, got %v want %v", bfPL, expectedBfPL)
	}

	alexPL := alex.ProfitLossCards["plastun"]
	expectedAlexPL := ProfitLossCard{TokenId: "plastun", SorareAddress: "alex19", Cost: 0.010, SoldFor: -1, Realized: 0, Unrealized: 0.005}
	if alexPL != expectedAlexPL {
		t.Errorf("Error with Alex Plastun, got %v want %v", alexPL, expectedAlexPL)
	}
}
