package keeper

import (
	"context"
	"crypto/sha256"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%s%s", msg.Creator, msg.Hashpath)))
	hash := h.Sum(nil)

	pathString := fmt.Sprintf("%x", hash)

	file := types.Files{
		Contents:      msg.Contents,
		Owner:         msg.Creator,
		ViewingAccess: msg.Viewers,
		EditAccess:    msg.Editors,
		Address:       pathString,
	}

	k.SetFiles(ctx, file)

	return &types.MsgPostFileResponse{Path: pathString}, nil
}
