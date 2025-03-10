package models

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2020  Bishop Fox

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
	"time"

	"gorm.io/gorm"
)

// KeyExHistory - Represents an implant
type KeyExHistory struct {
	Sha256    string    `gorm:"primaryKey;"`
	CreatedAt time.Time `gorm:"->;<-:create;"`
}

// BeforeCreate - GORM hook
func (k *KeyExHistory) BeforeCreate(tx *gorm.DB) (err error) {
	k.CreatedAt = time.Now()
	return nil
}
