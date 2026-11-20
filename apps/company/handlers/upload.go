package handlers

import (
	"bytes"
	"encoding/json/jsontext"
	"encoding/json/v2"
	"net/http"

	"noah/apps/company/data/schema"

	"gorm.io/datatypes"
)

const sgBusinessSource = "sg-business"

type companyUploadRequest struct {
	Data jsontext.Value `json:"data"`
}

type companyUploadResponse struct {
	ID     uint   `json:"id"`
	Source string `json:"source"`
}

func (deps registerDeps) UploadSGBusiness(w http.ResponseWriter, r *http.Request) {
	var req companyUploadRequest
	if err := json.UnmarshalRead(r.Body, &req); err != nil || len(req.Data) == 0 || bytes.Equal(req.Data, []byte("null")) {
		http.Error(w, "data must be a non-null JSON value", http.StatusBadRequest)
		return
	}

	upload := schema.CompanyUpload{
		Source: sgBusinessSource,
		Data:   datatypes.JSON(req.Data),
	}
	if err := deps.CompanyUploadRepo.Create(r.Context(), &upload); err != nil {
		http.Error(w, "failed to save company upload", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.MarshalWrite(w, companyUploadResponse{
		ID:     upload.ID,
		Source: upload.Source,
	})
	_, _ = w.Write([]byte("\n"))
}
