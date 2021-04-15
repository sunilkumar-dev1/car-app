package modelHandler

import (
	"../../entities"
	"../../service/modelService"
	"../../store"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var modeldel = []entities.Models{
	{1, "Swift", 1},
	{2, "i20", 1},
}

func TestHandler_ModelWithId(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Models
		stCode   int
		err      error
	}{
		{"1", modeldel[0], http.StatusOK, nil},
		{"2", modeldel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/model?id=" + testCase.id
		r := httptest.NewRequest("GET", link, nil)
		w := httptest.NewRecorder()
		id, _ := strconv.Atoi(testCase.id)
		serv := store.NewMockModelInterface(ctrl)
		serv.EXPECT().GetById(id).Return(testCase.expected)
		model := modelService.New(serv)
		handler := Mhandler{Mh: model}
		handler.GetById(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_ModelCreate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Models
		stCode   int
		err      error
	}{
		{"1", modeldel[0], http.StatusOK, nil},
		{"2", modeldel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		t := strconv.Itoa(testCase.expected.BrandId)
		link := "localhost:8000/model?id=" + testCase.id + "mname=" + testCase.expected.Name + "brandId=" + t
		r := httptest.NewRequest("POST", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockModelInterface(ctrl)
		serv.EXPECT().Create(testCase.expected).Return(testCase.expected)
		model := modelService.New(serv)
		handler := Mhandler{Mh: model}
		handler.Create(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_BrandDelete(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Models
		stCode   int
		err      error
	}{
		{"1", modeldel[0], http.StatusOK, nil},
		{"2", modeldel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/model?id=" + testCase.id
		r := httptest.NewRequest("DELETE", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockModelInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Delete(id).Return(testCase.expected)
		model := modelService.New(serv)
		handler := Mhandler{Mh: model}
		handler.Delete(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
func TestHandler_BrandUpdate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Models
		stCode   int
		err      error
	}{
		{"1", modeldel[0], http.StatusOK, nil},
		{"2", modeldel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/model?id=" + testCase.id + "mname=" + testCase.expected.Name
		r := httptest.NewRequest("PUT", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockModelInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Update(id, testCase.expected.Name).Return(testCase.expected)
		model := modelService.New(serv)
		handler := Mhandler{Mh: model}
		handler.Update(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
