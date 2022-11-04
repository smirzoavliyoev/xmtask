package handlers

import (
	"encoding/json"
	"net/http"

	company "github.com/smirzoavliyoev/xmtask/internal/companyservice"
	"github.com/smirzoavliyoev/xmtask/pkg/repositories/companies"
	"github.com/smirzoavliyoev/xmtask/pkg/responser"
)

func (h *Handlers) GetCompany(w http.ResponseWriter, r *http.Request) {
	var (
		body companies.CompanyFilter
		resp responser.Resp
	)

	defer func() {
		responser.Response(resp, w)
	}()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		resp = responser.BadRequest
		return
	}

	data, err := h.companyService.GetCompany(body)
	if err != nil {
		if err == company.ErrNotFound {
			resp = responser.NotFound
			return
		}
		resp = responser.InternalError
		return
	}

	resp = responser.Success
	resp.Body = data
}

func (h *Handlers) Create(w http.ResponseWriter, r *http.Request) {
	var (
		body companies.Company
		resp responser.Resp
	)

	defer func() {
		responser.Response(resp, w)
	}()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		resp = responser.BadRequest
		return
	}

	err = h.companyService.Create(body)
	if err != nil {
		resp = responser.InternalError
		return
	}

	resp = responser.Success
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	var (
		body companies.CompanyFilter
		resp responser.Resp
	)

	defer func() {
		responser.Response(resp, w)
	}()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		resp = responser.BadRequest
		return
	}

	err = h.companyService.Update(body)
	if err != nil {
		if err == company.ErrNotFound {
			resp = responser.NotFound
			return
		}
		resp = responser.InternalError
		return
	}

	resp = responser.Success
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		body struct {
			Ids []int `json:"ids"`
		}
		resp responser.Resp
	)

	defer func() {
		responser.Response(resp, w)
	}()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		resp = responser.BadRequest
		return
	}

	err = h.companyService.Delete(body.Ids...)
	if err != nil {
		if err == company.ErrNotFound {
			resp = responser.NotFound
			return
		}
		resp = responser.InternalError
		return
	}

	resp = responser.Success
}
