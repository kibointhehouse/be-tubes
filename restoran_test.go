package _714220031

import (
	"fmt"
	"testing"

	"github.com/ghaidafasya24/be-tubes/model"
	"github.com/ghaidafasya24/be-tubes/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// INSERT MENU
func TestInsertMenu(t *testing.T) {
	nama := "Kopi"
	harga := "10.000"
	deskripsi := "minuman dengan rasa spesial"

	var kategori = model.Kategori{
		Kategori: "Minuman",
	}

	var bahanBaku = model.BahanBaku{
		BahanBaku: "Bubuk kopi, Air, Gula",
		Jumlah:          "3",
	}

	menurestoran := model.Menu{
		Nama:      nama,
		Harga:     harga,
		Deskripsi: deskripsi,
		Kategori:  kategori,
		BahanBaku: bahanBaku,
	}

	insertedID, err := module.InsertMenu(module.MongoConn, "restoran", menurestoran)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

// BY ID 
func TestGetMenuFromID(t *testing.T) {
	id := "667e27a6cccefc9e0156f40d"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	menurestoran, err := module.GetMenuFromID(objectID, module.MongoConn, "restoran")
	if err != nil {
		t.Fatalf("error calling GetMenuFromID: %v", err)
	}
	fmt.Println(menurestoran)
}

// ALL
func TestGetAll(t *testing.T) {
	data := module.GetAllMenu(module.MongoConn, "restoran")
	fmt.Println(data)
}


// DELETE
func TestDeleteMenuByID(t *testing.T) {
	id := "667e5f1c0da481424d4fae0b" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteMenuByID(objectID, module.MongoConn, "restoran")
	if err != nil {
		t.Fatalf("error calling DeletePresensiByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetMenuFromID(objectID, module.MongoConn, "restoran")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}