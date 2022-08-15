package utils

import "testing"

func TestGetLocalIP(t *testing.T) {
	tests := []struct {
		name        string
		want        string
		wantErr     bool
		isErrorCase bool
	}{
		// TODO: Add test cases.
		// 需要修改ip
		{name: "测试服务器ip", want: "192.168.2.149", wantErr: false, isErrorCase: false},
		{name: "测试服务器ip", want: "192.168.2.148", wantErr: false, isErrorCase: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLocalIP()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocalIP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.isErrorCase {
				t.Errorf("GetLocalIP() got = %v, want %v", got, tt.want)
			}
		})
	}
}
