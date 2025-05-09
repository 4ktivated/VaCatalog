package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	curl "net/url"
)

const url = "https://api.hh.ru/vacancies"

// TODO: Протестировать параметры, потому что по сути вносить то надо только языки как я понял
var param = map[string]string{"text": "php", "page": "1", "per_page": "100"}

type HHparser struct {
	langs []string
}

func NewHHparser(langs []string) *HHparser {
	return &HHparser{langs: langs}
}

type Vacs []Vac

type HHoutData struct {
	Items Vacs `json:"items"`
}

type Vac struct {
	ID      string   `json:"id"`
	Lang    string   `json:"lang"`
	Title   string   `json:"name"`
	Company Employer `json:"employer"`
	URL     string   `json:"alternate_url"`
	Salary  Salary   `json:"salary"`
	Info    Snippet  `json:"snippet"`
}

type Salary struct {
	To      float64 `json:"to"`
	Curency string  `json:"currency"`
}

type Employer struct {
	Name string `json:"name"`
	Url  string `json:"alternate_url"`
}

type Snippet struct {
	Info string `json:"responsibility"`
}

func (h *HHparser) Pars() (*Vacs, error) {
	var bodyResponse *HHoutData

	resp, err := h.recvData()
	if err != nil {
		return nil, fmt.Errorf("cant get recv data error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &bodyResponse)
	if err != nil {
		return nil, err
	}

	return &bodyResponse.Items, nil
}

func (h *HHparser) recvData() (*http.Response, error) {

	params := curl.Values{}
	for prm, val := range param {
		params.Add(prm, val)
	}

	path, err := curl.Parse(url)
	if err != nil {
		return nil, err
	}
	path.RawQuery = params.Encode()

	response, err := http.Get(path.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response, nil

}
