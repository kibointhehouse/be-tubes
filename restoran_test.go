package _714220031

import (
	"fmt"
	"testing"

	"github.com/ghaidafasya24/be-tubes/model"
	"github.com/ghaidafasya24/be-tubes/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func TestInsertMenu(t *testing.T) {
	var kategori = model.Kategori{
		Nama_kategori: "Minuman",
	}
	var bahanBaku = model.BahanBaku{
		Nama_bahan_baku: "Susu",
		Jumlah:          "1",
	}

	nama := "Milk Shake"
	harga := 10000.0
	deskripsi := "Minuman dengan rasa spesial"

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