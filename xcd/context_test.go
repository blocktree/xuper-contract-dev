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
	"fmt"
	"github.com/xuperchain/xuperchain/core/contractsdk/go/code"
	"math/big"
	"testing"
)

type Issue struct {
	EventBase
	To     string   `json:"to"`
	Amount *big.Int `json:"amount"`
}

type Burn struct {
	EventBase
	From   string   `json:"from"`
	Amount *big.Int `json:"amount"`
}

type Transfer struct {
	EventBase
	From   string   `json:"from"`
	To     string   `json:"to"`
	Amount *big.Int `json:"amount"`
}

func transfer(ctx code.Context) error {
	eventCTX := NewContext(ctx)
	return eventCTX.Emit(Transfer{
		From:   "gogogo",
		To:     "heihei",
		Amount: big.NewInt(1234456),
	})
}

func TestEventContext_Method(t *testing.T) {
	store := NewSimulateContext()
	err := transfer(store)
	if err != nil {
		t.Errorf("transfer failed: %v", err)
		return
	}

}

func TestEventContext_Emit(t *testing.T) {
	store := NewContext(NewSimulateContext())
	//store := NewContext(NewSimulateContext(), EventKey("event"), EncodeTypeJSON)
	err := store.Emit(Issue{To: "qwer", Amount: big.NewInt(100)})
	if err != nil {
		t.Errorf("Emit failed: %v", err)
		return
	}

	err = store.Emit(Issue{To: "22222", Amount: big.NewInt(2222)})
	if err != nil {
		t.Errorf("Emit failed: %v", err)
		return
	}

	err = store.Emit(Burn{From: "22222", Amount: big.NewInt(2222)})
	if err != nil {
		t.Errorf("Emit failed: %v", err)
		return
	}

	err = transfer(store)
	if err != nil {
		t.Errorf("transfer failed: %v", err)
		return
	}
}

func TestContext_SetAndGet(t *testing.T) {

	ctx := NewContext(NewSimulateContext())

	/******* Test Set *******/

	if err := ctx.SetInt("keyInt", big.NewInt(11)); err != nil {
		t.Errorf("SetInt() error = %v", err)
	}

	if err := ctx.SetBytes("keyBytes", []byte{0x01, 0x02, 0x03}); err != nil {
		t.Errorf("SetBytes() error = %v", err)
	}

	if err := ctx.SetBool("keyBool", true); err != nil {
		t.Errorf("SetBool() error = %v", err)
	}

	if err := ctx.SetString("keyString", "hello"); err != nil {
		t.Errorf("SetString() error = %v", err)
	}

	if err := ctx.SetStrings("keyStrings", []string{"hello", "world"}); err != nil {
		t.Errorf("SetStrings() error = %v", err)
	}

	/******* Test Get *******/

	vInt := ctx.GetInt("keyInt", big.NewInt(0))
	fmt.Printf("GetInt = %v \n", vInt)

	vBytes := ctx.GetBytes("keyBytes", nil)
	fmt.Printf("GetBytes = %v \n", vBytes)

	vBool := ctx.GetBool("keyBool", false)
	fmt.Printf("GetBool = %v \n", vBool)

	vString := ctx.GetString("keyString", "")
	fmt.Printf("GetString = %v \n", vString)

	vStrings := ctx.GetStrings("keyStrings", nil)
	fmt.Printf("GetStrings = %v \n", vStrings)

	/******* Test Get Default Value *******/

	vInt2 := ctx.GetInt("keyInt2", big.NewInt(0))
	fmt.Printf("GetInt2 = %v \n", vInt2)

	vBytes2 := ctx.GetBytes("keyBytes2", nil)
	fmt.Printf("GetBytes2 = %v \n", vBytes2)

	vBool2 := ctx.GetBool("keyBool2", false)
	fmt.Printf("GetBool2 = %v \n", vBool2)

	vString2 := ctx.GetString("keyString2", "")
	fmt.Printf("GetString2 = %v \n", vString2)

	vStrings2 := ctx.GetStrings("keyStrings2", nil)
	fmt.Printf("GetStrings2 = %v \n", vStrings2)

}
