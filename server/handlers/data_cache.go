package handlers

/*
	CIPHERC2 Implant Framework
	Copyright (C) 2022  Bishop Fox

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
	"sync"

	"github.com/cs23m001/CIPHERC2/protobuf/CIPHERC2pb"
)

var (
	// TunnelID -> Sequence Number -> Data
	tunnelDataCache = dataCache{mutex: &sync.RWMutex{}, cache: map[uint64]map[uint64]*CIPHERC2pb.TunnelData{}}
)

type dataCache struct {
	mutex *sync.RWMutex
	cache map[uint64]map[uint64]*CIPHERC2pb.TunnelData
}

func (c *dataCache) Add(tunnelID uint64, sequence uint64, tunnelData *CIPHERC2pb.TunnelData) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.cache[tunnelID]; !ok {
		c.cache[tunnelID] = map[uint64]*CIPHERC2pb.TunnelData{}
	}

	c.cache[tunnelID][sequence] = tunnelData
}

func (c *dataCache) Get(tunnelID uint64, sequence uint64) (*CIPHERC2pb.TunnelData, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if _, ok := c.cache[tunnelID]; !ok {
		return nil, false
	}

	val, ok := c.cache[tunnelID][sequence]

	return val, ok
}

func (c *dataCache) DeleteTun(tunnelID uint64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.cache, tunnelID)
}

func (c *dataCache) DeleteSeq(tunnelID uint64, sequence uint64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.cache[tunnelID]; !ok {
		return
	}

	delete(c.cache[tunnelID], sequence)
}

func (c *dataCache) Len(tunnelID uint64) int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return len(c.cache[tunnelID])
}
