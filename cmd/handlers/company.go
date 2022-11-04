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
		// TODO:: add logger levels by clozhure or context
		h.logger.Info("can not parse data", err)
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
		h.logger.Error("error while trying to fetch data from repository", err)
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
		h.logger.Info("can not parse data", err)

		resp = responser.BadRequest
		return
	}

	err = h.companyService.Create(body)
	if err != nil {
		h.logger.Error("error while trying to create data from repository", err)
		resp = responser.InternalError
		return
	}
	go func() {
		data, err := json.Marshal(body)
		if err != nil {
			h.logger.Error("can nod marshal data", err)
			return
		}

		err = h.natsPub.Publish("clusterName", data)
		if err != nil {
			h.logger.Error("can not publish data after creation", err)
		}
	}()
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
		h.logger.Info("can not parse data", err)

		resp = responser.BadRequest
		return
	}

	err = h.companyService.Update(body)
	if err != nil {
		if err == company.ErrNotFound {
			resp = responser.NotFound
			return
		}
		h.logger.Error("error while trying to update data from repository", err)

		resp = responser.InternalError
		return
	}

	go func() {
		data, err := json.Marshal(body)
		if err != nil {
			h.logger.Error("can nod marshal data", err)
			return
		}

		err = h.natsPub.Publish("clusterName", data)
		if err != nil {
			h.logger.Error("can not publish data after update", err)
		}
	}()

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
		h.logger.Info("can not parse data", err)

		resp = responser.BadRequest
		return
	}

	err = h.companyService.Delete(body.Ids...)
	if err != nil {
		if err == company.ErrNotFound {
			resp = responser.NotFound
			return
		}
		h.logger.Error("error while trying to delete data from repository", err)

		resp = responser.InternalError
		return
	}

	go func() {
		data, err := json.Marshal(body)
		if err != nil {
			h.logger.Error("can nod marshal data", err)
			return
		}

		err = h.natsPub.Publish("clusterName", data)
		if err != nil {
			h.logger.Error("can not publish data after deletion", err)
		}
	}()

	resp = responser.Success
}
