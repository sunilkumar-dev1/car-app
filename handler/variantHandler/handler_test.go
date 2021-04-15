package variantHandler

import (
	"../../entities"
	"../../service/variantService"
	"../../store"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var variantdel = []entities.Variant{
	{1, "Swift", 1, "60", "60", 1},
	{2, "i20", 1, "50", "30", 1},
}

func TestHandler_VariantWithId(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Variant
		stCode   int
		err      error
	}{
		{"1", variantdel[0], http.StatusOK, nil},
		{"2", variantdel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/variant?id=" + testCase.id
		r := httptest.NewRequest("GET", link, nil)
		w := httptest.NewRecorder()
		id, _ := strconv.Atoi(testCase.id)
		serv := store.NewMockVariantInterface(ctrl)
		serv.EXPECT().GetById(id).Return(testCase.expected)
		variant := variantService.New(serv)
		handler := Vhandler{Vh: variant}
		handler.GetById(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_VariantCreate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Variant
		stCode   int
		err      error
	}{
		{"1", variantdel[0], http.StatusOK, nil},
		{"2", variantdel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		d := strconv.Itoa(testCase.expected.Displace)
		mid := strconv.Itoa(testCase.expected.ModelsId)
		link := "localhost:8000/variant?id=" + testCase.id + "name=" + testCase.expected.Name + "displace=" + d + "peakPower=" + testCase.expected.PeakPower + "peakTorque=" + testCase.expected.PeakTorque + "modelId=" + mid
		r := httptest.NewRequest("POST", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockVariantInterface(ctrl)
		serv.EXPECT().Create(testCase.expected).Return(testCase.expected)
		variant := variantService.New(serv)
		handler := Vhandler{Vh: variant}
		handler.Create(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}

func TestHandler_VariantDelete(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Variant
		stCode   int
		err      error
	}{
		{"1", variantdel[0], http.StatusOK, nil},
		{"2", variantdel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/variant?id=" + testCase.id
		r := httptest.NewRequest("DELETE", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockVariantInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Delete(id).Return(testCase.expected)
		variant := variantService.New(serv)
		handler := Vhandler{Vh: variant}
		handler.Delete(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
func TestHandler_VariantUpdate(t *testing.T) {
	testCases := []struct {
		id       string
		expected entities.Variant
		stCode   int
		err      error
	}{
		{"1", variantdel[0], http.StatusOK, nil},
		{"2", variantdel[1], http.StatusOK, nil},
	}
	//
	ctrl := gomock.NewController(t)

	for _, testCase := range testCases {
		link := "localhost:8000/variant?id=" + testCase.id + "name=" + testCase.expected.Name
		r := httptest.NewRequest("PUT", link, nil)
		w := httptest.NewRecorder()
		serv := store.NewMockVariantInterface(ctrl)
		id, _ := strconv.Atoi(testCase.id)
		serv.EXPECT().Update(id, testCase.expected.Name).Return(testCase.expected)
		variant := variantService.New(serv)
		handler := Vhandler{Vh: variant}
		handler.Update(w, r)
		if w.Code != testCase.stCode {
			fmt.Println("error")
		}
	}
}
