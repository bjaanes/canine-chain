package keeper

import (
	"context"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/rns/types"
)

func (k msgServer) Delist(goCtx context.Context, msg *types.MsgDelist) (*types.MsgDelistResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sale, found := k.GetForsale(ctx, msg.Name)

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name isn't listed.")
	}

	n, tld, err := getNameAndTLD(msg.Name)
	if err != nil {
		return nil, err
	}

	name, nfound := k.GetNames(ctx, n, tld)

	if !nfound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name does not exist or has expired.")
	}

	if sale.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "You do not own this listing.")
	}

	if name.Value != sale.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "This listing has expired.")
	}

	k.RemoveForsale(ctx, msg.Name)

	return &types.MsgDelistResponse{}, nil
}
