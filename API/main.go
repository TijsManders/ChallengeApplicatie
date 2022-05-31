package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TafelData struct {
	Tafel1JSON bool `json:"Tafel1Status"`
	Tafel2JSON bool `json:"Tafel2Status"`
}

var (
	tData TafelData
)

func main() {
	http.HandleFunc("/", OntvangAPI)
	http.HandleFunc("/get", StuurAPI)
	http.ListenAndServe("localhost:4000", nil)
}

func StuurAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		response, err := http.Get("https://challenge-cf3aa-default-rtdb.europe-west1.firebasedatabase.app/")
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(responseData, &tData)
		json.NewEncoder(w).Encode(tData)
		fmt.Fprintf(w, "")
		fmt.Println(tData.Tafel1JSON, tData.Tafel2JSON)
	}
}

func OntvangAPI(w http.ResponseWriter, r *http.Request) {
	// eerst naar database voor decoden
	if r.Method == http.MethodPost {
		reqBody, _ := ioutil.ReadAll(r.Body)
		json.Unmarshal(reqBody, &tData)
		json.NewEncoder(w).Encode(tData)
		fmt.Fprintf(w, "")

		TafelDataAPI := TafelData{
			Tafel1JSON: tData.Tafel1JSON,
			Tafel2JSON: tData.Tafel2JSON,
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(TafelDataAPI)
		resp, err := http.Post("https://challenge-cf3aa-default-rtdb.europe-west1.firebasedatabase.app/.json", "application/json", payloadBuf)
		if err != nil {
			log.Fatalf("An Error Occured %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)

	}
}
