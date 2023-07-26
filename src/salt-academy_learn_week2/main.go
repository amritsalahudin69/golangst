package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"salt-academy_learn_week2/domain/repository"
	"salt-academy_learn_week2/model"
	"salt-academy_learn_week2/model/mapper"
	repo "salt-academy_learn_week2/model/repository"
	"salt-academy_learn_week2/pkg/database/mysql"
	"salt-academy_learn_week2/presenter/response"
	"salt-academy_learn_week2/request"
	"strconv"
	"time"
)

var (
	ctx                      = context.Background()
	mysqlConnection          *sql.DB
	repositoryMysqlMahasiswa repository.MahasiswaTemplate
	repositoryMysqlCampus    repository.CampusTemplate
)

func main() {
	_ = godotenv.Load()

	mysqlConnection = mysql.InitMysqlDB()
	repositoryMysqlMahasiswa = repo.NewMahasiswaRepositoryMysql(mysqlConnection)
	repositoryMysqlCampus = repo.NewCampusRepositoryMysql(mysqlConnection)

	fmt.Println(os.Getenv("MYSQL_DB_NAME"))

	err := mysqlConnection.Ping()
	if err != nil {
		panic(err.Error())
		//fmt.Println(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/listmahasiswa", GetListMahasiswa).Methods(http.MethodGet)
	r.HandleFunc("/getmahasiswa", GetMahasiswa).Methods(http.MethodGet)
	r.HandleFunc("/addmahasiswa", CreateMahasiswa).Methods(http.MethodPost)
	r.HandleFunc("/ubahmahasiswa/{id}", UpdateMahasiswa).Methods(http.MethodPut)
	r.HandleFunc("/hapusmahasiswa/{id}", DeleteMahasiswa).Methods(http.MethodDelete)

	r.HandleFunc("/listCampus", GetListCampus).Methods(http.MethodGet)
	r.HandleFunc("/getCampus", GetCampus).Methods(http.MethodGet)
	r.HandleFunc("/addCampus", CreateCampus).Methods(http.MethodPost)
	r.HandleFunc("/ubahCampus/{id}", UpdateCampus).Methods(http.MethodPut)
	r.HandleFunc("/hapusCampus/{id}", DeleteCampus).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", r)
}

func DeleteCampus(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	param, _ := strconv.Atoi(key["id"])

	err := repositoryMysqlCampus.Delete(ctx, param)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN DELETE CAMPUS",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func UpdateCampus(w http.ResponseWriter, r *http.Request) {
	var req request.BodyCampus
	key := mux.Vars(r)
	param, _ := strconv.Atoi(key["id"])
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN UPDATE CAMPUS",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}
	err = repositoryMysqlCampus.Update(ctx, param, model.Campus{
		Name:    req.Nama,
		Email:   req.Email,
		Address: req.Address,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN UPDATE CAMPUS",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateCampus(w http.ResponseWriter, r *http.Request) {
	var req request.BodyCampus
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	err = repositoryMysqlCampus.Create(ctx, model.Campus{
		Name:    req.Nama,
		Email:   req.Email,
		Address: req.Address,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetCampus(w http.ResponseWriter, r *http.Request) {
	getCampus, err := repositoryMysqlCampus.Get(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN GET LIST CAMPUS",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapper.MapperModelCampusToJson(getCampus))
	return
}

func GetListCampus(w http.ResponseWriter, r *http.Request) {
	listCampus, err := repositoryMysqlCampus.GetList(ctx)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN GET LIST CAMPUS",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapper.MapperCollectionModelCampusToCollectionJSON(listCampus))
	return
}

func GetListMahasiswa(w http.ResponseWriter, r *http.Request) {
	listMahasiswa, err := repositoryMysqlMahasiswa.GetList(ctx)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapper.MapperCollectionModelMahasiswaToCollectionJSON(listMahasiswa))
	return
}

func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
	getMahasiswa, err := repositoryMysqlMahasiswa.Get(ctx)

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(mapper.MapperModelMahasiswaToJson(getMahasiswa))
	return
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var req request.BodyMahasiswa
	key := mux.Vars(r)
	param := key["id"]
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN UPDATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}
	DBirth, err := time.Parse("2006-01-02", req.BirthDay)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "DATE NOT MATCH",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	err = repositoryMysqlMahasiswa.Update(ctx, param, model.Mahasiswa{
		Name:      req.Nama,
		Gender:    req.Gender,
		BirthDate: DBirth,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN UPDATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var req request.BodyMahasiswa
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	DBirth, err := time.Parse("2006-01-02", req.BirthDay)
	if err != nil {
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "DATE NOT MATCH",
		}})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(responseStatusError)
		return
	}

	err = repositoryMysqlMahasiswa.Create(ctx, model.Mahasiswa{
		Name:      req.Nama,
		Gender:    req.Gender,
		BirthDate: DBirth,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN CREATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)
	param := key["id"]

	err := repositoryMysqlMahasiswa.Delete(ctx, param)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		responseStatusError, _ := json.Marshal(response.ResponseError{Status: response.StatusResponse{
			Code:    1,
			Message: "ERROR WHEN UPDATE MAHASISWA",
		}})
		w.Header().Set("Content-Type", "application/json")

		w.Write(responseStatusError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
