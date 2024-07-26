package parser

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

func HHparser(text string, area string, page string, per_page string) map[string]interface{} {

	var data map[string]interface{}

	params := url.Values{}
	params.Add("text", text)
	params.Add("area", area)
	params.Add("page", page)
	params.Add("per_page", per_page)

	path, err := url.Parse("https://api.hh.ru/vacancies")
	if err != nil {
		panic(err)
	}
	path.RawQuery = params.Encode()

	response, err := http.Get(path.String())
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(body, &data)
	return data
}
