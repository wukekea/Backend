package account

import (
	"Backend/modules/core/api"
	"Backend/modules/core/base"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type LoginParam struct {
	Name     string `json:"name" binding:"required" example:"wkk"`
	Password string `json:"password" binding:"required" example:"12345679"`
}

func Login(param LoginParam) (*api.AccountInfo, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(base.MongodbUri))
	if err != nil {
		log.Println("fail to get a new client.")
		return nil, err
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
		log.Println("fail to connect to a client.")
		return nil, err
	}

	filter := bson.M{"name": param.Name, "password": param.Password}
	collection := client.Database(base.DatabaseBlog).Collection(base.CollectionAccount)
	singleResult := collection.FindOne(ctx, filter)
	result := api.AccountInfo{}
	err = singleResult.Decode(&result)
	return &result, err
}
