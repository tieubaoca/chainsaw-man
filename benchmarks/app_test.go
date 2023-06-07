package benchmarks

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	"gitlab.com/onechain/saw/app"
)

func setup(db dbm.DB, withGenesis bool, invCheckPeriod uint, opts ...wasm.Option) (*app.App, app.GenesisState) {
	encodingConfig := app.MakeEncodingConfig()
	wasmApp := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, app.DefaultNodeHome, invCheckPeriod, encodingConfig, wasm.EnableAllProposals, app.EmptyBaseAppOptions{}, opts)
	if withGenesis {
		return wasmApp, app.NewDefaultGenesisState(app.MakeEncodingConfig().Codec)
	}
	return wasmApp, app.GenesisState{}
}

// SetupWithGenesisAccounts initializes a new WasmApp with the provided genesis
// accounts and possible balances.
func SetupWithGenesisAccounts(b testing.TB, db dbm.DB, genAccs []authtypes.GenesisAccount, balances ...banktypes.Balance) *app.App {
	wasmApp, genesisState := setup(db, true, 0)
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), genAccs)
	appCodec := wasmApp.AppCodec()

	genesisState[authtypes.ModuleName] = appCodec.MustMarshalJSON(authGenesis)

	totalSupply := sdk.NewCoins()
	for _, b := range balances {
		totalSupply = totalSupply.Add(b.Coins...)
	}

	bankGenesis := banktypes.NewGenesisState(banktypes.DefaultGenesisState().Params, balances, totalSupply, []banktypes.Metadata{})
	genesisState[banktypes.ModuleName] = appCodec.MustMarshalJSON(bankGenesis)

	stateBytes, err := json.MarshalIndent(genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	wasmApp.InitChain(
		abci.RequestInitChain{
			Validators: []abci.ValidatorUpdate{},
			ConsensusParams: &abci.ConsensusParams{
				Block: &abci.BlockParams{
					MaxBytes: 8000000,
					MaxGas:   1234000000,
				},
				Evidence: &tmproto.EvidenceParams{
					MaxAgeNumBlocks: 302400,
					MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
					MaxBytes:        10000,
				},
				Validator: &tmproto.ValidatorParams{
					PubKeyTypes: []string{
						tmtypes.ABCIPubKeyTypeEd25519,
					},
				},
			},
			AppStateBytes: stateBytes,
		},
	)

	wasmApp.Commit()
	wasmApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: wasmApp.LastBlockHeight() + 1}})

	return wasmApp
}

type AppInfo struct {
	App          *app.App
	MinterKey    *secp256k1.PrivKey
	MinterAddr   sdk.AccAddress
	ContractAddr string
	Denom        string
	AccNum       uint64
	SeqNum       uint64
	TxConfig     client.TxConfig
}

func InitializeWasmApp(b testing.TB, db dbm.DB, numAccounts int) AppInfo {
	// constants
	minter := secp256k1.GenPrivKey()
	addr := sdk.AccAddress(minter.PubKey().Address())
	denom := "uatom"

	// genesis setup (with a bunch of random accounts)
	genAccs := make([]authtypes.GenesisAccount, numAccounts+1)
	bals := make([]banktypes.Balance, numAccounts+1)
	genAccs[0] = &authtypes.BaseAccount{
		Address: addr.String(),
	}
	bals[0] = banktypes.Balance{
		Address: addr.String(),
		Coins:   sdk.NewCoins(sdk.NewInt64Coin(denom, 100000000000)),
	}
	for i := 0; i <= numAccounts; i++ {
		acct := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
		if i == 0 {
			acct = addr.String()
		}
		genAccs[i] = &authtypes.BaseAccount{
			Address: acct,
		}
		bals[i] = banktypes.Balance{
			Address: acct,
			Coins:   sdk.NewCoins(sdk.NewInt64Coin(denom, 100000000000)),
		}
	}
	wasmApp := SetupWithGenesisAccounts(b, db, genAccs, bals...)

	// add wasm contract
	height := int64(2)
	wasmApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: height, Time: time.Now()}})

	// upload the code

	// instantiate the contract

	// TODO: parse contract address better

	return AppInfo{
		App:        wasmApp,
		MinterKey:  minter,
		MinterAddr: addr,
		Denom:      denom,
		AccNum:     0,
		SeqNum:     2,
		TxConfig:   simappparams.MakeTestEncodingConfig().TxConfig,
	}
}

func GenSequenceOfTxs(b testing.TB, info *AppInfo, msgGen func(*AppInfo) ([]sdk.Msg, error), numToGenerate int) []sdk.Tx {
	fees := sdk.Coins{sdk.NewInt64Coin(info.Denom, 0)}
	txs := make([]sdk.Tx, numToGenerate)

	for i := 0; i < numToGenerate; i++ {
		msgs, err := msgGen(info)
		require.NoError(b, err)
		txs[i], err = helpers.GenTx(
			info.TxConfig,
			msgs,
			fees,
			1234567,
			"",
			[]uint64{info.AccNum},
			[]uint64{info.SeqNum},
			info.MinterKey,
		)
		require.NoError(b, err)
		info.SeqNum += 1
	}

	return txs
}

type balance struct {
	Address string `json:"address"`
	Amount  uint64 `json:"amount,string"`
}
