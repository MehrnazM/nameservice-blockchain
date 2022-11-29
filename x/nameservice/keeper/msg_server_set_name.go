package keeper

import (
	"context"

	"nameservice/x/nameservice/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	whois, found := k.GetWhois(ctx, msg.Name)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Name is not registered")
	}

	if whois.Owner != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner")
	}

	newWhois := types.Whois{
		Name:  msg.Name,
		Index: msg.Name,
		Value: msg.Value,
		Owner: whois.Owner,
		Price: whois.Price,
	}
	k.SetWhois(ctx, newWhois)

	return &types.MsgSetNameResponse{}, nil
}
