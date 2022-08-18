package main

import (
	"fmt"
	"strings"
)

// BackendServer 每一个后端服务定义
type BackendServer struct {
	// 实例权重
	Weight int
	// 当前的权重，初始为Weight
	currentWeight int
	// 后端服务名称
	ServerName string
}

// WeightServerRoundRobin 通过权重实现调用轮询的定义
type WeightServerRoundRobin struct {
	// 所有有效的权重总和
	effectiveWeight int
	// 后端服务列表
	backendServerList []*BackendServer
}

// NewWeightServerRoundRobin 创建一个负载轮询器
func NewWeightServerRoundRobin() *WeightServerRoundRobin {
	return &WeightServerRoundRobin{
		effectiveWeight: 0,
	}
}

// AddBackendServer 增加后端服务名称和权重
func (r *WeightServerRoundRobin) AddBackendServer(backendServer *BackendServer) {
	r.effectiveWeight += backendServer.Weight
	r.backendServerList = append(r.backendServerList, backendServer)
}

// GetBackendServer 更具权重获取一个后端服务名称
func (r *WeightServerRoundRobin) GetBackendServer() *BackendServer {
	var expectBackendServer *BackendServer
	for _, backendServer := range r.backendServerList {
		// 给每个后端服务增加自身权重
		backendServer.currentWeight += backendServer.Weight
		if expectBackendServer == nil {
			expectBackendServer = backendServer
		}
		if backendServer.currentWeight > expectBackendServer.currentWeight {
			expectBackendServer = backendServer
		}
	}
	r.VisitBackendServerCurrentWeight()
	// 把选择的后端服务权重减掉总权重
	expectBackendServer.currentWeight -= r.effectiveWeight
	return expectBackendServer
}

// VisitBackendServerCurrentWeight 打印后端服务的当前权重变化
func (r *WeightServerRoundRobin) VisitBackendServerCurrentWeight() {
	var serverListForLog []string
	for _, backendServer := range r.backendServerList {
		serverListForLog = append(serverListForLog,
			fmt.Sprintf("%v", backendServer.currentWeight))
	}
	fmt.Printf("(%v)\n", strings.Join(serverListForLog, ", "))
}
