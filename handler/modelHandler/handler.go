package modelHandler

import (
	"CarApp/entities"
	"CarApp/service"
	"fmt"
	"net/http"
	"strconv"
)

type Mhandler struct {
	Mh service.ModelService
}

func (mh Mhandler) Route(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		m := req.URL.Query()
		fmt.Println("query is : ", m)
		ms := make(map[string]string)
		for i, _ := range m {
			ms[i] = m[i][0]
		}
		if len(ms) > 1 {
			mh.ReadAll(w, req)
		} else if ms["id"] == "" {
			mh.ReadAll(w, req)
		} else {
			mh.GetById(w, req)
		}
	case "POST":
		mh.Create(w, req)
	case "PUT":
		mh.Update(w, req)
	case "DELETE":
		mh.Delete(w, req)
	default:
	}
}
func (h Mhandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res, err := h.Mh.GetById(id)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
func (h Mhandler) Create(w http.ResponseWriter, r *http.Request) {
	var Model entities.Models
	Model.Id, _ = strconv.Atoi(r.FormValue("id"))
	Model.Name = r.FormValue("mname")
	Model.BrandId, _ = strconv.Atoi(r.FormValue("brandId"))
	res, err := h.Mh.Create(Model)

	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Mhandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, _ := strconv.Atoi(id)
	res, err := h.Mh.Delete(i)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Mhandler) Update(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("mname")
	id, _ := strconv.Atoi(r.FormValue("id")) //type casting

	res, err := h.Mh.Update(id, name)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
func (h Mhandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	ms := make(map[string]string)
	fmt.Println(ms)
	for i, _ := range m {
		ms[i] = m[i][0]
	}
	fmt.Println(ms)
	e, err := h.Mh.ReadAll(ms)
	if err != nil {
		fmt.Println(w, "%v", err)
	} else {
		fmt.Fprintf(w, "%v", e)
	}
}
