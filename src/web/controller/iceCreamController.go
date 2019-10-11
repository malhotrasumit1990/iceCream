package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iceCream/src/model"
	"github.com/iceCream/src/service"
)

//InsertIceCream : is a controller method to be added in URL handler
var InsertIceCream = func(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Request json cannot be read")
	}
	ic := model.IceCream{}
	json.Unmarshal(reqBody, &ic)

	icObject := service.IceCreamOperationsImpl{}
	err = icObject.AddIceCreamObj(ic)
	if err != nil {
		fmt.Fprintf(w, "Error adding IceCream in DB . "+err.Error())
	}
	w.Write([]byte("Added"))

}

//GetIceCreamByProductID : is a controller method to be added in URL handler
var GetIceCreamByProductID = func(w http.ResponseWriter, r *http.Request) {

	productID := mux.Vars(r)["productID"]

	icObject := service.IceCreamOperationsImpl{}
	iceCream := icObject.GetIceCreamByProductID(productID)

	json.NewEncoder(w).Encode(iceCream)

}

//GetAllIceCreams : is a controller method to be added in URL handler
var GetAllIceCreams = func(w http.ResponseWriter, r *http.Request) {
	icObject := service.IceCreamOperationsImpl{}
	iceCreams := icObject.GetAllIceCreams()

	json.NewEncoder(w).Encode(iceCreams)
}
