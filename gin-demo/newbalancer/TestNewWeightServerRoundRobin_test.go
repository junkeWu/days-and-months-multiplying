package main

import (
	"fmt"
	"testing"
)

func TestNewWeightServerRoundRobin(t *testing.T) {
	weightServerRoundRobin := NewWeightServerRoundRobin()
	weightServerRoundRobin.AddBackendServer(&BackendServer{
		ServerName: "ServerA",
		Weight:     5,
	})
	weightServerRoundRobin.AddBackendServer(&BackendServer{
		ServerName: "ServerB",
		Weight:     3,
	})
	weightServerRoundRobin.AddBackendServer(&BackendServer{
		ServerName: "ServerC",
		Weight:     1,
	})

	expectServerNameList := []string{
		"ServerA", "ServerB", "ServerA", "ServerC", "ServerA", "ServerB", "ServerA", "ServerB", "ServerA",
		// "ServerA", "ServerB", "ServerA", "ServerC", "ServerA", "ServerB", "ServerA", "ServerB", "ServerA",
	}
	fmt.Printf("(A, B, C)\n")
	for ii, expectServerName := range expectServerNameList {
		weightServerRoundRobin.VisitBackendServerCurrentWeight()
		backendServer := weightServerRoundRobin.GetBackendServer()
		if backendServer.ServerName != expectServerName {
			t.Errorf("%v.%v.expect:%v, actual:%v", t.Name(), ii, expectServerName, backendServer.ServerName)
			return
		}
	}
}
