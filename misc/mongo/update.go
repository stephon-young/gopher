package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB(server string) (*mongo.Client, error) {
	// uri format: "mongodb://localhost:27017"
	uri := fmt.Sprintf("mongodb://%s", server)
	opts := options.Client().ApplyURI(uri)

	// 连接
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("connect error: %v", err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("ping error: %v", err)
	}

	return client, nil

	// 操作数据库的数据集
	//collection := client.Database("test").Collection("trainers")
}

type Template struct {
	ID       uint            `json:"id"`
	Sections json.RawMessage `json:"sections"`
}

type Courseware struct {
	Course   uint        `json:"id"`
	Version  uint        `json:"version"`
	Template Template    `json:"template"`
	Entity   interface{} `json:"entity,omitempty"`
}

type Meta struct {
	Course  uint `json:"course"`
	Version uint `json:"version"`
	BasedOn uint `json:"based_on"`
}

type Owner struct {
	Owner   string `json:"owner"`
	Corp    uint   `json:"corp"`
	Project uint   `json:"project"`
}

type CoursewareSpec struct {
	Owner  Owner       `json:"owner"`
	Meta   Meta        `json:"meta"`
	Entity interface{} `json:"entity"`
}

func main() {
	server := "localhost:27017"
	db, err := NewDB(server)
	if err != nil {
		fmt.Printf("error on NewDB: %v\n", err)
		return
	}

	f, err := ioutil.ReadFile("courseware.json")
	if err != nil {
		fmt.Println("read fail", err)
		return
	}

	v := map[string]interface{}{}

	err = json.Unmarshal(f, &v)
	if err != nil {
		fmt.Println("unmarshal fail", err)
		return
	}

	fmt.Printf("json: \n%v\n", v)

	spec := CoursewareSpec{
		Entity: v,
	}
	/*
		data, err := bson.Marshal(&spec)
		if err != nil {
			fmt.Printf("bson Marshal error: %v", err)
			return
		}
	*/

	collection := db.Database("courseware").Collection("courseware")

	resp, err := collection.InsertOne(context.TODO(), &spec)
	if err != nil {
		fmt.Printf("Add failed: %v", err)
		return
	}

	id := fmt.Sprintf("%s", resp.InsertedID)

	id = strings.Replace(id, "ObjectID(\"", "", -1)
	id = strings.Replace(id, "\")", "", -1)

	spec.Owner.Owner = "aux"

	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}
	update := bson.D{
		{"$set", bson.D{
			{"owner", spec.Owner},
		}},
	}

	//filter := bson.D{"_id": id}
	r, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Printf("update failed: %v", err)
		return
	}

	fmt.Printf("update succeed: %v", r)
}
