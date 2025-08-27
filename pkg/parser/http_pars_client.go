package parser

//"https://api.hh.ru/vacancies"
// params.Add("text", lang)
// params.Add("page", page)
// params.Add("per_page", per_page)

type ParseClient struct {
	pool []Parser
}

func NewParseClient(parsers ...Parser) *ParseClient {
	return &ParseClient{
		pool: parsers,
	}
}

func (c *ParseClient) GetPool() []Parser {
	return c.pool
}
