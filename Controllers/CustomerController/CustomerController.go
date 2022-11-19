package CustomerController

import (
	"encoding/json"
	"github.com/Shinizle/family-backend/Helpers"
	"github.com/Shinizle/family-backend/Models"
	"github.com/Shinizle/family-backend/Models/Structs"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var ResponseJson = Helpers.ResponseJson
var ResponseError = Helpers.ResponseError

func Index(w http.ResponseWriter, r *http.Request) {
	var customers []Models.Customer

	if err := Models.DB.Find(&customers).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseJson(w, http.StatusOK, customers)
}

func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var customers Models.Customer
	if err := Models.DB.Preload("FamilyList").Find(&customers).First(&customers, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ResponseError(w, http.StatusNotFound, "Data tidak ditemukan.")
			return
		default:
			ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	ResponseJson(w, http.StatusOK, customers)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var customers Models.Customer

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customers); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	tx := Models.DB.Begin()
	if err := tx.Create(&customers).Error; err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	ResponseJson(w, http.StatusCreated, customers)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		ResponseError(w, http.StatusBadRequest, err.Error())
	}

	var model Structs.UpdateCustomerRequestStruct

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		ResponseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer r.Body.Close()

	tx := Models.DB.Begin()
	if tx.Where("id = ?", id).Updates(&model.Customer).RowsAffected >= 1 {
		var newFamilies = model.FamilyList
		tx.Delete(Models.FamilyList{}, "customer_id = ?", id)
		tx.Create(&newFamilies)

		model.Customer.Id = id
		tx.Commit()

		Models.DB.Preload("FamilyList").Find(&model.Customer)
		ResponseJson(w, http.StatusOK, model.Customer)
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

	var customers Models.Customer
	tx := Models.DB.Begin()
	tx.Delete(Models.FamilyList{}, "customer_id = ?", id)
	if tx.Delete(&customers, id).RowsAffected == 0 {
		tx.Rollback()
		ResponseError(w, http.StatusBadRequest, "Tidak dapat menghapus data")
		return
	}

	tx.Commit()

	response := map[string]string{"message": "Data berhasil dihapus"}
	ResponseJson(w, http.StatusOK, response)
}
