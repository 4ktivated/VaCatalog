package app


import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"some_code/apperror"

	"github.com/gocraft/web"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

type URLRequest struct {
	URL     string `json:"url"`
	TTLDays int    `json:"ttlDays"`
}

func (h *Handler) Main_Page(rw web.ResponseWriter, req *web.Request) (interface{}, error) {
	bodyData, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var reqJson URLRequest
	if err := json.Unmarshal(bodyData, &reqJson); err != nil {
		return nil, err
	}

	if _, err := url.ParseRequestURI(reqJson.URL); err != nil {
		return nil, apperror.ErrBadRequest
	}

	return h.service.Main_Page(req.Context(), reqJson.URL, reqJson.TTLDays)
}

func (h *Handler) GetbyLang(rw web.ResponseWriter, req *web.Request) (interface{}, error) {
	val := req.PathParams["lang"]
	return h.service.GetbyLang(req.Context(), val)
}

func (h *Handler) Ping(rw web.ResponseWriter, req *web.Request) (interface{}, error) {
	return nil, nil
}

type EndpointHandler func(rw web.ResponseWriter, req *web.Request) (interface{}, error)

func WrapEndpoint(h EndpointHandler) interface{} {
	fn := func(rw web.ResponseWriter, req *web.Request, h EndpointHandler) error {
		result, err := h(rw, req)
		if err != nil {
			return err
		}

		data, err := json.Marshal(result)
		if err != nil {
			return err
		}

		_, err = rw.Write(data)
		return err
	}
	return func(rw web.ResponseWriter, req *web.Request) {
		err := fn(rw, req, h)
		if err != nil {
			fmt.Println(err.Error())
			writeHttpCode(rw, err)
		}
	}
}

func writeHttpCode(rw web.ResponseWriter, err error) {
	switch err {
	case apperror.ErrNotFound:
		rw.WriteHeader(http.StatusNotFound)
	case apperror.ErrBadRequest:
		rw.WriteHeader(http.StatusBadRequest)
	default:
		rw.WriteHeader(http.StatusInternalServerError)
	}
}