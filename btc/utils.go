/*
Copyright 2017 Idealnaya rabota LLC
Licensed under Multy.io license.
See LICENSE for details
*/
package btc

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Appscrunch/Multy-back/store"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func newEmptyTx(userID string) store.TxRecord {
	return store.TxRecord{
		UserID:       userID,
		Transactions: []store.MultyTX{},
	}
}
func newAddresAmount(address string, amount int64) store.AddresAmount {
	return store.AddresAmount{
		Address: address,
		Amount:  amount,
	}
}

//func newMultyTX(tx store.MultyTX) store.MultyTX {
//	return store.MultyTX{
//		TxID:              tx.TxID,
//		TxHash:            tx.TxHash,
//		TxOutScript:       tx.TxOutScript,
//		TxAddress:         tx,
//		TxStatus:          txStatus,
//		TxOutAmount:       txOutAmount,
//		TxOutID:           txOutID,
//		WalletIndex:       walletindex,
//		BlockTime:         blockTime,
//		BlockHeight:       blockHeight,
//		TxFee:             fee,
//		MempoolTime:       mempoolTime,
//		StockExchangeRate: stockexchangerate,
//		TxInputs:          inputs,
//		TxOutputs:         outputs,
//	}
//}

func rawTxByTxid(txid string) (*btcjson.TxRawResult, error) {
	hash, err := chainhash.NewHashFromStr(txid)
	if err != nil {
		return nil, err
	}
	previousTxVerbose, err := rpcClient.GetRawTransactionVerbose(hash)
	if err != nil {
		return nil, err
	}
	return previousTxVerbose, nil
}

func fetchWalletAndAddressIndexes(wallets []store.Wallet, address string) (int, int) {
	var walletIndex int
	var addressIndex int
	for _, wallet := range wallets {
		for _, addr := range wallet.Adresses {
			if addr.Address == address {
				walletIndex = wallet.WalletIndex
				addressIndex = addr.AddressIndex
				break
			}
		}
	}
	return walletIndex, addressIndex
}

func setTransactionInfo(multyTx *store.MultyTX, txVerbose *btcjson.TxRawResult) error {

	inputs := []store.AddresAmount{}
	outputs := []store.AddresAmount{}
	var inputSum float64
	var outputSum float64

	for _, out := range txVerbose.Vout {
		for _, address := range out.ScriptPubKey.Addresses {
			amount := int64(out.Value * 100000000)
			outputs = append(outputs, newAddresAmount(address, amount))
		}
		outputSum += out.Value
	}
	for _, input := range txVerbose.Vin {
		hash, err := chainhash.NewHashFromStr(input.Txid)
		if err != nil {
			log.Errorf("txInfo:chainhash.NewHashFromStr: %s", err.Error())

		}
		previousTxVerbose, err := rpcClient.GetRawTransactionVerbose(hash)
		if err != nil {
			log.Errorf("txInfo:rpcClient.GetRawTransactionVerbose: %s", err.Error())
		}

		for _, address := range previousTxVerbose.Vout[input.Vout].ScriptPubKey.Addresses {
			amount := int64(previousTxVerbose.Vout[input.Vout].Value * 100000000)
			inputs = append(inputs, newAddresAmount(address, amount))
		}
		inputSum += previousTxVerbose.Vout[input.Vout].Value
	}
	fee := int64((inputSum - outputSum) * 100000000)

	multyTx.TxInputs = inputs
	multyTx.TxOutputs = outputs
	multyTx.TxFee = fee

	return nil
}

/*

Main process BTC transaction method

can be called from:
- Mempool
- New block
- Resync

*/

