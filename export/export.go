package export

import (
	"fmt"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/state/txindex/kv"
	"github.com/tendermint/tendermint/store"
	dbm "github.com/tendermint/tm-db"
)

func Run(likedPath string) {
	likedDataDir := fmt.Sprintf("%s/data", likedPath)
	blockDB, err := dbm.NewGoLevelDB("blockstore", likedDataDir)
	if err != nil {
		fmt.Println("Cannot initialize blockstore database from liked data", err)
		os.Exit(0)
	}
	defer blockDB.Close()
	blockStore := store.NewBlockStore(blockDB)
	txIndexDB, err := dbm.NewGoLevelDB("tx_index", likedDataDir)
	if err != nil {
		fmt.Println("Cannot initialize tx_index database from liked data", err)
		os.Exit(0)
	}
	defer txIndexDB.Close()
	txIndexer := kv.NewTxIndex(txIndexDB)
	maxHeight := blockStore.Height()

	var startHeight int64 = 1
	for height := startHeight; height < maxHeight; height++ {
		block := blockStore.LoadBlock(height)
		txs := block.Data.Txs
		for _, tx := range txs {
			txHash := bytes.HexBytes(tx.Hash())
			txResult, err := txIndexer.Get(txHash)
			var txRes sdk.TxResponse
			if err != nil || txResult == nil {
				//logger.L.Warnw("Invalid transaction result, replacing with empty result", "txhash", txHash)
				txRes = sdk.TxResponse{Height: height, TxHash: txHash.String()}
			} else {
				txRes, err = formatTxResult(txHash, txResult, block)
				if err != nil {
					//logger.L.Panicw("Cannot parse transaction", "txhash", txHash, "tx_raw", txResult.Tx, "error", err)
				}
			}
			txResJSON, err := encodingConfig.Marshaler.MarshalJSON(&txRes)
			if err == nil {
				fmt.Println(string(txResJSON))
			}
		}
	}
}
