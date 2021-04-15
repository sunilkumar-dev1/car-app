package brandHandler

import (
	"CarApp/entities"
	"CarApp/service"
	"fmt"
	"net/http"
	"strconv"
)

type Bhandler struct {
	Bh service.BrandService
}

func (bh Bhandler) Route(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		bh.GetById(w, req)
	case "POST":
		bh.Create(w, req)
	case "PUT":
		bh.Update(w, req)
	case "DELETE":
		bh.Delete(w, req)
	default:
	}
}
func (h Bhandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res, err := h.Bh.GetById(id)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
func (h Bhandler) Create(w http.ResponseWriter, r *http.Request) {
	var Brand entities.Brand
	Brand.Id, _ = strconv.Atoi(r.FormValue("id"))
	Brand.Name = r.FormValue("bname")

	res, err := h.Bh.Create(Brand)

	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Bhandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, _ := strconv.Atoi(id)
	res, err := h.Bh.Delete(i)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Bhandler) Update(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("bname")
	id, _ := strconv.Atoi(r.FormValue("id")) //type casting

	res, err := h.Bh.Update(id, name)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
