package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Tax-Sandbox Calculation Engine & Lacerte/ProConnect Simulator running on http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Tax-Sandbox Engine Active. Welcome Sireesh! Ready for HNI 1040 & 1065 Partnership workflows.")
	})
	http.ListenAndServe(":8080", nil)
}
