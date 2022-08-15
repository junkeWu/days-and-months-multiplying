package utils

import (
	"testing"
)

func TestGetSelfFuncName(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{name: "main", want: "true"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSelfFuncName(); got != tt.want {
				t.Errorf("GetSelfFuncName() = %v, want %v", got, tt.want)
			}
		})
	}
}
