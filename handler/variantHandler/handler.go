package variantHandler

import (
	"CarApp/entities"
	"CarApp/service"
	"fmt"
	"net/http"
	"strconv"
)

type Vhandler struct {
	Vh service.VariantService
}

func (vh Vhandler) Route(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		m := req.URL.Query()
		fmt.Println("query is : ", m)
		ms := make(map[string]string)
		for i, _ := range m {
			ms[i] = m[i][0]
		}
		if len(ms) > 1 {
			vh.ReadAll(w, req)
		} else if ms["id"] == "" {
			vh.ReadAll(w, req)
		} else {
			vh.GetById(w, req)
		}
	case "POST":
		vh.Create(w, req)
	case "PUT":
		vh.Update(w, req)
	case "DELETE":
		vh.Delete(w, req)
	default:
	}
}
func (h Vhandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	res, err := h.Vh.GetById(id)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
func (h Vhandler) Create(w http.ResponseWriter, r *http.Request) {
	var Variant entities.Variant
	Variant.Id, _ = strconv.Atoi(r.FormValue("id"))
	Variant.Name = r.FormValue("name")
	Variant.Displace, _ = strconv.Atoi(r.FormValue("displace"))
	Variant.PeakPower = r.FormValue("peakPower")
	Variant.PeakTorque = r.FormValue("peakTorque")
	Variant.ModelsId, _ = strconv.Atoi(r.FormValue("modelId")) //

	res, err := h.Vh.Create(Variant)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Vhandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	i, _ := strconv.Atoi(id)
	res, err := h.Vh.Delete(i)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}

func (h Vhandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id")) //type casting
	name := r.FormValue("name")
	res, err := h.Vh.Update(id, name)
	if err.Error() != "Failed" {
		fmt.Fprintf(w, "%v", res)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
func (h Vhandler) ReadAll(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query()
	ms := make(map[string]string)
	fmt.Println(ms)
	for i, _ := range m {
		ms[i] = m[i][0]
	}
	fmt.Println(ms)
	e, err := h.Vh.ReadAll(ms)
	fmt.Println(err, e)
	if err != nil {
		fmt.Println(w, "%v", err)
	} else {
		fmt.Fprintf(w, "%v", e)
	}
}
