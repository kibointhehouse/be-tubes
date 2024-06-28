package module

import (
	"context"
	// "errors"
	"fmt"
	"github.com/ghaidafasya24/be-tubes/model"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// INSERT MENU
func InsertMenu(db *mongo.Database, col string, menu model.Menu) (insertedID primitive.ObjectID, err error) {
	// Membuat dokumen BSON untuk disimpan di MongoDB
	menurestoran := bson.M{
		"nama":       menu.Nama,
		"harga":      menu.Harga,
		"deskripsi":  menu.Deskripsi,
		"kategori":   menu.Kategori,
		"bahan_baku": menu.BahanBaku,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), menurestoran)
	if err != nil {
		fmt.Printf("InsertMenu: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}