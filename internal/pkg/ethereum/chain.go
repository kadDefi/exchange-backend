package ethereum

import (
	"context"
	"math/big"
)

func (c *Client) ChainID(ctx context.Context) (*big.Int, error) {
	if c.chainID == nil {
		id, err := c.Client.ChainID(ctx)
		if err != nil {
			return nil, err
		}

		c.chainID = id
	}

	return c.chainID, nil
}