func processTransaction(blockChainBlockHeight int64, txVerbose *btcjson.TxRawResult) {
	var multyTx *store.MultyTX = parseRawTransaction(blockChainBlockHeight, txVerbose)
	if multyTx != nil {

		setTransactionInfo(multyTx, txVerbose)
		log.Debugf("processTransaction:setTransactionInfo %v", multyTx)

		transactions := splitTransaction(*multyTx, blockChainBlockHeight)
		log.Debugf("processTransaction:splitTransaction %v", transactions)

		for _, transaction := range transactions {

			finalizeTransaction(&transaction, txVerbose)

			updateWalletAndAddressDate(transaction)

			saveMultyTransaction(transaction)
			sendNotifyToClients(transaction)
		}
	}
}

/*
This method should parse raw transaction from BTC node

_________________________
Inputs:
* blockChainBlockHeight int64 - could be:
-1 in case of mempool call
>1 in case of block transaction
max chain height in case of resync

*txVerbose - raw BTC transaction
_________________________
Output:
* multyTX - multy transaction Structure

*/
func parseRawTransaction(blockChainBlockHeight int64, txVerbose *btcjson.TxRawResult) *store.MultyTX {
	multyTx := store.MultyTX{}

	err := parseInputs(txVerbose, blockChainBlockHeight, &multyTx)
	if err != nil {
		log.Errorf("parseRawTransaction:parseInputs: %s", err.Error())
	}

	err = parseOutputs(txVerbose, blockChainBlockHeight, &multyTx)
	if err != nil {
		log.Errorf("parseRoawTransaction:parseOutputs: %s", err.Error())
	}

	if multyTx.TxID != "" {
		multyTx.TxOutScript = txVerbose.Hex

		return &multyTx
	} else {
		return nil
	}
}

/*
This method need if we have one transaction with more the one u wser'sallet
That means that from one btc transaction we should build more the one Multy Transaction
*/
func splitTransaction(multyTx store.MultyTX, blockHeight int64) []store.MultyTX {
	// transactions := make([]store.MultyTX, 1)
	transactions := []store.MultyTX{}

	currentBlockHeight, err := rpcClient.GetBlockCount()
	if err != nil {
		log.Errorf("splitTransaction:getBlockCount: %s", err.Error())
	}

	blockDiff := currentBlockHeight - blockHeight

	//This is implementatios for single wallet transaction for multi addresses not for multi wallets!
	if multyTx.WalletsInput != nil && len(multyTx.WalletsInput) > 0 {
		outgoingTx := newEntity(multyTx)
		// outgoingTx.WalletsOutput = make([]store.WalletForTx, 1)
		outgoingTx.WalletsOutput = []store.WalletForTx{}

		for _, walletOutput := range multyTx.WalletsOutput {
			for _, walletInput := range outgoingTx.WalletsInput {
				if walletInput.UserId == walletOutput.UserId && walletInput.WalletIndex == walletOutput.WalletIndex {
					outgoingTx.WalletsOutput = append(outgoingTx.WalletsOutput, walletOutput)
				}
			}
		}

		setTransactionStatus(&outgoingTx, blockDiff, currentBlockHeight, true)
		transactions = append(transactions, outgoingTx)
	}

	if multyTx.WalletsOutput != nil && len(multyTx.WalletsOutput) > 0 {
		for _, walletOutput := range multyTx.WalletsOutput {
			var alreadyAdded = false
			for i := 0; i < len(transactions); i++ {
				//for _, splitedTx := range transactions{
				//Check if our output wallet is in the inputs
				//var walletOutputExistInInputs = false
				if transactions[i].WalletsInput != nil && len(transactions[i].WalletsInput) > 0 {
					for _, splitedInput := range transactions[i].WalletsInput {
						if splitedInput.UserId == walletOutput.UserId && splitedInput.WalletIndex == walletOutput.WalletIndex {
							alreadyAdded = true
						}
					}
				}

				if transactions[i].WalletsOutput != nil && len(transactions[i].WalletsOutput) > 0 {
					//var alreadyInOutputs = false
					for j := 0; j < len(transactions[i].WalletsOutput); j++ {
						if walletOutput.UserId == transactions[i].WalletsOutput[j].UserId && walletOutput.WalletIndex == transactions[i].WalletsOutput[j].WalletIndex { //&& walletOutput.Address.Address != transactions[i].WalletsOutput[j].Address.Address Don't think this ckeck we need
							//We have the same wallet index in output but different addres
							alreadyAdded = true
							//alreadyInOutputs = true
							transactions[i].WalletsOutput = append(transactions[i].WalletsOutput, walletOutput)
						}
						//else if walletOutput.UserId == transactions[i].WalletsOutput[j].UserId && walletOutput.WalletIndex == transactions[i].WalletsOutput[j].WalletIndex && walletOutput.Address.Address == transactions[i].WalletsOutput[j].Address.Address{
						//	//alreadyInOutputs = true
						//	alreadyAdded = true
						//}
					}
				}

			}

			if alreadyAdded {
				continue
			} else {
				//Add output transaction here
				incomingTx := newEntity(multyTx)
				incomingTx.WalletsInput = nil
				// incomingTx.WalletsOutput = make([]store.WalletForTx, 1)
				incomingTx.WalletsOutput = []store.WalletForTx{}
				incomingTx.WalletsOutput = append(incomingTx.WalletsOutput, walletOutput)
				setTransactionStatus(&incomingTx, blockDiff, currentBlockHeight, false)
				transactions = append(transactions, incomingTx)
			}
		}

	}

	return transactions
}

