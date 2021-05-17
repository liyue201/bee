package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	crypto2 "github.com/ethersphere/bee/pkg/crypto"
	"github.com/ethersphere/bee/pkg/settlement/swap/chequebook"
	"math/big"
)

func newSigner(privateKey string, chainId int64) chequebook.ChequeSigner {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return chequebook.NewChequeSigner(crypto2.NewDefaultSigner(key), chainId)
}

func main() {
	privateKey := flag.String("private_key", "", "private key")
	payout := flag.Int64("payout", 0, "payout amount")
	chequeBook := flag.String("chequebook", "", "chequebook contract address")
	beneficiary := flag.String("beneficiary", "", "beneficiary address")

	flag.Parse()

	signer := newSigner(*privateKey, 5)
	payoutAmount := big.NewInt(*payout)

	cheque := &chequebook.Cheque{
		Chequebook:       common.HexToAddress(*chequeBook),
		Beneficiary:      common.HexToAddress(*beneficiary),
		CumulativePayout: payoutAmount,
	}

	data, err := signer.Sign(cheque)
	if err != nil {
		fmt.Printf("failed to sign: %v\n", err)
		return
	}
	fmt.Printf("0x%v\n", hex.EncodeToString(data))
}
