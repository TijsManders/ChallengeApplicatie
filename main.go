package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Pagina struct {
	Tafel1 bool
	Tafel2 bool
}

type TafelData struct {
	Activatie bool `json:"Activatie"`
	Alarm     bool `json:"Alarm"`
}

var (
	Tafel1Value bool
	Tafel2Value bool
)

func main() {
	http.HandleFunc("/", RadioButtons)
	http.ListenAndServe(":80", nil)
}

func StuurNaarAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("Dit werkt ook")
		TafelDataAPI := []TafelData{
			{Activatie: Tafel1Value},
			{Alarm: Tafel2Value},
		}
		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(TafelDataAPI)
		req, _ := http.NewRequest("POST", "localhost", payloadBuf)
		if req == nil {
			fmt.Println("hallo reg is nil")
		}
	}
}

func RadioButtons(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		Tafel1V, err := strconv.ParseBool(r.Form.Get("ActivatieStatus"))
		if err != nil {
			log.Fatal(err)
		}
		Tafel1Value = Tafel1V
		fmt.Println(Tafel1Value, "Tijs")

		Tafel2V, err := strconv.ParseBool(r.Form.Get("AlarmStatus"))
		if err != nil {
			log.Fatal(err)
		}
		Tafel2Value = Tafel2V
		fmt.Println(Tafel2Value, "Niet Tijs")
	}
	data := Pagina{
		Tafel1: Tafel1Value,
		Tafel2: Tafel2Value,
	}
	tmpl.Execute(w, data)
	StuurNaarAPI(w, r)
}
