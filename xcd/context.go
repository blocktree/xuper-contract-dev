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
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
)

type (
	EventKey   string
	EncodeType int
)

const (
	DefaultEventKey = EventKey("com.github.blocktree.xcd.event")
)

const (
	_ EncodeType = iota
	EncodeTypeJSON
	EncodeTypeRLP
)

type Event interface {
	EventName(e Event) string
}

type EventBase struct{}

func (base EventBase) EventName(e Event) string {
	return reflect.TypeOf(e).Name()
}

type EventLog struct {
	Event string      `json:"event"`
	Value interface{} `json:"value"`
}

type Context struct {
	code.Context

	keyName    EventKey
	encodeType EncodeType
}

// NewContext
func NewContext(ctx code.Context, args ...interface{}) *Context {
	var (
		keyName    = DefaultEventKey
		encodeType = EncodeTypeJSON
	)

	store, ok := ctx.(*Context)
	if ok {
		return store
	}

	for _, arg := range args {
		switch obj := arg.(type) {
		case EventKey:
			keyName = obj
		case EncodeType:
			encodeType = obj
		}
	}

	return newContext(ctx, keyName, encodeType)
}

// newContext
func newContext(ctx code.Context, keyName EventKey, encodeType EncodeType) *Context {
	store := &Context{
		Context:    ctx,
		keyName:    keyName,
		encodeType: encodeType,
	}

	//清楚旧值
	store.DeleteObject([]byte(keyName))
	return store
}

func (ctx *Context) Emit(event Event) error {
	name := event.EventName(event)
	eventLog := EventLog{
		Event: name,
		Value: event,
	}

	eventLogs, err := ctx.getEventLogs()
	if err != nil {
		return err
	}

	eventLogs = append(eventLogs, eventLog)
	err = ctx.setEventLogs(eventLogs)
	if err != nil {
		return err
	}

	return nil
}

// getEventLogs
func (ctx *Context) getEventLogs() ([]EventLog, error) {
	logBytes, err := ctx.GetObject([]byte(ctx.keyName))
	if err != nil || logBytes == nil {
		return []EventLog{}, nil
	}

	var logs []EventLog
	switch ctx.encodeType {
	case EncodeTypeJSON:

		err = json.Unmarshal(logBytes, &logs)
		if err != nil {
			return nil, err
		}
	case EncodeTypeRLP:
		return []EventLog{}, nil
	}
	return logs, nil
}

// setEventLogs
func (ctx *Context) setEventLogs(logs []EventLog) error {

	var (
		logBytes []byte
		err      error
	)
	switch ctx.encodeType {
	case EncodeTypeJSON:
		logBytes, err = json.Marshal(logs)
	case EncodeTypeRLP:
		err = fmt.Errorf("EncodeTypeRLP is not implemented")
	}
	if err != nil {
		return err
	}
	return ctx.PutObject([]byte(ctx.keyName), logBytes)
}

func (ctx *Context) SetInt(key string, value *big.Int) error {
	return ctx.PutObject([]byte(key), value.Bytes())
}

func (ctx *Context) SetBool(key string, value bool) error {
	var b byte
	if value {
		b = 0x01
	} else {
		b = 0x00
	}
	return ctx.PutObject([]byte(key), []byte{b})
}

func (ctx *Context) SetString(key string, value string) error {
	return ctx.PutObject([]byte(key), []byte(value))
}

func (ctx *Context) SetStrings(key string, value []string) error {
	vals := strings.Join(value, "*;*")
	return ctx.PutObject([]byte(key), []byte(vals))
}

func (ctx *Context) SetBytes(key string, value []byte) error {
	return ctx.PutObject([]byte(key), value)
}

func (ctx *Context) GetInt(key string, def *big.Int) *big.Int {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return def
	}
	intv := new(big.Int)
	intv.SetBytes(value)
	return intv
}

func (ctx *Context) GetBool(key string, def bool) bool {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return def
	}
	bInt := new(big.Int)
	bInt.SetBytes(value)
	if bInt.Cmp(big.NewInt(0)) > 0 {
		return true
	}
	return false
}

func (ctx *Context) GetString(key string, def string) string {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return def
	}
	return string(value)
}

func (ctx *Context) GetStrings(key string, def []string) []string {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return def
	}
	return strings.Split(string(value), "*;*")
}

func (ctx *Context) GetBytes(key string, def []byte) []byte {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return def
	}
	return value
}

func (ctx *Context) GetJSON(key string) *gjson.Result {
	value, err := ctx.GetObject([]byte(key))
	if err != nil {
		return nil
	}
	result := gjson.ParseBytes(value)
	return &result
}

func (ctx *Context) ArgToInt(key string) *big.Int {
	if arg := ctx.Args()[key]; arg != nil {
		intv := new(big.Int)
		intv.SetBytes(arg)
		return intv
	}
	return big.NewInt(0)
}

func (ctx *Context) ArgToBool(key string) bool {
	if arg := ctx.Args()[key]; arg != nil {
		bInt := new(big.Int)
		bInt.SetBytes(arg)
		if bInt.Cmp(big.NewInt(0)) > 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

func (ctx *Context) ArgToString(key string) string {
	if arg := ctx.Args()[key]; arg != nil {
		return string(arg)
	}
	return ""
}

func (ctx *Context) ArgToStrings(key string) []string {
	if arg := ctx.Args()[key]; arg != nil {
		return strings.Split(string(arg), "*;*")
	}
	return nil
}

func (ctx *Context) ArgToBytes(key string) []byte {
	if arg := ctx.Args()[key]; arg != nil {
		return arg
	}
	return nil
}

func (ctx *Context) ArgToJSON(key string) *gjson.Result {
	if arg := ctx.Args()[key]; arg != nil {
		result := gjson.ParseBytes(arg)
		return &result
	}
	return nil
}
