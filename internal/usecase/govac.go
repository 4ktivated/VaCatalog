package usecase

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"some_app/internal/repository"
)

func SuckLie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "suck lie")
}

func ApiLang(w http.ResponseWriter, r *http.Request) {
	//TODO: заменить на норм базу
	data := repository.Data{}

	//При фромате запроса api/php
	lang, _ := strings.CutPrefix(r.RequestURI, "/api/")

	//При фромате запроса api?lang=php
	// lang := r.URL.Query().Get("lang")

	resp, err := data.GetDataVac(lang)
	//TODO: конвертировать данные надо где-то в иделе структура респонса

	if err != nil {
		log.Fatalf("Data send an error %s", err)
	}

	// s, _ := json.Marshal(data)
	fmt.Fprintf(w, "%s", string(resp))
}
