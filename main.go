package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Pagina struct {
	Tafel1 bool
	Tafel2 bool
}

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
	http.HandleFunc("/", RadioButtons)
	// http.HandleFunc("/stuur", StuurNaarAPI)
	http.HandleFunc("/api", OntvangAPI)
	http.ListenAndServe("localhost:80", nil)
}

func OntvangAPI(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var tData TafelData
	json.Unmarshal(reqBody, &tData)
	BarTijs = append(BarTijs, tData)
	json.NewEncoder(w).Encode(tData)
	fmt.Println(ioutil.ReadAll(r.Body))
	fmt.Fprintf(w, "")
}

func StuurNaarAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		TafelDataAPI := TafelData{
			Tafel1JSON: Tafel1Value,
			Tafel2JSON: Tafel2Value,
		}
		// body, _ := json.Marshal(TafelDataAPI)
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(TafelDataAPI)
		req, _ := http.NewRequest("POST", "/api", payloadBuf)
		if req == nil {
			fmt.Println("hallo req is nil")
		}
		// req, _ := http.Post("</api>", "application/json", bytes.NewBuffer(body))
		// if req == nil {
		// 	fmt.Println("hallo req is nil")
		// }
		// defer req.Body.Close()
		// if req.StatusCode == http.StatusCreated {
		// 	body, err := ioutil.ReadAll(req.Body)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	jsonStr := string(body)
		// 	fmt.Println("Response: ", jsonStr)

		// } else {
		// 	fmt.Println("Get failed with error: ", req.Status)
		// }
		fmt.Println("Sturen", payloadBuf)

	}
}

func RadioButtons(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		Tafel1V, err := strconv.ParseBool(r.Form.Get("Tafel1I"))
		if err != nil {
			log.Fatal(err)
		}
		Tafel1Value = Tafel1V

		Tafel2V, err := strconv.ParseBool(r.Form.Get("Tafel2I"))
		if err != nil {
			log.Fatal(err)
		}
		Tafel2Value = Tafel2V
	}
	data := Pagina{
		Tafel1: Tafel1Value,
		Tafel2: Tafel2Value,
	}
	tmpl.Execute(w, data)
	StuurNaarAPI(w, r)
}
