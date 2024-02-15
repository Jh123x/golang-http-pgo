package main

import "sync"

type SafeMap[Key, Value any] struct {
	safeMap    *sync.Map
	emptyValue Value
}

func (sm *SafeMap[Key, Value]) Load(key Key) Value {
	value, ok := sm.safeMap.Load(key)
	if !ok {
		return sm.emptyValue
	}
	return value.(Value)
}

func (sm *SafeMap[Key, Value]) Store(key Key, value Value) {
	sm.safeMap.Store(key, value)
}

func (sm *SafeMap[Key, Value]) Delete(key Key) {
	sm.safeMap.Delete(key)
}

func NewSafeMap[Key, Value any]() *SafeMap[Key, Value] {
	return &SafeMap[Key, Value]{safeMap: &sync.Map{}}
}
