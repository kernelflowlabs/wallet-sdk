package chainrpc

import (
	"context"
	"fmt"
	"strconv"
)

func Confirmations(ctx context.Context, txHeight string, required uint64, latestHeight func(context.Context) (string, error)) (uint64, error) {
	if required == 0 {
		return 0, nil
	}
	height, err := strconv.ParseUint(txHeight, 10, 64)
	if err != nil || height == 0 {
		return 0, fmt.Errorf("invalid tx height %q", txHeight)
	}
	latestStr, err := latestHeight(ctx)
	if err != nil {
		return 0, fmt.Errorf("failed to get latest height: %w", err)
	}
	latest, err := strconv.ParseUint(latestStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid latest height %q", latestStr)
	}
	if latest < height {
		return 0, nil
	}
	return latest - height, nil
}
