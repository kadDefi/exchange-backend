package ethereum

import (
	"context"
	pkgabi "exchange-backend/internal/pkg/abi"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
)

type multiCallInstance struct {
	client *Client

	methods []abi.Method
	calls   []pkgabi.Struct0
	err     error
}

func (c *Client) NewMultiCall() *multiCallInstance {
	return &multiCallInstance{
		client:  c,
		methods: []abi.Method{},
		calls:   []pkgabi.Struct0{},
	}
}

func (i *multiCallInstance) Add(
	target common.Address,
	method abi.Method,
	args ...interface{},
) *multiCallInstance {
	if i.err != nil {
		return i
	}

	if method.ID == nil {
		i.err = errors.Errorf("call %d: method ID is nil", len(i.methods))
		return i
	}

	call, err := method.Inputs.Pack(args...)
	if err != nil {
		i.err = errors.Wrapf(err, "call %d - %s: failed to pack call", len(i.methods), method.Name)
		return i
	}

	i.methods = append(i.methods, method)
	i.calls = append(i.calls, pkgabi.Struct0{
		Target:   target,
		CallData: append(method.ID, call...),
	})

	return i
}

func (i *multiCallInstance) Call(
	ctx context.Context,
	opts *bind.CallOpts,
) (*big.Int, []interface{}, error) {
	if i.err != nil {
		return nil, nil, i.err
	}

	if opts == nil {
		opts = &bind.CallOpts{}
	}
	if opts.Context == nil {
		opts.Context = ctx
	}

	aggregate := new([]interface{})
	if err := i.client.multiCallCallerRaw.Call(opts, aggregate, "aggregate", i.calls); err != nil {
		return nil, nil, errors.Wrapf(err, "failed to call aggregate")
	}

	blockNumber := (*aggregate)[0].(*big.Int)
	returnData := (*aggregate)[1].([][]byte)

	results := make([]interface{}, 0, len(i.methods))
	for i, method := range i.methods {
		data := returnData[i]
		output, err := method.Outputs.Unpack(data)
		if err != nil {
			return nil, nil, errors.Wrapf(err, "call %d - %s: failed to unpack return data", i, method.Name)
		}

		results = append(results, output)
	}

	return blockNumber, results, nil
}
