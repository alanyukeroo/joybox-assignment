package main

import (
	"fmt"
	"net/http"

	"github.com/alanyukeroo/joybox-assignment/rest"
)

func main() {

	rest.Run()

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)

}
