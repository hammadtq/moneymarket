package types

import (
	"fmt"
	"strings"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Whois is a struct that contains all the metadata of a name
type Whois struct {
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
	Price sdk.Coins      `json:"price"`
}

// Initial Starting Price for a name that was never previously owned
var MinNamePrice = sdk.Coins{sdk.NewInt64Coin("nametoken", 1)}

// Returns a new Whois with the minprice as the price
func NewWhois() Whois {
	return Whois{
		Price: MinNamePrice,
	}
}

// implement fmt.Stringer
func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

type MoneyMarket struct {
	Name string         `json:"name"`
	Symbol string       `json:"symbol"`
	Owner sdk.AccAddress `json:"owner"`
	InterestRate sdk.Int	`json:"interestRate"`
	TokenSupply sdk.Coins `json:"tokenSupply"`
	TokenBorrows sdk.Coins `json:"tokenBorrows"`
	BorrowCollateral sdk.Coins `json:"BorrowCollateral"`
	TokenName string 	`json:"tokenName"`
	CollateralToken string `json:"collateralToken"`
}


// Initial Starting Price for a market that was never previously owned
//var MinMarketPrice = sdk.Coins{sdk.NewInt64Coin("bnbtoken", 1)}

// Returns a new MoneyMarket with default rate
func NewMarket() MoneyMarket {
	return MoneyMarket{
		InterestRate: sdk.NewInt(0),
		TokenSupply: sdk.Coins{sdk.NewInt64Coin("nametoken", 0)},
		TokenBorrows: sdk.Coins{sdk.NewInt64Coin("nametoken", 0)},
		BorrowCollateral: sdk.Coins{sdk.NewInt64Coin("xyztoken", 0)},
	}
}

// implement fmt.Stringer
func (w MoneyMarket) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Name: %s
Symbol: %s
InterestRate: %s`, w.Owner, w.Name, w.Symbol, w.InterestRate))
}

type MarketPosition struct {
	Owner sdk.AccAddress `json:"owner"`
	Market string `json:"market"`
	LendTokens sdk.Coins	`json:"lendTokens"`
	BorrowTokens sdk.Coins `json:"borrowTokens"`
	BorrowCollateral sdk.Coins `json:"BorrowCollateral"`
}

// Returns a new MarketPosition with default rate
func NewMarketPosition() MarketPosition {
	return MarketPosition{
		LendTokens: sdk.Coins{sdk.NewInt64Coin("nametoken", 0)},
		BorrowTokens: sdk.Coins{sdk.NewInt64Coin("nametoken", 0)},
		BorrowCollateral: sdk.Coins{sdk.NewInt64Coin("xyztoken", 0)},
	}
}

// implement fmt.Stringer
func (w MarketPosition) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Market: %s
LendTokens: %s
BorrowTokens: %s`, w.Owner, w.Market, w.LendTokens, w.BorrowTokens))
}
