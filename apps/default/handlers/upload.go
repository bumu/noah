package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const baseUploadPath = "./uploads"

func (deps registerDeps) UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Limit file size (e.g., 10MB)
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to parse uploaded file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 3. Generate UUID v7 (time-ordered)
	// Benefits of v7: file names are sorted by upload order on disk
	u, err := uuid.NewV7()
	if err != nil {
		http.Error(w, "Failed to generate UUID v7", http.StatusInternalServerError)
		return
	}
	fileID := u.String()

	// 4. Create save path: ./uploads/<uuid_v7>/<filename>
	saveDir := filepath.Join(baseUploadPath, fileID)
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	dstPath := filepath.Join(saveDir, handler.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 5. Copy uploaded content to target file
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to write to file", http.StatusInternalServerError)
		return
	}

	// 6. Return result
	fmt.Printf("[log] Received file: %s, assigned ID: %s\n", handler.Filename, fileID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"status":"success", "id":"%s", "file":"%s"}`+"\n", fileID, handler.Filename)
}
