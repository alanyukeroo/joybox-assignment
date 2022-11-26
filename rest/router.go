package rest

import "net/http"

func Run() {

	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi there!"))
	}

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/submit", SubmitAppointment)
	http.HandleFunc("/list", GetListBookBySubject)

}
