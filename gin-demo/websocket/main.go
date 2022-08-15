package main

import (
	"errors"
	"sync"

	"github.com/gorilla/websocket"
)

// 封装websocket并发读写操作

type Connection struct {
	WsConn    *websocket.Conn
	InChan    chan []byte
	OutChan   chan models.BayDataS
	CloseChan chan byte
	Mutex     sync.Mutex
	IsClosed  bool
}

func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		WsConn:    wsConn,
		InChan:    make(chan []byte, 1000),
		OutChan:   make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	// 读协程
	go conn.ReadLoop()
	// 写协程
	go conn.WriteLoop()
	return
}

func (conn *Connection) ReadMess() (data []byte, err error) {
	select {
	case data = <-conn.InChan:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) WriteMes(data []byte) (err error) {
	select {
	case conn.OutChan <- data:
	case <-conn.CloseChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close() {
	conn.Close() // 本身线程安全，可重入
	// 加锁，只能执行一次
	conn.Mutex.Lock()
	if !conn.IsClosed {
		close(conn.CloseChan)
		conn.IsClosed = true
	}
}

// 具体实现读消息
func (conn *Connection) ReadLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.WsConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.InChan <- data:
		case <-conn.CloseChan:
			goto ERR
		}
	}
ERR:
	conn.Close()
}

// 具体实现写消息
func (conn *Connection) WriteLoop() {
	var (
		data models.BayDataS
		err  error
	)
	for {
		select {
		case data = <-conn.OutChan:
		case <-conn.CloseChan:
			goto ERR
		}
		if err = conn.WsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
