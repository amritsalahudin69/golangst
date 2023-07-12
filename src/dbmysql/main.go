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

	listSiswa, err := repoSiswa.GetSiswalist(ctx)
	//for _, Siswa := range listSiswa {
	//	fmt.Println(Siswa.Name)
	//}

	jsonData, err := json.Marshal(listSiswa)

	if err != nil {
		log.Fatal(err)
	}

	jsonString := string(jsonData)
	fmt.Println(jsonString)
}
