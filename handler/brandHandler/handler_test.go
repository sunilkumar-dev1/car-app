package brandHandler

import (
	"CarApp/entities"
	"CarApp/service/brandService"
	"CarApp/store"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
)

var branddel = []entities.Brand{
	{1, "Maruti"},
	{2, "Hyundai"},
}

func TestHandler_BrandWithId(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Brand
		stCode   int
		err      error
	}{
		{"1", branddel[0], http.StatusOK, nil},
		{"2", branddel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	//h:= Bhandler{serv}
	for _, testCase := range testCases {
		link := "localhost:8000/brand?id=" + testCase.id
		r := httptest.NewRequest("GET", link, nil)
		w := httptest.NewRecorder()
		id, _ := strconv.Atoi(testCase.id)
		serv := store.NewMockBrandInterface(ctrl)
		serv.EXPECT().GetById(id).Return(testCase.expected)
		brand := brandService.New(serv)
		handler := Bhandler{Bh: brand}
		handler.GetById(w, r)
		fmt.Println(w.Code, testCase.stCode)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_BrandCreate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Brand
		stCode   int
		err      error
	}{
		{"1", branddel[0], http.StatusOK, nil},
		{"2", branddel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/brand?id=" + testCase.id + "bname=" + testCase.expected.Name
		r := httptest.NewRequest("POST", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockBrandInterface(ctrl)
		serv.EXPECT().Create(testCase.expected).Return(testCase.expected)
		brand := brandService.New(serv)
		handler := Bhandler{Bh: brand}
		handler.Create(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_BrandDelete(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Brand
		stCode   int
		err      error
	}{
		{"1", branddel[0], http.StatusOK, nil},
		{"2", branddel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/brand?id=" + testCase.id
		r := httptest.NewRequest("DELETE", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockBrandInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Delete(id).Return(testCase.expected)
		brand := brandService.New(serv)
		handler := Bhandler{Bh: brand}
		handler.Delete(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
func TestHandler_BrandUpdate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Brand
		stCode   int
		err      error
	}{
		{"1", branddel[0], http.StatusOK, nil},
		{"2", branddel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/brand?id=" + testCase.id + "bname=" + testCase.expected.Name
		r := httptest.NewRequest("PUT", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockBrandInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Update(id, testCase.expected.Name).Return(testCase.expected)
		brand := brandService.New(serv)
		handler := Bhandler{Bh: brand}
		handler.Update(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
