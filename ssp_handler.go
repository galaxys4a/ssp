package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	AddEndpoint("/ssp_handler/id1")

	port := 4545
	fmt.Println(fmt.Sprintf("Server is listening on port %d...", port))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	fmt.Println(err)
}

func AddEndpoint(pattern string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		//resp, err := http.Get("http://127.0.0.1:7171/s/sb")
		//if err != nil {
		//	fmt.Fprint(w, err)
		//}
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	})
}
