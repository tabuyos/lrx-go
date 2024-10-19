// Copyright 2024 tabuyos. All rights reserved.
//
// @author tabuyos
// @since 2024/10/11
// @description file desc
package map_test

import (
	"sync"
	"testing"
)

func NewSyncMap() *sync.Map {
	return &sync.Map{}
}

func Test_Sync_Map(t *testing.T) {
	nodeMap := NewSyncMap()

	// AddTicketEvent Func
	_ = func(id, code, name string, ch chan any) {
		key := code + name
		v, _ := nodeMap.LoadOrStore(id, NewSyncMap())
		if chMap, ok := v.(*sync.Map); ok {
			chMap.Store(key, ch)
		}
	}

	_ = func() {
		nodeMap.Range(func(_, chMap any) bool {
			// for xxx
			return true
		})
	}
}
