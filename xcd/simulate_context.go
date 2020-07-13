/*
 * Copyright 2019 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package xcd

import (
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/pb"
	"math/big"
)

type SimulateContext struct {
	store map[string][]byte
}

func NewSimulateContext() *SimulateContext {
	s := &SimulateContext{store: make(map[string][]byte)}
	return s
}

func (ctx *SimulateContext) Args() map[string][]byte {
	return nil
}
func (ctx *SimulateContext) Caller() string {
	return ""
}
func (ctx *SimulateContext) Initiator() string {
	return ""
}
func (ctx *SimulateContext) AuthRequire() []string {
	return nil
}

func (ctx *SimulateContext) PutObject(key []byte, value []byte) error {
	ctx.store[string(key)] = value
	return nil
}
func (ctx *SimulateContext) GetObject(key []byte) ([]byte, error) {
	return ctx.store[string(key)], nil
}
func (ctx *SimulateContext) DeleteObject(key []byte) error {
	delete(ctx.store, string(key))
	return nil
}
func (ctx *SimulateContext) NewIterator(start, limit []byte) code.Iterator {
	return nil
}

func (ctx *SimulateContext) QueryTx(txid string) (*pb.Transaction, error) {
	return nil, nil
}
func (ctx *SimulateContext) QueryBlock(blockid string) (*pb.Block, error) {
	return nil, nil
}
func (ctx *SimulateContext) Transfer(to string, amount *big.Int) error {
	return nil
}
func (ctx *SimulateContext) TransferAmount() (*big.Int, error) {
	return nil, nil
}
func (ctx *SimulateContext) Call(module, contract, method string, args map[string][]byte) (*code.Response, error) {
	return nil, nil
}
func (ctx *SimulateContext) CrossQuery(uri string, args map[string][]byte) (*code.Response, error) {
	return nil, nil
}

func (ctx *SimulateContext) Logf(fmt string, args ...interface{}) {

}
