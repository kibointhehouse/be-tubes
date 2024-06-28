package module

import (
	"context"
	"errors"
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

func GetAllMenu(db *mongo.Database, col string) (data []model.Menu) {
	menurestoran := db.Collection(col)
	filter := bson.M{}
	cursor, err := menurestoran.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}


func GetMenuFromID(_id primitive.ObjectID, db *mongo.Database, col string) (menu model.Menu, errs error) {
	menurestoran := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := menurestoran.FindOne(context.TODO(), filter).Decode(&menu)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return menu, fmt.Errorf("no data found for ID %s", _id.Hex())
		}
		return menu, fmt.Errorf("error retrieving data for ID %s: %s", _id.Hex(), err.Error())
	}
	return menu, nil
}

func UpdateMenu(db *mongo.Database, col string, id primitive.ObjectID, nama string, harga float64, deskripsi string, kategori model.Kategori, bahanBaku model.BahanBaku) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
			"nama":       nama,
			"harga":      harga,
			"deskripsi":  deskripsi,
			"kategori":   kategori,
			"bahan_baku": bahanBaku,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateMenu: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified ID")
		return
	}
	return nil
}

func DeleteMenuByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := karyawan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}