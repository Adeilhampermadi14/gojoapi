package gojoapi

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

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

func InsertPresensi(checkin string, id_mhs string) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.ID_mhs = id_mhs
	return InsertOneDoc("dbmhs", "presensi", presensi)
}

func InsertMahasiswa(nama string, phone_number string, prodi string, kelas string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Phone_number = phone_number
	mahasiswa.Prodi = prodi
	mahasiswa.Kelas = kelas
	return InsertOneDoc("dbmhs", "mahasiswa", mahasiswa)
}

func InsertMatakuliah(nama_matkul string, id_dosen string) (InsertedID interface{}) {
	var matakuliah Matakuliah
	matakuliah.Nama_matkul = nama_matkul
	matakuliah.ID_dosen = id_dosen
	return InsertOneDoc("dbmhs", "matakuliah", matakuliah)
}

func InsertDosen(nama_dosen string, npm string) (InsertedID interface{}) {
	var dosen Dosen
	dosen.Nama_dosen = nama_dosen
	dosen.Npm = npm
	return InsertOneDoc("dbmhs", "dosen", dosen)
}

func InsertJam_matkul(id_matkul string, jam_masuk string, jam_keluar string) (InsertedID interface{}) {
	var jam_matkul Jam_matkul
	jam_matkul.ID_matkul = id_matkul
	jam_matkul.Jam_masuk = jam_masuk
	jam_matkul.Jam_keluar = jam_keluar
	return InsertOneDoc("dbmhs", "Jam_matkul", jam_matkul)
}

func GetMahasiswaFromPhone(phone_number string) (staf Mahasiswa) {
	mahasiswa := MongoConnect("dbmhs").Collection("mahasiswa")
	// filter := bson.M{"phone_number": phone_number}
	filter := bson.M{"phone_number": phone_number}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	// err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}
