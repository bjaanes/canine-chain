package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/x/filetree/keeper"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FiletreeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
