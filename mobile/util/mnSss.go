package util

import (
	"strings"

	"github.com/kernelflowlabs/wallet-sdk/crypto"
)

const (
	TotalShares     = 2
	ThresholdShares = 2
)

func SssSplit(mnemonic string) (string, error) {
	shares, err := crypto.PerformSplit(mnemonic, TotalShares, ThresholdShares)
	if err != nil {
		return "", err
	}
	return strings.Join(shares, ","), nil
}

func SssRecover(shares string) (string, error) {
	var clean []string
	for _, line := range strings.Split(shares, ",") {
		if s := strings.TrimSpace(line); s != "" {
			clean = append(clean, s)
		}
	}
	return crypto.PerformRecover(clean)
}
