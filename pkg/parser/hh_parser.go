package parser

import (
	"encoding/json"
	"io"
	"net/http"
	curl "net/url"
	"sync"

	"go.uber.org/zap"
)

const url = "https://api.hh.ru/vacancies"

var param = map[string]string{"text": "", "page": "1", "per_page": "100"}

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

func (h *HHparser) Pars(logger *zap.SugaredLogger) *Vacs {
	var bodyResponse *HHoutData
	vacCh := make(chan *HHoutData, 200)
	parsedVac := make(Vacs, len(h.langs))

	var wg sync.WaitGroup
	for _, lang := range h.langs {
		param["text"] = lang
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := h.recvData()
			if err != nil {
				logger.Error("cant get recv data error", zap.Error(err))
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				logger.Error("cant read body", zap.Error(err))
				return
			}

			err = json.Unmarshal(body, &bodyResponse)
			if err != nil {
				logger.Error("cant unmarshal body", zap.Error(err))
				return
			}

			vacCh <- bodyResponse
		}()
	}

	wg.Wait()

	close(vacCh)
	for  data := range vacCh {
		parsedVac = append(parsedVac, data.Items...)	
	}

	return &parsedVac
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