func newEntity(multyTx store.MultyTX) store.MultyTX {
	newTx := store.MultyTX{
		TxID:        multyTx.TxID,
		TxHash:      multyTx.TxHash,
		TxOutScript: multyTx.TxOutScript,
		TxAddress:   multyTx.TxAddress,
		TxStatus:    multyTx.TxStatus,
		TxOutAmount: multyTx.TxOutAmount,
		//TxOutIndexes:      multyTx.TxOutIndexes,
		//TxInAmount:        multyTx.TxInAmount,
		//TxInIndexes:       multyTx.TxInIndexes,
		BlockTime:         multyTx.BlockTime,
		BlockHeight:       multyTx.BlockHeight,
		TxFee:             multyTx.TxFee,
		MempoolTime:       multyTx.MempoolTime,
		StockExchangeRate: multyTx.StockExchangeRate,
		TxInputs:          multyTx.TxInputs,
		TxOutputs:         multyTx.TxOutputs,
		WalletsInput:      multyTx.WalletsInput,
		WalletsOutput:     multyTx.WalletsOutput,
	}
	return newTx
}

func saveMultyTransaction(tx store.MultyTX) {

	// This is splited transaction! That means that transaction's WalletsInputs and WalletsOutput have the same WalletIndex!

	//Here we have outgoing transaction for exact wallet!
	if tx.WalletsInput != nil && len(tx.WalletsInput) >0{
		sel := bson.M{"userid": tx.WalletsInput[0].UserId, "transactions.txid": tx.TxID, "transactions.walletsinput.walletindex": tx.WalletsInput[0].WalletIndex}
		update := bson.M{
			"$set": bson.M{
				"transactions.$.txstatus":    tx.TxStatus,
				"transactions.$.blockheight": tx.BlockHeight,
				"transactions.$.blocktime":   tx.BlockTime,
			},
		}
		err := txsData.Update(sel, update)
		if err != nil {
			log.Errorf("saveMultyTransaction:txsData.Update %s", err.Error())
		}

		if err == mgo.ErrNotFound {
			sel := bson.M{"userid": tx.WalletsInput[0].UserId}
			update := bson.M{"$push": bson.M{"transactions": tx}}
			err := txsData.Update(sel, update)
			if err != nil {
				log.Errorf("parseInput.Update add new tx to user: %s", err.Error())
			}
		}
	} else if tx.WalletsOutput != nil && len(tx.WalletsOutput) > 0{

		sel := bson.M{"userid": tx.WalletsOutput[0].UserId, "transactions.txid": tx.TxID, "transactions.walletsoutput.walletindex": tx.WalletsOutput[0].WalletIndex}
		update := bson.M{
			"$set": bson.M{
				"transactions.$.txstatus":    tx.TxStatus,
				"transactions.$.blockheight": tx.BlockHeight,
				"transactions.$.blocktime":   tx.BlockTime,
			},
		}
		err := txsData.Update(sel, update)

		if err == mgo.ErrNotFound {
			sel := bson.M{"userid": tx.WalletsOutput[0].UserId}
			update := bson.M{"$push": bson.M{"transactions": tx}}
			err := txsData.Update(sel, update)
			if err != nil {
				log.Errorf("parseInput.Update add new tx to user: %s", err.Error())
			}
		}
	}

	//
	//
	//
	//for _, walletInput := range tx.WalletsInput {
	//
	//	sel := bson.M{"userid": walletInput.UserId, "transactions.txid": tx.TxID, "transactions.walletsoutput.walletindex": walletInput.WalletIndex}
	//	update := bson.M{
	//		"$set": bson.M{
	//			"transactions.$.txstatus":    tx.TxStatus,
	//			"transactions.$.blockheight": tx.BlockHeight,
	//			"transactions.$.blocktime":   tx.BlockTime,
	//		},
	//	}
	//	err := txsData.Update(sel, update)
	//	if err != nil {
	//		log.Errorf("saveMultyTransaction:txsData.Update %s", err.Error())
	//	}
	//
	//	if err == mgo.ErrNotFound {
	//		sel := bson.M{"userid": walletInput.UserId}
	//		update := bson.M{"$push": bson.M{"transactions": tx}}
	//		err := txsData.Update(sel, update)
	//		if err != nil {
	//			log.Errorf("parseInput.Update add new tx to user: %s", err.Error())
	//		}
	//	}
	//}
	//
	//for _, walletOutput := range tx.WalletsOutput {
	//	sel := bson.M{"userid": walletOutput.UserId, "transactions.txid": tx.TxID, "transactions.walletsoutput.walletindex": walletOutput.WalletIndex}
	//	update := bson.M{
	//		"$set": bson.M{
	//			"transactions.$.txstatus":    tx.TxStatus,
	//			"transactions.$.blockheight": tx.BlockHeight,
	//			"transactions.$.blocktime":   tx.BlockTime,
	//		},
	//	}
	//	err := txsData.Update(sel, update)
	//
	//	if err == mgo.ErrNotFound {
	//		sel := bson.M{"userid": walletOutput.UserId}
	//		update := bson.M{"$push": bson.M{"transactions": tx}}
	//		err := txsData.Update(sel, update)
	//		if err != nil {
	//			log.Errorf("parseInput.Update add new tx to user: %s", err.Error())
	//		}
	//	}
	//}

}

