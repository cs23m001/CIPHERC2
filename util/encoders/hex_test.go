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
	"bytes"
	"testing"
)

func TestHex(t *testing.T) {
	sample := randomData()

	// Server-side
	x := new(Hex)
	output, _ := x.Encode(sample)
	data, err := x.Decode(output)
	if err != nil {
		t.Errorf("hex decode returned an error %v", err)
	}
	if !bytes.Equal(sample, data) {
		t.Errorf("sample does not match returned\n%#v != %#v", sample, data)
	}

}
