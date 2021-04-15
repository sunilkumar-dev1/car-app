package main

import (
	"CarApp/driver"
	"CarApp/handler/brandHandler"
	"CarApp/handler/modelHandler"
	"CarApp/handler/variantHandler"
	"CarApp/service/brandService"
	"CarApp/service/modelService"
	"CarApp/service/variantService"
	"CarApp/store/brand"
	"CarApp/store/model"
	"CarApp/store/variant"
	"log"
	"net/http"
)

func main() {

	DB := driver.DbConn()
	brandstore := brand.New(DB)               // at store layer return interface of store
	brandServ := brandService.New(brandstore) // at service layer return interface of services
	Bh := brandHandler.Bhandler{brandServ}    // http layer

	modelstore := model.New(DB)
	modelServ := modelService.New(modelstore)
	Mh := modelHandler.Mhandler{Mh: modelServ}

	variantstore := variant.New(DB)
	variantServ := variantService.New(variantstore)
	Vh := variantHandler.Vhandler{Vh: variantServ}

	http.HandleFunc("/brand", Bh.Route)
	http.HandleFunc("/model", Mh.Route)
	http.HandleFunc("/variant", Vh.Route)
	log.Println("Server started on: http://localhost:5000")
	http.ListenAndServe("0.0.0.0:5000", nil)
}
