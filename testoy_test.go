package gojoapi

import (
	"fmt"
	"testing"
)

func TestInsertMahasiswa(t *testing.T) {
	phone_number := "6811110023231"
	nama := "aderai"
	prodi := "D4 TI"
	kelas := "2A"
	hasil := InsertMahasiswa(nama, phone_number, prodi, kelas)
	fmt.Println(hasil)

}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
	phone_number := "6811110023231"
	biodata := GetMahasiswaFromPhone(phone_number)
	fmt.Println(biodata)
}
