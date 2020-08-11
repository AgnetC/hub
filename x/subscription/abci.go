package subscription

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	end := ctx.BlockTime().Add(-1 * k.CancelDuration(ctx))
	k.IterateCancelSubscriptions(ctx, end, func(_ int, item types.Subscription) (stop bool) {
		if item.Plan == 0 {
			consumed := hub.NewBandwidthFromInt64(0, 0)
			k.IterateQuotas(ctx, item.ID, func(_ int, item types.Quota) bool {
				consumed = consumed.Add(item.Current)
				return false
			})

			amount := item.Deposit.Sub(item.Amount(consumed))
			if err := k.SubtractDeposit(ctx, item.Address, amount); err != nil {
				panic(err)
			}
		}

		k.DeleteCancelSubscriptionAt(ctx, item.StatusAt, item.ID)

		item.Status = hub.StatusInactive
		item.StatusAt = ctx.BlockTime()
		k.SetSubscription(ctx, item)

		return false
	})

	return nil
}
