package Structs

import "github.com/Shinizle/family-backend/Models"

type UpdateCustomerRequestStruct struct {
	Customer   Models.Customer
	FamilyList []Models.FamilyList
}
