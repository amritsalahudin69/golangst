package main

import (
	"context"
	"dbmysql/app/db/config"
	"dbmysql/app/db/repositories"
	"encoding/json"
	"fmt"
	"log"
)

var (
	ctx             = context.Background()
	mysqlConnection = config.MysqlDBInit()
	repoSiswa       = repositories.NewSiswaRepoMysql(mysqlConnection)
)

func main() {
	//fmt.Println(mysqlConnection)
	err := mysqlConnection.Ping()

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("----------getlist------------")
	//listsiswa

	listSiswa, err := repoSiswa.GetSiswalist(ctx)
	for _, Siswa := range listSiswa {
		fmt.Println(Siswa.Name)
	}
	jsonData, err := json.Marshal(listSiswa)
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	//listsiswa

	fmt.Println("----------getone------------")
	//getsiswaone

	idsiswa := 2
	siswa, err := repoSiswa.GetSiswa(idsiswa, ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(siswa)

	//getsiswaone

	fmt.Println("----------adddatasiswa------------")
	//adddatasiswa

	dummydata := models.Siswa{
		Name:      "Mawar",
		DateBirth: time.Now().AddDate(1991, 2, 6),
		Gender:    "Perempuan",
	}
	fmt.Println(dummydata)

	err = repoSiswa.AddSiswa(dummydata, ctx)

	if err != nil {
		fmt.Println(err)
	}
	jsonData, err := json.Marshal("data Sukses di input")
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	//adddatasiswa

	fmt.Println("----------delete------------")

	IdDelete := 4
	repoSiswa.DeleteSiwa(IdDelete, ctx)

	fmt.Println("----------update------------")
	//updatedatasiswa
	dummyupdate := "Bambang"
	fmt.Println(dummyupdate)

	idupdate := 5
	err = repoSiswa.UpdateSiswa(idupdate, dummyupdate, ctx)

	if err != nil {
		fmt.Println(err)
	}
	jsonData, err := json.Marshal("data Sukses di Ubah")
	if err != nil {
		log.Fatal(err)
	}
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	//updatedatasiswa
}
