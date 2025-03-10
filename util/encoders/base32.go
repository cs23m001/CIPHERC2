package encoders

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"encoding/base32"
)

// Base32 Encoder
type Base32 struct{}

// Missing chars: s, i, l, o
const base32Alphabet = "ab1c2d3e4f5g6h7j8k9m0npqrtuvwxyz"

var CIPHERC2Base32 = base32.NewEncoding(base32Alphabet).WithPadding(base32.NoPadding)

// Encode - Base32 Encode
func (e Base32) Encode(data []byte) ([]byte, error) {
	return []byte(CIPHERC2Base32.EncodeToString(data)), nil
}

// Decode - Base32 Decode
func (e Base32) Decode(data []byte) ([]byte, error) {
	return CIPHERC2Base32.DecodeString(string(data))
}
