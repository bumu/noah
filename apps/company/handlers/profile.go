package handlers

import (
	"encoding/json"
	"net/http"
	"noah/apps/company/data/schema"
	companyv1 "noah/gen/go/company/v1"

	"github.com/go-chi/render"
)

func (deps registerDeps) GetCompanyProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get company profile"))
	offset := 0
	limit := 50

	res, err := deps.CompanyProfileRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &res)
}

func (deps registerDeps) ListCompanyProfile(w http.ResponseWriter, r *http.Request) {
	offset := 0
	limit := 50

	res, err := deps.CompanyProfileRepo.List(r.Context(), limit, offset)
	if err != nil {
		w.Write([]byte("get key error\n"))
	}

	render.JSON(w, r, &res)
}

func (deps registerDeps) CreateCompanyProfile(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("create company profile"))
	if r.Method != http.MethodPost {
		w.Write([]byte("method not allowed\n"))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req companyv1.CreateCompanyProfileRequest
	var resp companyv1.CreateCompanyProfileResponse

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.Status = http.StatusBadRequest

		render.JSON(w, r, &resp)
		return
	}

	schema := schema.CompnayProfile{
		Name:    req.Name,
		UEN:     req.Uen,
		Website: req.Website,
		Phone:   req.Phone,
		Address: req.Address,
		Status:  req.Status,
	}

	err = deps.CompanyProfileRepo.Create(r.Context(), &schema)
	if err != nil {
		resp.Status = http.StatusBadRequest

		render.JSON(w, r, &resp)
		return
	}

	// w.Write([]byte("record success"))
	w.WriteHeader(http.StatusOK)

}

func (deps registerDeps) UpdateCompanyProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update company profile"))
}

func (deps registerDeps) DeleteCompanyProfile(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete company profile"))
}
