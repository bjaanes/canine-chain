package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveEditors = "remove_editors"

var _ sdk.Msg = &MsgRemoveEditors{}

func NewMsgRemoveEditors(creator string, editorIds string, address string, fileowner string, notifyEditors string, notiForEditors string) *MsgRemoveEditors {
	return &MsgRemoveEditors{
		Creator:        creator,
		EditorIds:      editorIds,
		Address:        address,
		Fileowner:      fileowner,
		NotifyEditors:  notifyEditors,
		NotiForEditors: notiForEditors,
	}
}

func (msg *MsgRemoveEditors) Route() string {
	return RouterKey
}

func (msg *MsgRemoveEditors) Type() string {
	return TypeMsgRemoveEditors
}

func (msg *MsgRemoveEditors) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveEditors) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveEditors) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
