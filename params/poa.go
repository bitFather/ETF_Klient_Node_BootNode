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

package params

import (
	"math/big"

	"github.com/Exgibichi/go-etf/common"
)

// POA fork params
var (
	POAForkBlockExtra          = common.FromHex("0x796f62610000000000000000000000000000000000000000000000000000000001a4260e6348ddd60ee1a597d1c6a102a076e9040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	POAForkExtraRange          = big.NewInt(10)
	POAConfig                  = &CliqueConfig{Period: 5, Epoch: 30000}
	POAForkGasLimit   uint64   = 105e+6
	POAChainId        *big.Int = big.NewInt(455446)
	POAFoundersReward *big.Int = new(big.Int).Mul(big.NewInt(1e+18), big.NewInt(10e+6))
)

// POA Founders accounts
func POA() []common.Address {
	return []common.Address{
		common.HexToAddress("0x01a4260e6348ddd60ee1a597d1c6a102a076e904"),
		common.HexToAddress("0x13678c1285ff6b9b84e566d2e374e7d1be45fc8f"),
		common.HexToAddress("0x939f2f2fc4f1650a28733e3e52d844be12ce7fa7"),
		common.HexToAddress("0x983e7d95b9629ee41e44592fba8ebdf72bb5260c"),
		common.HexToAddress("0x69b565fdacb27c7ae652e3aa13322347e5398630"),
		common.HexToAddress("0xe939ec7525f8787f8c0ec4eef25be6e2b5dd34f0"),
		common.HexToAddress("0x8986578a8d9951a4B55038B01262d81b2a030D33"),
	}
}
