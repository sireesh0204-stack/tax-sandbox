package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>ProConnect Tax Engine</title></head><body style='background:#0f172a;color:#f1f5f9;font-family:Arial,sans-serif;display:flex;min-height:100vh'><aside style='width:280px;background:rgba(30,41,59,0.95);padding:2rem;text-align:left'><div style='color:#38bdf8;font-weight:700;font-size:1.3rem;margin-bottom:2rem'>ProConnect Tax</div><div style='color:#94a3b8'>ProConnect Tax Engine Active. Ready for HNI 1040 & 1065 Partnership workflows.</div></aside><main style='flex:1;padding:3rem'><h1 style='font-size:2rem;margin-bottom:1rem'>Form 1040</h1><p style='color:#94a3b8'>David Vance - High Net Worth 2026</p></main></body></html>")
	})
	http.HandleFunc("/api/overview", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"running","client":"David Vance","total_income":1250000,"tax_year":2026}`)
	})
	fmt.Println("ProConnect Tax Engine running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}