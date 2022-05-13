package simulation

import (
	"math/rand"

	"github.com/Stride-labs/stride/x/stakeibc/keeper"
	"github.com/Stride-labs/stride/x/stakeibc/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgLiquidStake(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLiquidStake{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the LiquidStake simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "LiquidStake simulation not implemented"), nil, nil
	}
}