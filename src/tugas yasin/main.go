package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"salt-academy_learn_week2/model/repository"
	"salt-academy_learn_week2/pkg/database/mysql"
	"salt-academy_learn_week2/presenter/mahasiswa"
	"salt-academy_learn_week2/presenter/response"
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

	r := mux.NewRouter()
	r.HandleFunc("/list-mahasiswa", GetListMahasiwa)
	http.ListenAndServe(":8080", r)
}

func GetListMahasiwa(w http.ResponseWriter, router *http.Request) {
	listMahasiswa, err := repositoryMysql.GetListMahasiswa(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN GET LIST MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	responseSuccess, _ := json.Marshal(response.ResponseSuccess{
		Status: response.StatusResponse{
			Code:    0,
			Message: "SUCCESS",
		},
		Data: mahasiswa.MapperCollectionModelToCollectionJSON(listMahasiswa),
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseSuccess)
	return
}
