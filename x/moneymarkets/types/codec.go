package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "moneymarket/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "moneymarket/BuyName", nil)
	cdc.RegisterConcrete(MsgCreateMarket{}, "moneymarket/CreateMarket", nil)
	cdc.RegisterConcrete(MsgSupplyMarket{}, "moneymarket/SupplyMarket", nil)
	cdc.RegisterConcrete(MsgBorrowFromMarket{}, "moneymarket/BorrowFromMarket", nil)
	cdc.RegisterConcrete(MsgRedeemFromMarket{}, "moneymarket/RedeemFromMarket", nil)
	cdc.RegisterConcrete(MsgRepayToMarket{}, "moneymarket/RepayToMarket", nil)
}
