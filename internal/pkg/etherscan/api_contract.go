package etherscan

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

type GetContractCreationArg struct {
	ContractAddresses []string `json:"contract_addresses"`
}
type GetContractCreationResult []struct {
	ContractAddress string `json:"contractAddress"`
	ContractCreator string `json:"contractCreator"`
	TxHash          string `json:"txHash"`
}

func (c *Client) GetContractCreation(ctx context.Context, arg *GetContractCreationArg) (GetContractCreationResult, error) {
	resp, err := c.rc.R().SetQueryParams(map[string]string{
		"module":            "contract",
		"action":            "getcontractcreation",
		"contractaddresses": strings.Join(arg.ContractAddresses, ","),
	}).Get("/api")

	if err != nil {
		return nil, errors.Wrapf(err, "failed to get contract creation")
	}

	if gjson.GetBytes(resp.Body(), "status").String() != "1" {
		return nil, errors.Errorf("%s", gjson.GetBytes(resp.Body(), "result").String())
	}

	var result GetContractCreationResult
	if err := json.Unmarshal(
		[]byte(gjson.GetBytes(resp.Body(), "result").Raw),
		&result,
	); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal contract creation response")
	}

	return result, nil
}
