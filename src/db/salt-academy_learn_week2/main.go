package main

import (
	"context"
	"fmt"
	"salt-academy_learn_week2/model/repository"
	"salt-academy_learn_week2/pkg/database/mysql"
)

var (
	ctx             = context.Background()
	mysqlConnection = mysql.InitMysqlDB()
	repositoryMysql = repository.NewMahasiswaRepositoryMysql(mysqlConnection)
)

func main() {
	err := mysqlConnection.Ping()

	if err != nil {
		panic(err.Error())
	}

	//fmt.Println("Connection Success")

	/*mahasiswa, err := repositoryMysql.GetMahasiswa(ctx)
	if err != nil {
		return
	}

	fmt.Println(mahasiswa)*/

	listMahasiswa, err := repositoryMysql.GetListMahasiswa(ctx)
	for _, mahasiswa := range listMahasiswa {
		fmt.Println(mahasiswa.Name)
		fmt.Println(mahasiswa.Address)

	}
}
