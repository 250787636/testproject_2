package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo" //MongoDB的Go驱动包
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

type com struct {
	Articleid      string    `json:"articleid"`
	Content        string    `json:"content"`
	Userid         string    `json:"userid"`
	Nickname       string    `json:"nickname"`
	Createdatetime time.Time `json:"createdatetime"`
	Likenum        int       `json:"likenum"`
	State          string    `json:"state"`
}

func Test_Mongo(t *testing.T) {

	// 链接地址
	url := fmt.Sprintf("mongodb://localhost:27017")
	clientOPtions := options.Client().ApplyURI(url)

	// 建立客户端连接
	client, err := mongo.Connect(context.TODO(), clientOPtions)
	if err != nil {
		fmt.Println(err)
	}

	// 检查连接情况
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("connect right")

	collection := client.Database("test").Collection("comment")

	get(1, client, collection)

	go func() {
		err = client.Disconnect(context.TODO())
		if err != nil {
			fmt.Println(err)
		}
	}()
}

func get(num int, client *mongo.Client, collection *mongo.Collection) {
	switch num {
	case 1:
		id := insertSensor(client, collection)
		fmt.Println(id)
	case 2:
		c := findSensor(client, collection)
		fmt.Println(c)
	case 3:
		isSuccess := delSensor(client, collection)
		if isSuccess {
			fmt.Println("delSuccess")
		} else {
			fmt.Println("delFail")
		}
	case 4:
		isSuccess := updateSensor(client, collection)
		if isSuccess {
			fmt.Println("update success")
		} else {
			fmt.Println("update fail")
		}
	}
}

// 插入数据
func insertSensor(client *mongo.Client, collection *mongo.Collection) (id primitive.ObjectID) {
	result, err := collection.InsertOne(context.TODO(), &com{
		Articleid:      "100002",
		Content:        "这个不错",
		Userid:         "1010",
		Nickname:       "黑蛋",
		Createdatetime: time.Now(),
		Likenum:        2,
		State:          "1",
	})
	if err != nil {
		fmt.Println(err)
	}
	id = result.InsertedID.(primitive.ObjectID)
	return id
}

// 删除数据
func delSensor(client *mongo.Client, collection *mongo.Collection) bool {

	filter := bson.M{
		"_id": "4",
	}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 更新数据
func updateSensor(client *mongo.Client, collection *mongo.Collection) bool {
	filter := bson.M{
		"_id": "3",
	}
	updateData := bson.M{
		"$set": bson.M{
			"articleid":      "100002",
			"content":        "不错不错不错",
			"userid":         "1005",
			"nickname":       "shi蛋shi蛋",
			"createdatetime": time.Now(),
			"likenum":        4,
			"state":          "1",
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 查询数据
func findSensor(client *mongo.Client, collection *mongo.Collection) (c com) {
	filter := bson.M{
		"_id": "4",
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
