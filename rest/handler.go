package rest

import (
	"encoding/json"
	"fmt"
	"github/alanyukeroo/joybox-assignment/model"
	"net/http"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"
)

// A user will need to know the Title, Author, and Edition Number of a book.
// A librarian will need to know the book information + pick up schedule for the book.

func SubmitAppointment(w http.ResponseWriter, r *http.Request) {
	var (
		lender       model.LenderInfo
		pickUpFormat = "2006-01-02"
	)

	if r.Method != "POST" {
		w.Write([]byte("Method Not Allowed!"))
		return
	}

	subject := r.URL.Query().Get("subject")
	lenderName := r.URL.Query().Get("lender_name")
	pickUpDate := r.URL.Query().Get("pickup_date")
	editionCount := r.URL.Query().Get("editionCount") //Book ID

	if subject == "" || lenderName == "" {
		w.Write([]byte("Please input subject & lender_name"))
		return
	}

	date, err := time.Parse(pickUpFormat, pickUpDate)
	if err != nil {
		w.Write([]byte("Wrong Pickup Date!"))
		fmt.Println(err)
		return
	}

	edCount, err := strconv.ParseInt(editionCount, 10, 64)
	if err != nil {
		w.Write([]byte("Please input correct editionCount!"))
		return
	}

	lender.Name = lenderName
	lender.PickUpDate = date
	lender.Subject = subject

	availableBooks := getListBookBySubject(subject)

	if len(availableBooks) > 0 {

		//Check if book exist
		for _, book := range availableBooks {
			if edCount == book.EditionCount {
				lender.BorrowedBook = book
				lender.Message = "success"
				break
			}
		}
	} else {
		w.Write([]byte("your book doesnt exist!"))
		return
	}

	if lender.Message != "success" {
		w.Write([]byte(fmt.Sprintf("Cant find editionCount %s with subject %s", editionCount, subject)))
		return
	}

	resp, err := json.Marshal(lender)
	if err != nil {
		w.Write([]byte("something wrong!"))
	}

	postToWebhook(lender)
	w.Write(resp)
}

func GetListBookBySubject(w http.ResponseWriter, r *http.Request) {
	var (
		input model.Subject
	)
	if r.Method != "GET" {
		w.Write([]byte("Method Not Allowed!"))
		return
	}

	client := resty.New()

	subject := r.URL.Query().Get("subject")

	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json", subject)

	resp, err := client.R().
		EnableTrace().
		Get(url)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(resp.Body(), &input)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	data, err := json.Marshal(input.Works)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write(data)

}

func getListBookBySubject(subject string) []model.Works {

	var (
		input model.Subject
	)

	client := resty.New()

	url := fmt.Sprintf("https://openlibrary.org/subjects/%s.json", subject)

	resp, err := client.R().
		EnableTrace().
		Get(url)

	if err != nil {
		return nil
	}

	err = json.Unmarshal(resp.Body(), &input)
	if err != nil {
		return nil
	}

	return input.Works

}

func postToWebhook(data model.LenderInfo) {
	client := resty.New()

	//Save successful history appoinment to webhook
	//Please check in https://webhook.site/#!/60ee70bc-e520-4b9d-9c6a-223226a4f836/7bf71cf7-2c0f-4581-8d20-f6cad907e0ba/1
	_, _ = client.R().
		SetBody(data).
		Post("https://webhook.site/60ee70bc-e520-4b9d-9c6a-223226a4f836")

}