func sendNotify(txMsq *BtcTransactionWithUserID) {
	newTxJSON, err := json.Marshal(txMsq)
	if err != nil {
		log.Errorf("sendNotifyToClients: [%+v] %s\n", txMsq, err.Error())
		return
	}

	err = nsqProducer.Publish(TopicTransaction, newTxJSON)
	if err != nil {
		log.Errorf("nsq publish new transaction: [%+v] %s\n", txMsq, err.Error())
		return
	}
	return
}

func sendNotifyToClients(tx store.MultyTX) {

	for _, walletOutput := range tx.WalletsOutput {
		txMsq := BtcTransactionWithUserID{
			UserID: walletOutput.UserId,
			NotificationMsg: &BtcTransaction{
				TransactionType: tx.TxStatus,
				Amount:          tx.TxOutAmount,
				TxID:            tx.TxID,
				Address:         walletOutput.Address.Address,
			},
		}
		sendNotify(&txMsq)
	}

	for _, walletInput := range tx.WalletsInput {
		txMsq := BtcTransactionWithUserID{
			UserID: walletInput.UserId,
			NotificationMsg: &BtcTransaction{
				TransactionType: tx.TxStatus,
				Amount:          tx.TxOutAmount,
				TxID:            tx.TxID,
				Address:         walletInput.Address.Address,
			},
		}
		sendNotify(&txMsq)
	}

	//TODO make it correct
	//func sendNotifyToClients(txMsq *BtcTransactionWithUserID) {
	//	newTxJSON, err := json.Marshal(txMsq)
	//	if err != nil {
	//		log.Errorf("sendNotifyToClients: [%+v] %s\n", txMsq, err.Error())
	//		return
	//	}
	//
	//	err = nsqProducer.Publish(TopicTransaction, newTxJSON)
	//	if err != nil {
	//		log.Errorf("nsq publish new transaction: [%+v] %s\n", txMsq, err.Error())
	//		return
	//	}
	//	return
}

