package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const baseDownloadPath = "./downloads"
const skillFileName = "skill.md"

func (deps registerDeps) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Check if base download directory exists
	if _, err := os.Stat(baseDownloadPath); os.IsNotExist(err) {
		http.Error(w, "Download directory not found", http.StatusInternalServerError)
		fmt.Printf("[log] Error: base download directory does not exist: %s\n", baseDownloadPath)
		return
	}

	// 3. Construct file path: ./downloads/skill.md
	filePath := filepath.Join(baseDownloadPath, skillFileName)

	// 4. Check if file exists
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			fmt.Printf("[log] Error: skill.md file not found at: %s\n", filePath)
		} else {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			fmt.Printf("[log] Error: failed to open file: %v\n", err)
		}
		return
	}
	defer file.Close()

	// 5. Get file info for size
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	// 6. Set response headers for download
	w.Header().Set("Content-Type", "text/markdown; charset=utf-8")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", skillFileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 7. Copy file content to response
	if _, err := io.Copy(w, file); err != nil {
		fmt.Printf("[log] Error downloading skill.md: %v\n", err)
		return
	}

	fmt.Printf("[log] skill.md downloaded successfully\n")
