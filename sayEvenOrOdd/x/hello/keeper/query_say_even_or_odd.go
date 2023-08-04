package keeper

import (
	"context"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hello/x/hello/types"
)

func (k Keeper) SayEvenOrOdd(goCtx context.Context, req *types.QuerySayEvenOrOddRequest) (*types.QuerySayEvenOrOddResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	number := req.Number
	intNumber, err := strconv.Atoi(number)
	if err != nil {
		return &types.QuerySayEvenOrOddResponse{Result: fmt.Sprintf("%s is NOT a number", number)}, nil
	} else {
		if intNumber%2 == 0 {
			return &types.QuerySayEvenOrOddResponse{Result: fmt.Sprintf("%d is even", intNumber)}, nil
		} else {
			return &types.QuerySayEvenOrOddResponse{Result: fmt.Sprintf("%d is odd", intNumber)}, nil
		}
	}
}
