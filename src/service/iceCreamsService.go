package service

import (
	"github.com/iceCream/src/dataAccess"
	"github.com/iceCream/src/model"
)

//IceCreamOperations consist of Crud operations on Icecream DB
type IceCreamOperations interface {
	AddIceCreamObj(iceCream model.IceCream)
	//AddIceCreamObjs(iceCreams []model.IceCream)
	GetAllIceCreams() []model.IceCream
	//UpdateIcreamByProductID(productID string, iceCreamObj model.IceCream) model.IceCream
	GetIceCreamByProductID(productID string)
}

//IceCreamOperationsImpl implements methods if IceCreamOperations interface
type IceCreamOperationsImpl struct {
}

//AddIceCreamObj in DB. This method adds one object at a time.
func (ic *IceCreamOperationsImpl) AddIceCreamObj(iceCream model.IceCream) error {

	return dataAccess.AddIceCreamObj(iceCream)
}

//GetAllIceCreams gets all rows present in IceCream DB.
func (ic *IceCreamOperationsImpl) GetAllIceCreams() *[]model.IceCream {

	return dataAccess.GetAllIceCreams()
}

//GetIceCreamByProductID return one IceCream object with specific product Id
func (ic *IceCreamOperationsImpl) GetIceCreamByProductID(productID string) *model.IceCream {
	return dataAccess.GetIceCreamByProductID(productID)
}
