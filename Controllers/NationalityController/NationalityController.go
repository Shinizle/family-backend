package NationalityController

import (
	"encoding/json"
	"github.com/Shinizle/family-backend/Helpers"
	"github.com/Shinizle/family-backend/Models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var ResponseJson = Helpers.ResponseJson
var ResponseError = Helpers.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var nationalities []Models.Nationality

	if err := Models.DB.Find(&nationalities).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, nationalities)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var nationalities Models.Nationality
	if err := Models.DB.Preload("FamilyList").Find(&nationalities).First(&nationalities, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data tidak ditemukan.")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, nationalities)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var nationalities Models.Nationality

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nationalities); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	tx := Models.DB.Begin()
	if err := tx.Create(&nationalities).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	ResponseJson(w, http.StatusCreated, nationalities)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var nationalities Models.Nationality

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nationalities); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	tx := Models.DB.Begin()
	if tx.Where("id = ?", id).Updates(&nationalities).RowsAffected >= 1 {
		nationalities.Id = id

		tx.Commit()
		ResponseJson(w, http.StatusOK, nationalities)
	} else {
		tx.Rollback()
		ResponseError(w, http.StatusBadRequest, "Tidak dapat mengupdate data")
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var nationalities Models.Nationality
	tx := Models.DB.Begin()
	if tx.Delete(&nationalities, id).RowsAffected == 0 {
		tx.Rollback()
		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus data")
		return
	}

	tx.Commit()

	response := map[string]string{"message": "Data berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
