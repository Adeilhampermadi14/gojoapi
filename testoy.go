package gojoapi

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "presensiMahasiswa",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertPresensi(db *mongo.Database, checkin string, id_mhs string) (InsertedID interface{}) {
	var presensi Presensi
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = checkin
	presensi.ID_mhs = id_mhs
	return InsertOneDoc(db, "presensi", presensi)
}

func InsertMahasiswa(db *mongo.Database, nama string, phone_number string, prodi string, kelas string) (InsertedID interface{}) {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Phone_number = phone_number
	mahasiswa.Prodi = prodi
	mahasiswa.Kelas = kelas
	return InsertOneDoc(db, "mahasiswa", mahasiswa)
}

func InsertMatakuliah(db *mongo.Database, nama_matkul string, id_dosen string) (InsertedID interface{}) {
	var matakuliah Matakuliah
	matakuliah.Nama_matkul = nama_matkul
	matakuliah.ID_dosen = id_dosen
	return InsertOneDoc(db, "matakuliah", matakuliah)
}

func InsertDosen(db *mongo.Database, nama_dosen string, npm string) (InsertedID interface{}) {
	var dosen Dosen
	dosen.Nama_dosen = nama_dosen
	dosen.Npm = npm
	return InsertOneDoc(db, "dosen", dosen)
}

func InsertJam_matkul(db *mongo.Database, id_matkul string, jam_masuk string, jam_keluar string) (InsertedID interface{}) {
	var jam_matkul Jam_matkul
	jam_matkul.ID_matkul = id_matkul
	jam_matkul.Jam_masuk = jam_masuk
	jam_matkul.Jam_keluar = jam_keluar
	return InsertOneDoc(db, "Jam_matkul", jam_matkul)
}

func GetMahasiswaFromPhone(db *mongo.Database, phone_number string) (mhs Mahasiswa) {
	mahasiswa := db.Collection("mahasiswa")
	filter := bson.M{"phone_number": phone_number}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&mhs)
	if err != nil {
		fmt.Printf("GetKelasFromKodeKelas: %v\n", err)
	}
	return mhs
}
