module github/alanyukeroo/joybox-assignment

replace github.com/alanyukeroo/joybox-assignment/rest => ./rest

go 1.17

require github.com/alanyukeroo/joybox-assignment/rest v0.0.0-00010101000000-000000000000

require (
	github.com/go-resty/resty/v2 v2.7.0 // indirect
	golang.org/x/net v0.0.0-20211029224645-99673261e6eb // indirect
)
