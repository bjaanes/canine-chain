package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgPostproof = "postproof"

var _ sdk.Msg = &MsgPostproof{}

func NewMsgPostproof(creator string, item string, hashlist string, cid string) *MsgPostproof {
	return &MsgPostproof{
		Creator:  creator,
		Cid:      cid,
		Item:     item,
		Hashlist: hashlist,
	}
}

func (msg *MsgPostproof) Route() string {
	return RouterKey
}

func (msg *MsgPostproof) Type() string {
	return TypeMsgPostproof
}

func (msg *MsgPostproof) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPostproof) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPostproof) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
