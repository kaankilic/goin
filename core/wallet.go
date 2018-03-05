package wallet

import (
	"crypto/sha256"
	"time"
    "encoding/base64"
	"strconv"
	"fmt"
)
type Wallet struct{
	walletHash string
	balance int
}
func createWallet() Wallet{
	sha := sha256.New()
	sha.Write([]byte(time.Now().String()))
	return Wallet{walletHash:base64.URLEncoding.EncodeToString(sha.Sum(nil)),balance:0}
}