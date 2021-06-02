package module

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/container"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	"github.com/cosmos/cosmos-sdk/x/params/types"
)

type Inputs struct {
	container.StructArgs

	Codec        codec.BinaryCodec
	LegacyAmino  *codec.LegacyAmino
	Key          *sdk.KVStoreKey
	TransientKey *sdk.TransientStoreKey
}

type Outputs struct {
	container.StructArgs

	Keeper paramskeeper.Keeper `security-role:"admin"`
}

func (m Module) NewAppModule(inputs Inputs) (module.AppModule, Outputs, error) {
	keeper := paramskeeper.NewKeeper(inputs.Codec, inputs.LegacyAmino, inputs.Key, inputs.TransientKey)
	appMod := params.NewAppModule(keeper)

	return appMod, Outputs{Keeper: keeper}, nil
}

func (m Module) Provide(registrar container.Registrar) error {
	return registrar.RegisterProvider(func(scope container.Scope, keeper paramskeeper.Keeper) types.Subspace {
		return keeper.Subspace(string(scope))
	})
}