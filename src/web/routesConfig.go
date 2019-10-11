package web

import (
	"github.com/gorilla/mux"
	"github.com/iceCream/src/constants"
	"github.com/iceCream/src/web/controller"
)

//SetupRoutes sets up all incoming request urls for application
var SetupRoutes = func() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc(constants.AddIC, controller.InsertIceCream).Methods("Post")
	//router.HandleFunc(constants.AddICs, controller.InsertIceCreams).Methods("Post")
	router.HandleFunc(constants.GetAllIc, controller.GetAllIceCreams).Methods("Get")
	router.HandleFunc(constants.GetICByProductID, controller.GetIceCreamByProductID).Methods("Get")

	return router

}
