package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		response, err := http.Get("http://app2.default.svc.cluster.local")
		if err != nil {
			handleError(err, &w)
			return
		}

		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			handleError(err, &w)
			return
		}

		responseString := fmt.Sprintf("Hello %s", string(body))
		io.WriteString(w, responseString)
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}

func handleError(err error, w *http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(*w, "Internal server error", http.StatusInternalServerError)
}
