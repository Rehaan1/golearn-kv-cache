package kvcache

import (
	"bytes"
	"testing"
)

func TestCache(t *testing.T) {
	tests := []struct {
		name   string
		setups []struct {
			key   string
			value []byte
		}
		getKey    string
		wantValue []byte
		wantOk    bool
	}{
		{
			name:      "get on empty cache returns miss",
			setups:    nil,
			getKey:    "missing",
			wantValue: nil,
			wantOk:    false,
		},
		{
			name: "set then get same key returns hit",
			setups: []struct {
				key   string
				value []byte
			}{
				{key: "a", value: []byte("hello")},
			},
			getKey:    "a",
			wantValue: []byte("hello"),
			wantOk:    true,
		},
		{
			name: "set same key twice, latest value wins",
			setups: []struct {
				key   string
				value []byte
			}{
				{key: "a", value: []byte("first")},
				{key: "a", value: []byte("second")},
			},
			getKey:    "a",
			wantValue: []byte("second"),
			wantOk:    true,
		},
		{
			name: "set empty value is still a hit",
			setups: []struct {
				key   string
				value []byte
			}{
				{key: "empty", value: []byte{}},
			},
			getKey:    "empty",
			wantValue: []byte{},
			wantOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Cache

			for _, s := range tt.setups {
				c.Set(s.key, s.value)
			}

			gotValue, gotOk := c.Get(tt.getKey)

			if gotOk != tt.wantOk {
				t.Errorf("Get(%q) ok = %v, want %v", tt.getKey, gotOk, tt.wantOk)
			}
			if !bytes.Equal(gotValue, tt.wantValue) {
				t.Errorf("Get(%q) value = %v, want %v", tt.getKey, gotValue, tt.wantValue)
			}
		})
	}
}
