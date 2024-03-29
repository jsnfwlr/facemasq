package events

import (
	"sort"
	"sync"
)

// listenerMap wraps sync.Map and uses the following types:
// key:   string
// value: Listener
type listenerMap struct {
	data sync.Map
}

// Delete the key from the map
func (m *listenerMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns Listener or bool.
// A false return indicates either the key was not found
// or the value is not of type Listener
func (m *listenerMap) Load(key string) (Listener, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return nil, false
	}
	s, ok := i.(Listener)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *listenerMap) LoadOrStore(key string, value Listener) (Listener, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.(Listener)
	return s, ok
}

// Range over the Listener values in the map
func (m *listenerMap) Range(f func(key string, value Listener) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(Listener)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a Listener in the map
func (m *listenerMap) Store(key string, value Listener) {
	m.data.Store(key, value)
}

// Keys returns a list of keys in the map
func (m *listenerMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value Listener) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