func parseInputs(txVerbose *btcjson.TxRawResult, blockHeight int64, multyTx *store.MultyTX) error {
	//NEW LOGIC
	user := store.User{}
	//Ranging by inputs
	for _, input := range txVerbose.Vin {

		//getting previous verbose transaction from BTC Node for checking addresses
		previousTxVerbose, err := rawTxByTxid(input.Txid)
		if err != nil {
			log.Errorf("parseInput:rawTxByTxid: %s", err.Error())
			continue
		}

		for _, txInAddress := range previousTxVerbose.Vout[input.Vout].ScriptPubKey.Addresses {
			query := bson.M{"wallets.addresses.address": txInAddress}

			err := usersData.Find(query).One(&user)
			if err != nil {
				continue
				// is not our user
			}
			fmt.Println("[ITS OUR USER] ", user.UserID)

			walletIndex, addressIndex := fetchWalletAndAddressIndexes(user.Wallets, txInAddress)

			txInAmount := int64(100000000 * previousTxVerbose.Vout[input.Vout].Value)

			currentWallet := store.WalletForTx{UserId: user.UserID, WalletIndex: walletIndex}

			//if multyTx.TxInputs == nil {
			//	multyTx.TxInputs = make([]store.AddresAmount, 2)
			//}

			if multyTx.WalletsInput == nil {
				// multyTx.WalletsInput = make([]store.WalletForTx, 2)
				multyTx.WalletsInput = []store.WalletForTx{}
			}

			currentWallet.Address = store.AddressForWallet{Address: txInAddress, AddressIndex: addressIndex, Amount: txInAmount}
			multyTx.WalletsInput = append(multyTx.WalletsInput, currentWallet)

			//multyTx.TxInputs = append(multyTx.TxInputs, store.AddresAmount{Address: txInAddress, Amount: txInAmount})

			multyTx.TxID = txVerbose.Txid
			multyTx.TxHash = txVerbose.Hash

		}

	}

	return nil

	//OLD LOGIC
	//user := store.User{}
	//blockTimeUnix := time.Now().Unix()
	//
	////Ranging by inputs
	//for _, input := range txVerbose.Vin {
	//
	//	//getting previous verbose transaction from BTC Node for checking addresses
	//	previousTxVerbose, err := rawTxByTxid(input.Txid)
	//	if err != nil {
	//		log.Errorf("parseInput:rawTxByTxid: %s", err.Error())
	//		continue
	//	}
	//
	//	for _, address := range previousTxVerbose.Vout[input.Vout].ScriptPubKey.Addresses {
	//		query := bson.M{"wallets.addresses.address": address}
	//		// Is it's our user transaction.
	//		err := usersData.Find(query).One(&user)
	//		if err != nil {
	//			continue
	//			// Is not our user.
	//		}
	//
	//		log.Debugf("[ITS OUR USER] %s", user.UserID)
	//
	//		inputs, outputs, fee, err := txInfo(txVerbose)
	//		if err != nil {
	//			log.Errorf("parseInput:txInfo:input: %s", err.Error())
	//			continue
	//		}
	//
	//
	//		walletIndex := fetchWalletIndex(user.Wallets, address)
	//
	//
	//
	//		// Is our user already have this transactions.
	//		sel := bson.M{"userid": user.UserID, "transactions.txid": txVerbose.Txid, "transactions.txaddress": address}
	//		err = txsData.Find(sel).One(nil)
	//		if err == mgo.ErrNotFound {
	//			// User have no transaction like this. Add to DB.
	//			txOutAmount := int64(100000000 * previousTxVerbose.Vout[input.Vout].Value)
	//
	//			// Set bloct time -1 if tx from mempool.
	//			blockTime := blockTimeUnix
	//			if blockHeight == -1 {
	//				blockTime = int64(-1)
	//			}
	//
	//			newTx := newMultyTX(txVerbose.Txid, txVerbose.Hash, previousTxVerbose.Vout[input.Vout].ScriptPubKey.Hex, address, txStatus, int(previousTxVerbose.Vout[input.Vout].N), walletIndex, txOutAmount, blockTime, blockHeight, fee, blockTimeUnix, exRates, inputs, outputs)
	//			sel = bson.M{"userid": user.UserID}
	//			update := bson.M{"$push": bson.M{"transactions": newTx}}
	//			err = txsData.Update(sel, update)
	//			if err != nil {
	//				log.Errorf("parseInput:txsData.Update add new tx to user: %s", err.Error())
	//			}
	//			continue
	//		} else if err != nil && err != mgo.ErrNotFound {
	//			log.Errorf("parseInput:txsData.Find: %s", err.Error())
	//			continue
	//		}
	//
	//		// User have this transaction but with another status.
	//		// Update statsus, block height and block time.
	//		sel = bson.M{"userid": user.UserID, "transactions.txid": txVerbose.Txid, "transactions.txaddress": address}
	//		update = bson.M{
	//			"$set": bson.M{
	//				"transactions.$.txstatus":    txStatus,
	//				"transactions.$.blockheight": blockHeight,
	//				"transactions.$.blocktime":   blockTimeUnix,
	//			},
	//		}
	//		err = txsData.Update(sel, update)
	//		if err != nil {
	//			log.Errorf("parseInput:txsData.Update: %s", err.Error())
	//		}
	//	}
	//}
	//return nil
}

