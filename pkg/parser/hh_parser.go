package parser

import (
	"encoding/json"
	"io"
	"net/http"
)

// HHstruct
// TODO: сделать структуру для парсера, подумать хранить ли данныне в парсере
// или парсер будет производить данные а сам лишь хранить данные для запроса
type HHparser struct {
	params map[string]string // может вынести парметры в файл с конфигом что бы туда давболять и они мапились
	data   int               // под вопросов
}

func NewHHparser(params map[string]string) *HHparser {
	return &HHparser{params: params}
}

type HHoutData struct {
	// Url    string
	// Params map[string]string
	Items Vacs `json:"items"`
}

type Vacs []Vac

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

// TODO: соеднить слои приложения как описаннов обсидиане
func (*HHparser) Parse(response *http.Response) (Vacs, error) {
	var bodyResponse *HHoutData

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &bodyResponse)
	if err != nil {
		return nil, err
	}

	return bodyResponse.Items, nil
}

func (*HHparser) recvData() (*http.Response, error) {
	return nil, nil
}

func (*HHparser) sendData(vacs Vacs) error {
	return nil
}

