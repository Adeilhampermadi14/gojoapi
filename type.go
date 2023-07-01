package gojoapi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Prodi        string             `bson:"Prodi,omitempty" json:"Prodi,omitempty"`
	Kelas        string             `bson:"Kelas,omitempty" json:"Kelas,omitempty"`
}

type Presensi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Checkin  string             `bson:"checkin,omitempty" json:"checkin,omitempty"`
	ID_mhs   string             `bson:"ID_mhs,omitempty" json:"ID_mhs,omitempty"`
}

type Matakuliah struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_matkul string             `bson:"Nama_matkul,omitempty" json:"Nama_matkul,omitempty"`
	ID_dosen    string             `bson:"Dosen,omitempty" json:"Dosen,omitempty"`
}

type Dosen struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama_dosen string             `bson:"nama_dosen,omitempty" json:"nama_dosen,omitempty"`
	Npm        string             `bson:"npm,omitempty" json:"npm,omitempty"`
}
type Jam_matkul struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID_matkul  string             `bson:"ID_matkul,omitempty" json:"ID_matkul,omitempty"`
	Jam_masuk  string             `bson:"Jam_masuk,omitempty" json:"Jam_masuk,omitempty"`
	Jam_keluar string             `bson:"Jam_keluar,omitempty" json:"Jam_keluar,omitempty"`
}
