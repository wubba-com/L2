package main

import (
	chainRes "L2/task1/chainRes/pattern"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/welcome/",
		chainRes.Chain(
			welcome,
			chainRes.Logging(),
			chainRes.Method("GET"),
			chainRes.SetCORS(),
			chainRes.SetHeaderText(),
		),
	)

	log.Fatal(http.ListenAndServe(":3000", r))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello server</h1>"))
}
