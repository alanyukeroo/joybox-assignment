package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/list", GetListBookBySubject).Methods("GET")
	router.HandleFunc("/submit", SubmitAppointment).Methods("POST")
	return router

}

func TestGetListBookBySubject(t *testing.T) {
	request, _ := http.NewRequest("GET", "/list?subject=education", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

func TestSubmitAppointment(t *testing.T) {
	request, _ := http.NewRequest("POST", "/submit?pickup_date=2021-11-13&editionCount=772&subject=education&lender_name=alan", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}
