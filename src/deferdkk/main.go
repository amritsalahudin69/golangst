package main

import (
	"encoding/json"
	"fmt"
)

type detailSiswa struct {
	Alamat       string `json:"alamat"`
	JenisKelamin string `json:"jenis_kelamin"`
}

func main() {
	dSiswa := detailSiswa{
		Alamat:       "Jakarta",
		JenisKelamin: "Laki-laki",
	}

	jsonBytes, errJson := json.Marshal(dSiswa)
	if errJson != nil {
		fmt.Println(errJson)
	}

	fmt.Println(string(jsonBytes))

	jsonArray := `[{"alamat":"Jakarta","jenis_kelamin":"Laki-laki"},{"alamat":"Depok","jenis_kelamin":"Perempuan"},{"alamat":"Bogor","jenis_kelamin":"Laki-laki"}]`

	var listDetailSiswa []*detailSiswa

	errUnmarshal := json.Unmarshal([]byte(jsonArray), &listDetailSiswa)
	if errUnmarshal != nil {
		fmt.Println(errUnmarshal)
	}

	fmt.Println(listDetailSiswa)

	for _, value := range listDetailSiswa {
		fmt.Println("Kota :", value.Alamat, ", Jenis Kelamin :", value.JenisKelamin)
	}

	//errUnmarshal2 := json.Unmarshal([]byte(jsonString2), dSiswa2)
	//if errUnmarshal2 != nil {
	//	fmt.Println(errUnmarshal2)
	//}
	//
	//errUnmarshal3 := json.Unmarshal([]byte(jsonString3), dSiswa3)
	//if errUnmarshal3 != nil {
	//	fmt.Println(errUnmarshal3)
	//}
	//
	//fmt.Println(dSiswa1)
	//fmt.Println(dSiswa2)
	//fmt.Println(dSiswa3)

	//runApplication(true)
	//siswa := Siswa{Name: "Latief"}
	//
	//err := SayHello(siswa)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//result := random()
	//
	//switch typeResult := result.(type) {
	//case string:
	//	fmt.Println("Tipe Datanya String", typeResult)
	//case int:
	//	fmt.Println("Tipe Datanya Integer", typeResult)
	//}
}

//
//func random() interface{} {
//	return "Saya Aziz"
//}
//
//type Siswa struct {
//	Name string
//}
//
//func (s Siswa) GetName() string {
//	return s.Name
//}
//
//type HasName interface {
//	GetName() string
//}
//
//func SayHello(hasName HasName) error {
//	fmt.Println("Hello", hasName.GetName())
//	return errors.New("error")
//}
//
//func logging() {
//	message := recover()
//	fmt.Println("Terjadi Error", message)
//	fmt.Println("Selesai memanggiil function")
//}
//
//func runApplication(err bool) {
//	defer logging()
//	fmt.Println()
//	if err {
//		panic("PANIC ERROR")
//	}
//}
