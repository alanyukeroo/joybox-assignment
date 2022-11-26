package model

import "time"

type Subject struct {
	Key          string  `json:"key"`
	Name         string  `json:"name"`
	Subject_type string  `json:"subject_type"`
	Work_count   int64   `json:"work_count"`
	Works        []Works `json:"works"`
}

type Works struct {
	EditionCount int64    `json:"edition_count"`
	Title        string   `json:"title"`
	Author       []Author `json:"authors"`
}

type Author struct {
	Name string `json:"name"`
}

type LenderInfo struct {
	Name         string    `json:"lender_name"`
	Subject      string    `json:"subject"`
	PickUpDate   time.Time `json:"pickup_date"`
	BorrowedBook Works     `json:"borrowed_book"`
	Message      string    `json:"message"`
}
