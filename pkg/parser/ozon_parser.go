package parser

// OzonParser реализует интерфейс Parser
// Можно добавить поля, если потребуется (например, фильтры)
// type OzonParser struct {
// 	langs []string
// }

// func NewOzonParser() *OzonParser {
// 	return &OzonParser{}
// }

// // Pars реализует основной парсинг вакансий Ozon
// func (o *OzonParser) Pars(logger *zap.SugaredLogger) *Vacs {
// 	var wg sync.WaitGroup
// 	for _, lang := range o.langs {
// 		param["text"] = lang
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			resp, err := o.recvData()
// 			if err != nil {
// 				logger.Error("cant get recv data error", zap.Error(err))
// 				return
// 			}

// 			body, err := io.ReadAll(resp.Body)
// 			if err != nil {
// 				logger.Error("cant read body", zap.Error(err))
// 				return
// 			}

// 			err = json.Unmarshal(body, &bodyResponse)
// 			if err != nil {
// 				logger.Error("cant unmarshal body", zap.Error(err))
// 				return
// 			}

// 			vacCh <- bodyResponse
// 		}()
// 	}

// 	wg.Wait()

// 	close(vacCh)
// 	for data := range vacCh {
// 		parsedVac = append(parsedVac, data.Items...)
// 	}

// 	return &parsedVac
// }

// // recvData для Playwright не используется, возвращаем ошибку
// func (o *OzonParser) recvData() (*http.Response, error) {
// 	return nil, fmt.Errorf("OzonParser использует playwright, а не http запросы")
// }

// // Основная логика парсинга через playwright
// func (o *OzonParser) parseOzonVacanciesPlaywright(logger *zap.SugaredLogger) (*Vacs, error) {
// 	pw, err := playwright.Run()
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка запуска playwright: %w", err)
// 	}
// 	defer pw.Stop()

// 	browser, err := pw.Chromium.Launch()
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка запуска браузера: %w", err)
// 	}
// 	defer browser.Close()

// 	page, err := browser.NewPage()
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка создания страницы: %w", err)
// 	}

// 	_, err = page.Goto("https://job.ozon.ru/vacancy/?tech=Go")
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка перехода на страницу: %w", err)
// 	}

// 	// Ждём, пока вакансии появятся на странице
// 	if _, err := page.WaitForSelector("a.vacancy-card"); err != nil {
// 		return nil, fmt.Errorf("вакансии не найдены: %w", err)
// 	}

// 	elems, err := page.QuerySelectorAll("a.vacancy-card")
// 	if err != nil {
// 		return nil, fmt.Errorf("ошибка поиска карточек: %w", err)
// 	}

// 	var vacs Vacs
// 	for _, el := range elems {
// 		titleEl, _ := el.QuerySelector(".vacancy-card__title")
// 		locationEl, _ := el.QuerySelector(".vacancy-card__location")
// 		link, _ := el.GetAttribute("href")

// 		titleText, _ := titleEl.TextContent()
// 		locationText, _ := locationEl.TextContent()

// 		if link != "" && titleText != "" {
// 			vac := Vac{
// 				ID:      "", // Ozon не даёт id, можно сгенерировать если нужно
// 				Lang:    "Go",
// 				Title:   titleText,
// 				Company: Employer{Name: "Ozon", Url: "https://job.ozon.ru"},
// 				URL:     "https://job.ozon.ru" + link,
// 				Salary:  Salary{}, // Зарплата не парсится
// 				Info:    Snippet{Info: locationText},
// 			}
// 			vacs = append(vacs, vac)
// 		}
// 	}

// 	return &vacs, nil
// }
