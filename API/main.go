package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TafelData struct {
	Tafel1JSON bool `json:"Tafel1Status"`
	Tafel2JSON bool `json:"Tafel2Status"`
}

var (
	Tafel1Value bool
	Tafel2Value bool
	BarTijs     []TafelData
)

func main() {
	http.HandleFunc("/", OntvangAPI)
	http.ListenAndServe("localhost:4000", nil)
}

func OntvangAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		tData := TafelData{}
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &tData)
		BarTijs = append(BarTijs, tData)
		json.NewEncoder(w).Encode(tData)
		fmt.Fprintf(w, "Welkom op de API pagina")
		fmt.Println(BarTijs)

	}
}