func parseOutputs(txVerbose *btcjson.TxRawResult, blockHeight int64, multyTx *store.MultyTX) error {

	user := store.User{}

	for _, output := range txVerbose.Vout {
		for _, txOutAddress := range output.ScriptPubKey.Addresses {
			query := bson.M{"wallets.addresses.address": txOutAddress}

			err := usersData.Find(query).One(&user)
			if err != nil {
				continue
				// is not our user
			}
			fmt.Println("[ITS OUR USER] ", user.UserID)

			walletIndex, addressIndex := fetchWalletAndAddressIndexes(user.Wallets, txOutAddress)

			currentWallet := store.WalletForTx{UserId: user.UserID, WalletIndex: walletIndex}

			if multyTx.TxOutputs == nil {
				// multyTx.TxOutputs = make([]store.AddresAmount, 2)
				multyTx.TxOutputs = []store.AddresAmount{}
			}

			if multyTx.WalletsOutput == nil {
				// multyTx.WalletsOutput = make([]store.WalletForTx, 2)
				multyTx.WalletsOutput = []store.WalletForTx{}
			}

			currentWallet.Address = store.AddressForWallet{Address: txOutAddress, AddressIndex: addressIndex, Amount: int64(100000000 * output.Value)}
			multyTx.WalletsOutput = append(multyTx.WalletsOutput, currentWallet)

			multyTx.TxOutputs = append(multyTx.TxOutputs, store.AddresAmount{Address: txOutAddress, Amount: int64(100000000 * output.Value)})

			//if len(multyTx.WalletsOutput) > 0{
			//	var haveTheSameWalletIndex = false
			//	//Check if we already have the same wallet index
			//	for _, walletOutForTx := range multyTx.WalletsOutput{
			//		if walletOutForTx.WalletIndex == currentWallet.WalletIndex{
			//			haveTheSameWalletIndex = true
			//		}
			//	}
			//	if !haveTheSameWalletIndex{
			//		//This is not stored wallet
			//		currentWallet.Address = store.AddressWorWallet{Address:txOutAddress, AddressIndex:addressIndex, Amount:int64(100000000 * output.Value)}
			//		multyTx.WalletsOutput = append(multyTx.WalletsOutput, currentWallet)
			//
			//		multyTx.TxOutputs = append(multyTx.TxOutputs, store.AddresAmount{Address:txOutAddress, Amount:int64(100000000 * output.Value)})
			//
			//
			//	}
			//} else {
			//	currentWallet.Address = store.AddressWorWallet{Address:txOutAddress, AddressIndex:addressIndex, Amount:int64(100000000 * output.Value)}
			//	multyTx.WalletsOutput = append(multyTx.WalletsOutput, currentWallet)
			//}

			multyTx.TxID = txVerbose.Txid
			multyTx.TxHash = txVerbose.Hash

		}
	}
	return nil
}

