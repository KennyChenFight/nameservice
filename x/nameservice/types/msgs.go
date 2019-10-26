package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

const RouterKey = ModuleName

type MsgSetName struct {
	Name string 			`json:"name"`
	Value string 			`json:"value"`
	Owner sdk.AccAddress	`json:"owner"`
}

func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

func (msg MsgSetName) Route() string {
	return RouterKey
}

func (msg MsgSetName) Type() string {
	return "set_name"
}

func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

func (msg MsgSetName) GetSignBytes() []byte {
	return sdk.MustSortJSON(auth.ModuleCdc.MustMarshalJSON(msg))
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

type MsgBuyName struct {
	Name string 			`json:"name"`
	Bid sdk.Coins 			`json:"bid"`
	Buyer sdk.AccAddress	`json:"buyer"`
}

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
	return sdk.MustSortJSON(auth.ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgBuyName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}

