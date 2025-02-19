package repository

type Data struct{}

//TODO: с базы мы берём  "буфер" иа его кже парсим там где нужно во что нужно

type UrlParam struct {
	Url   string
	Param map[string]string
}

// hardcode
var url = "https://api.hh.ru/vacancies"
var param = map[string]string{"text": "php", "page": "1", "per_page": "100"}

func (d *Data) GetDataVac(lang string) ([]byte, error) {
	//запрашиваем вакансии по языку
	return []byte(lang), nil
}

func (d *Data) GetOptUrl(brand string) (*UrlParam, error) {
	//запрашиваем из бд данные по бренд
	// up, err := json.Marshal(UrlParam{url, param})
	// if err != nil {
	// 	return nil, err
	// }
	up := &UrlParam{url, param}

	return up, nil
}
