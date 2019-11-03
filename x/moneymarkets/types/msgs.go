package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string         `json:"name"`
	Value string         `json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// Route should return the name of the module
func (msg MsgSetName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetName) Type() string { return "set_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}



//BUY NAMES MESSAGE

// MsgBuyName defines the BuyName message
type MsgBuyName struct {
	Name  string         `json:"name"`
	Bid   sdk.Coins      `json:"bid"`
	Buyer sdk.AccAddress `json:"buyer"`
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

// Route should return the name of the module
func (msg MsgBuyName) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBuyName) Type() string { return "buy_name" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBuyName) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.IsAllPositive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBuyName) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}


//CREATE NEW MONEY MARKET

// MsgBuyName defines the BuyName message
type MsgCreateMarket struct {
	Name  string         `json:"name"`
	Symbol   string      `json:"symbol"`
	InterestRate sdk.Coins		`json:"interestRate"`
	Buyer sdk.AccAddress `json:"buyer"`
	TokenName string `json:"tokenName"`
	CollateralToken string `json:"tokenName"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgCreateMarket(name string, symbol string, interestRate sdk.Coins, buyer sdk.AccAddress, tokenName string, collateralToken string) MsgCreateMarket {

	return MsgCreateMarket{
		Name:  name,
		Symbol:   symbol,
		InterestRate: interestRate,
		Buyer: buyer,
		TokenName: tokenName,
		CollateralToken: collateralToken,
	}
}

// Route should return the name of the module
func (msg MsgCreateMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateMarket) Type() string { return "create_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateMarket) ValidateBasic() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrInvalidAddress(msg.Buyer.String())
	}
	if len(msg.Name) == 0 {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if len(msg.Symbol) == 0 {
		return sdk.ErrInsufficientCoins("Symbol cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

//SUPPLY TO MARKET

type MsgSupplyMarket struct {
	Market  string         `json:"market"`
	LendTokens sdk.Coins      `json:"lendTokens"`
	Supplier sdk.AccAddress `json:"supplier"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgSupplyMarket(market string, coins sdk.Coins, supplier sdk.AccAddress) MsgSupplyMarket {
	return MsgSupplyMarket{
		Market:  market,
		LendTokens: coins,
		Supplier: supplier,
	}
}

// Route should return the name of the module
func (msg MsgSupplyMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSupplyMarket) Type() string { return "supply_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgSupplyMarket) ValidateBasic() sdk.Error {
	if msg.Supplier.Empty() {
		return sdk.ErrInvalidAddress(msg.Supplier.String())
	}
	if len(msg.Market) == 0 {
		return sdk.ErrUnknownRequest("Market cannot be empty")
	}
	if len(msg.LendTokens) == 0 {
		return sdk.ErrInsufficientCoins("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSupplyMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSupplyMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Supplier}
}


//Borrow from Market

type MsgBorrowFromMarket struct {
	Market  string         `json:"market"`
	BorrowTokens sdk.Coins      `json:"borrowTokens"`
	CollateralTokens sdk.Coins      `json:"collateralTokens"`
	Supplier sdk.AccAddress `json:"supplier"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgBorrowFromMarket(market string, coins sdk.Coins, collateralcoins sdk.Coins, supplier sdk.AccAddress) MsgBorrowFromMarket {
	return MsgBorrowFromMarket{
		Market:  market,
		BorrowTokens: coins,
		CollateralTokens: collateralcoins,
		Supplier: supplier,
	}
}

// Route should return the name of the module
func (msg MsgBorrowFromMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBorrowFromMarket) Type() string { return "borrow_from_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgBorrowFromMarket) ValidateBasic() sdk.Error {
	if msg.Supplier.Empty() {
		return sdk.ErrInvalidAddress(msg.Supplier.String())
	}
	if len(msg.Market) == 0 {
		return sdk.ErrUnknownRequest("Market cannot be empty")
	}
	if len(msg.BorrowTokens) == 0 {
		return sdk.ErrInsufficientCoins("You must borrow at least one token")
	}
	if len(msg.CollateralTokens) == 0 {
		return sdk.ErrInsufficientCoins("You must supply at least one collateral token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBorrowFromMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBorrowFromMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Supplier}
}

//Redeem from Market

type MsgRedeemFromMarket struct {
	Market  string         `json:"market"`
	RedeemTokens sdk.Coins      `json:"redeemTokens"`
	Supplier sdk.AccAddress `json:"supplier"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgRedeemFromMarket(market string, coins sdk.Coins, supplier sdk.AccAddress) MsgRedeemFromMarket {
	return MsgRedeemFromMarket{
		Market:  market,
		RedeemTokens: coins,
		Supplier: supplier,
	}
}

// Route should return the name of the module
func (msg MsgRedeemFromMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRedeemFromMarket) Type() string { return "redeem_from_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRedeemFromMarket) ValidateBasic() sdk.Error {
	if msg.Supplier.Empty() {
		return sdk.ErrInvalidAddress(msg.Supplier.String())
	}
	if len(msg.Market) == 0 {
		return sdk.ErrUnknownRequest("Market cannot be empty")
	}
	if len(msg.RedeemTokens) == 0 {
		return sdk.ErrInsufficientCoins("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRedeemFromMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRedeemFromMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Supplier}
}

//Repay to Market

type MsgRepayToMarket struct {
	Market  string         `json:"market"`
	RepayTokens sdk.Coins      `json:"repayTokens"`
	Borrower sdk.AccAddress `json:"supplier"`
}

// NewMsgCreateMarket is the constructor function for MsgBuyName
func NewMsgRepayToMarket(market string, coins sdk.Coins, supplier sdk.AccAddress) MsgRepayToMarket {
	return MsgRepayToMarket{
		Market:  market,
		RepayTokens: coins,
		Borrower: supplier,
	}
}

// Route should return the name of the module
func (msg MsgRepayToMarket) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRepayToMarket) Type() string { return "repay_to_market" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRepayToMarket) ValidateBasic() sdk.Error {
	if msg.Borrower.Empty() {
		return sdk.ErrInvalidAddress(msg.Borrower.String())
	}
	if len(msg.Market) == 0 {
		return sdk.ErrUnknownRequest("Market cannot be empty")
	}
	if len(msg.RepayTokens) == 0 {
		return sdk.ErrInsufficientCoins("You must supply at least one token")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRepayToMarket) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRepayToMarket) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Borrower}
}