func GetLatestExchangeRate() ([]store.ExchangeRatesRecord, error) {
	selGdax := bson.M{
		"stockexchange": "Gdax",
	}
	selPoloniex := bson.M{
		"stockexchange": "Poloniex",
	}
	stocksGdax := store.ExchangeRatesRecord{}
	err := exRate.Find(selGdax).Sort("-timestamp").One(&stocksGdax)
	if err != nil {
		return nil, err
	}

	stocksPoloniex := store.ExchangeRatesRecord{}
	err = exRate.Find(selPoloniex).Sort("-timestamp").One(&stocksPoloniex)
	if err != nil {
		return nil, err
	}
	return []store.ExchangeRatesRecord{stocksPoloniex, stocksGdax}, nil

}

func updateWalletAndAddressDate(tx store.MultyTX) {

	for _, walletOutput := range tx.WalletsOutput {

		// update addresses last action time
		sel := bson.M{"userID": walletOutput.UserId, "wallets.addresses.address": walletOutput.Address}
		update := bson.M{
			"$set": bson.M{
				"wallets.$.addresses.$[].lastActionTime": time.Now().Unix(),
			},
		}
		err := usersData.Update(sel, update)
		if err != nil {
			log.Errorf("updateWalletAndAddressDate:usersData.Update: %s", err.Error())
		}

		// update wallets last action time
		// Set status to OK if some money transfered to this address
		sel = bson.M{"userID": walletOutput.UserId, "wallets.walletIndex": walletOutput.WalletIndex}
		update = bson.M{
			"$set": bson.M{
				"wallets.$.status":         store.WalletStatusOK,
				"wallets.$.lastActionTime": time.Now().Unix(),
			},
		}
		err = usersData.Update(sel, update)
		if err != nil {
			log.Errorf("updateWalletAndAddressDate:usersData.Update: %s", err.Error())
		}

	}

	for _, walletInput := range tx.WalletsInput {
		// update addresses last action time
		sel := bson.M{"userID": walletInput.UserId, "wallets.addresses.address": walletInput.Address}
		update := bson.M{
			"$set": bson.M{
				"wallets.$.addresses.$[].lastActionTime": time.Now().Unix(),
			},
		}
		err := usersData.Update(sel, update)
		if err != nil {
			log.Errorf("updateWalletAndAddressDate:usersData.Update: %s", err.Error())
		}

		// update wallets last action time
		sel = bson.M{"userID": walletInput.UserId, "wallets.walletIndex": walletInput.WalletIndex}
		update = bson.M{
			"$set": bson.M{
				"wallets.$.lastActionTime": time.Now().Unix(),
			},
		}
		err = usersData.Update(sel, update)
		if err != nil {
			log.Errorf("updateWalletAndAddressDate:usersData.Update: %s", err.Error())
		}
	}

	/*
		// Update wallets last action time on every new transaction.
		// Set status to OK if some money transfered to this address
		sel := bson.M{"userID": user.UserID, "wallets.walletIndex": walletIndex}
		update := bson.M{
			"$set": bson.M{
				"wallets.$.status":         store.WalletStatusOK,
				"wallets.$.lastActionTime": time.Now().Unix(),
			},
		}
		err = usersData.Update(sel, update)
		if err != nil {
			log.Errorf("parseOutput:restClient.userStore.Update: %s", err.Error())
		}

		// Update wallets last action time on every new transaction.
		sel := bson.M{"userID": user.UserID, "wallets.walletIndex": walletIndex}
		update := bson.M{
			"$set": bson.M{
				"wallets.$.lastActionTime": time.Now().Unix(),
			},
		}
		err := usersData.Update(sel, update)
		if err != nil {
			log.Errorf("parseOutput:restClient.userStore.Update: %s", err.Error())
		}

		// Update address last action time on every new transaction.
		sel = bson.M{"userID": user.UserID, "wallets.addresses.address": address}
		update = bson.M{
			"$set": bson.M{
				"wallets.$.addresses.$[].lastActionTime": time.Now().Unix(),
			},
		}
		err = usersData.Update(sel, update)
		if err != nil {
			log.Errorf("parseOutput:restClient.userStore.Update: %s", err.Error())
		}
	*/
}

