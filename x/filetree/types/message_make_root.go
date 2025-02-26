package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMakeRoot = "make_root"

var _ sdk.Msg = &MsgMakeRoot{}

func NewMsgMakeRoot(creator string, account string, rootHashPath string, contents string, editors string, viewers string, trackingNumber string) *MsgMakeRoot {
	return &MsgMakeRoot{
		Creator:        creator,
		Account:        account,
		RootHashPath:   rootHashPath,
		Contents:       contents,
		Editors:        editors,
		Viewers:        viewers,
		TrackingNumber: trackingNumber,
	}
}

func (msg *MsgMakeRoot) Route() string {
	return RouterKey
}

func (msg *MsgMakeRoot) Type() string {
	return TypeMsgMakeRoot
}

func (msg *MsgMakeRoot) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMakeRoot) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMakeRoot) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
