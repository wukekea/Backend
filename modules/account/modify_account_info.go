package account

import (
	"Backend/modules/core/api"
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

type ModifyAccountInfoParam struct {
	MID      string      `json:"id" bson:"id" binding:"required"`    // 账户id
	Name     *string     `json:"name,omitempty" bson:"name"`         // 用户名
	Password *string     `json:"password,omitempty" bson:"password"` // 用户密码
	Age      *int8       `json:"age,omitempty" bson:"age"`           // 用户年龄
	Sex      *int8       `json:"sex" bson:"sex" binding:"required"`  // 用户性别；0：未知；1：男；2：女
	Hobbies  []api.Hobby `json:"hobbies" bson:"hobbies"`             // 用户爱好
	Address  *string     `json:"address,omitempty" bson:"address"`   // 用户地址
}

func ModifyAccountInfo(param *ModifyAccountInfoParam) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(base.MongodbUri))
	if err != nil {
		log.Println("fail to get a client.")
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
		log.Println("fail to connect to database.")
		return err
	}

	filter := bson.M{"id": param.MID}
	updateFilter := bson.M{}
	if param.Name != nil && *param.Name != ""{
		updateFilter["name"] = *param.Name
	}
	if param.Password != nil {
		updateFilter["password"] = *param.Password
	}
	if param.Age != nil {
		updateFilter["age"] = *param.Age
	}
	updateFilter["sex"] = param.Sex
	updateFilter["hobbies"] = param.Hobbies
	if param.Address != nil && *param.Address != "" {
		updateFilter["address"] = param.Address
	}

	//updateInfo := bson.D{{"$set", bson.D{{"name", "wkk"}}}}
	updateInfo := bson.M{"$set": updateFilter}
	updateResult, err := client.Database(base.DatabaseBlog).Collection(base.CollectionAccount).UpdateOne(ctx, filter, updateInfo)
	if err != nil {
		log.Println("fail to update a document from database.")
		return err
	}
	if updateResult.MatchedCount == 0 {
		log.Printf("fail to get the document. matchedCount: %d\n", updateResult.ModifiedCount)
		return errors.New(fmt.Sprintf("fail to get the document. matchedCount: %d\n", updateResult.ModifiedCount))
	}
	return nil
}
