package account

import (
	"Backend/modules/core/api"
	"Backend/modules/core/base"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type RegisterParam struct {
	Name     string `json:"name" binding:"required" example:"wkk"`
	Password string `json:"password" binding:"required" example:"123456789"`

	Age     int8        `json:"age" example:"20"`        // 用户年龄
	Sex     int8        `json:"sex" example:"1"`         // 用户性别；0：未知；1：男；2：女
	Hobbies []api.Hobby `json:"hobbies"`                 // 用户爱好
	Address string      `json:"address" example:"China"` // 用户地址
}

func Register(param RegisterParam) (ID interface{}, err error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(base.MongodbUri))
	if err != nil {
		log.Println("fail to get a new client.")
		return
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
		return
	}

	collection := client.Database(base.DatabaseBlog).Collection(base.CollectionAccount)
	accountInfo := api.AccountInfo{
		ID:       primitive.NewObjectID(),
		MID:      "1",
		Name:     param.Name,
		Password: param.Password,
		Age:      param.Age,
		Sex:      param.Sex,
		Hobbies:  param.Hobbies,
		Address:  param.Address,
	}
	res, err := collection.InsertOne(ctx, accountInfo)
	if err != nil {
		log.Println("fail to insert an account information to database.")
		return
	}
	ID = res.InsertedID
	return
}
