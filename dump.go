package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
		SetAuth(options.Credential{
			AuthSource: "dbinfo", Username: "appuser", Password: "Mko0Zaq1",
		})

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var contentArray []interface{}

	contentArray = append(contentArray, bson.M{"Title": "DevOps - Alpha", "Description": "#01 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Bravo", "Description": "#02 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Charlie", "Description": "#03 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Delta", "Description": "#04 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Echo", "Description": "#05 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Fhox", "Description": "#06 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Golf", "Description": "#07 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Hotel", "Description": "#08 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - India", "Description": "#09 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Juliet", "Description": "#10 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Kilo", "Description": "#11 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Lima", "Description": "#12 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Mike", "Description": "#13 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - November", "Description": "#14 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Oscar", "Description": "#15 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Papa", "Description": "#16 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Quebe", "Description": "#17 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Romeo", "Description": "#18 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Sierra", "Description": "#19 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Tango", "Description": "#20 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Uniform", "Description": "#21 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Victor", "Description": "#22 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Whiskey", "Description": "#23 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - X-ray", "Description": "#24 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Yankee", "Description": "#25 Example by Anderson Braz."})
	contentArray = append(contentArray, bson.M{"Title": "DevOps - Zulu", "Description": "#26 Example by Anderson Braz."})

	collection := client.Database("dbinfo").Collection("infos")

	if res, err := collection.InsertMany(context.TODO(), contentArray); err == nil {
		log.Println("Inserted record", res.InsertedIDs)
	} else {
		log.Fatal("Error inserting record:", err)
	}

}
