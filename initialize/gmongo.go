package initialize

import (
	"context"
	"gin-demo/global"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Database

func MongoInit() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(global.GVA_CONFIG.MongoDB.Timeout))
	defer cancel()
	o := options.Client().ApplyURI(global.GVA_CONFIG.MongoDB.MongoDBURI)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(uint64(global.GVA_CONFIG.MongoDB.ConnectNum))
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil
	}
	// 返回 client
	MongoClient = client.Database(global.GVA_CONFIG.MongoDB.Database)
	return nil
}

func Collection(name string) *mongo.Collection {
	return MongoClient.Collection(name)
}
