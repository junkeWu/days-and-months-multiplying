package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mog *mongo.Client

func init() {
	var err error
	// 连接client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:37017")

	// 到MongoDB
	mog, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// check link
	err = mog.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to mongodb")
}