func setTransactionStatus(tx *store.MultyTX, blockDiff int64, currentBlockHeight int64, fromInput bool) {
	transactionTime := time.Now().Unix()
	if blockDiff > currentBlockHeight {
		//This call was made from memPool
		if fromInput {
			tx.TxStatus = TxStatusAppearedInMempoolOutcoming
			tx.MempoolTime = transactionTime
			tx.BlockTime = -1
		} else {
			tx.TxStatus = TxStatusAppearedInMempoolIncoming
			tx.MempoolTime = transactionTime
			tx.BlockTime = -1
		}

	} else if blockDiff >= 0 && blockDiff < 6 {
		//This call was made from block or resync
		//Transaction have no enough confirmations
		if fromInput {
			tx.TxStatus = TxStatusAppearedInBlockOutcoming
			tx.BlockTime = transactionTime
		} else {
			tx.TxStatus = TxStatusAppearedInBlockIncoming
			tx.BlockTime = transactionTime
		}
	} else if blockDiff >= 6 && blockDiff < currentBlockHeight {
		//This call was made from resync
		//Transaction have enough confirmations
		if fromInput {
			tx.TxStatus = TxStatusInBlockConfirmedOutcoming
		} else {
			tx.TxStatus = TxStatusInBlockConfirmedIncoming
		}
	}
}


func finalizeTransaction(tx *store.MultyTX, txVerbose *btcjson.TxRawResult) {

	if tx.TxAddress == nil {
		// tx.TxAddress = make([]string, 1)
		tx.TxAddress = []string{}
	}

	if tx.TxStatus == TxStatusInBlockConfirmedOutcoming || tx.TxStatus == TxStatusAppearedInBlockOutcoming || tx.TxStatus == TxStatusAppearedInMempoolOutcoming {
		for _, walletInput := range tx.WalletsInput {
			tx.TxOutAmount += walletInput.Address.Amount
			tx.TxAddress = append(tx.TxAddress, walletInput.Address.Address)
		}
	} else {
		for i := 0; i < len(tx.WalletsOutput); i++ {
			tx.TxOutAmount += tx.WalletsOutput[i].Address.Amount
			tx.TxAddress = append(tx.TxAddress, tx.WalletsOutput[i].Address.Address)

			for _, output := range txVerbose.Vout {
				for _, outAddr := range output.ScriptPubKey.Addresses {
					if tx.WalletsOutput[i].Address.Address == outAddr {
						tx.WalletsOutput[i].Address.AddressOutIndex = int(output.N)
					}
				}
			}
		}
		//TxOutIndexes we need only for incoming transactions
	}

	rates, err := GetLatestExchangeRate()
	if err != nil {
		log.Errorf("processTransaction:ExchangeRates: %s", err.Error())
	}

	tx.StockExchangeRate = rates


}