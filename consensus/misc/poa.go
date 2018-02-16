// Copyright 2018 The go-etf Authors
// This file is part of the go-etf library.
//
// The go-etf library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etf library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etf library. If not, see <http://www.gnu.org/licenses/>.

package misc

import (
	"bytes"
	"errors"

	"github.com/Exgibichi/go-etf/core/state"
	"github.com/Exgibichi/go-etf/core/types"
	"github.com/Exgibichi/go-etf/params"
)

var (
	ErrBadPOAExtra = errors.New("bad POA fork extra-data")
)

// check POA fork extra data
func VerifyPOAHeaderExtraData(config *params.ChainConfig, header *types.Header) error {
	if config.POABlock == nil {
		return nil
	}
	if config.POABlock != nil && config.POABlock.Cmp(header.Number) != 0 {
		return nil
	}
	if !bytes.Equal(header.Extra, params.POAForkBlockExtra) {
		return ErrBadPOAExtra
	}
	return nil
}

// ApplyDAOHardFork modifies the state database
// add founders accounts
func ApplyPOAHardFork(statedb *state.StateDB) {
	if statedb != nil {
		for _, addr := range params.POA() {
			if !statedb.Exist(addr) {
				statedb.CreateAccount(addr)
				statedb.SetBalance(addr, params.POAFoundersReward)
			}
		}
	}
}
