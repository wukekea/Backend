package account

import (
	"Backend/modules/core/base"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type LogoffParam struct {
	Name     string `json:"name" binding:"required" example:"wkk"`
	Password string `json:"password" binding:"required" example:"123456789"`
}

func Logoff(param *LogoffParam) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(base.MongodbUri))
	if err != nil {
		log.Println("fail to get a new client")
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Connect(ctx)
	if err != nil {
		log.Println("fail to client to a database.")
		return err
	}

	filter := bson.M{"name": param.Name, "password": param.Password}
	deleteResult, err := client.Database(base.DatabaseBlog).Collection(base.CollectionAccount).DeleteOne(ctx, filter)
	if err != nil {
		log.Println("fail to delete an account.")
		return err
	}
	if deleteResult.DeletedCount != 1 {
		log.Printf("DeletedCount != 1. DeletedCount: %d", deleteResult.DeletedCount)
		return errors.New(fmt.Sprintf("DeleteCount != 1. DeleteCount: %d", deleteResult.DeletedCount))
	}
	return nil
}
