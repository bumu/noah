package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const baseDownloadPath = "./downloads"

func (deps registerDeps) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Only allow GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2. Create base download directory if it doesn't exist
	if _, err := os.Stat(baseDownloadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(baseDownloadPath, os.ModePerm); err != nil {
			http.Error(w, "Failed to create download directory", http.StatusInternalServerError)
			fmt.Printf("[log] Error: failed to create download directory: %v\n", err)
			return
		}
		fmt.Printf("[log] Created download directory: %s\n", baseDownloadPath)
	}

	// 3. Extract filename from URL path
	requestPath := r.URL.Path
	var fileName string
	if len(requestPath) > len("/downloads/") {
		fileName = requestPath[len("/downloads/"):]
	} else if len(requestPath) > len("/dl/") {
		fileName = requestPath[len("/dl/"):]
	}

	if fileName == "" {
		http.Error(w, "File name is required", http.StatusBadRequest)
		return
	}

	// 4. Prevent path traversal attacks
	if filepath.IsAbs(fileName) || len(filepath.Clean(fileName)) != len(fileName) {
		http.Error(w, "Invalid file name", http.StatusForbidden)
		return
	}

	// 5. Construct file path
	filePath := filepath.Join(baseDownloadPath, fileName)

	// 6. Verify the resolved path is still within baseDownloadPath
	absFilePath, err := filepath.Abs(filePath)
	absBasePath, _ := filepath.Abs(baseDownloadPath)
	if err != nil {
		http.Error(w, "Invalid file path", http.StatusBadRequest)
		return
	}
	if !isPathWithinDirectory(absFilePath, absBasePath) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// 7. Check if file exists
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
			fmt.Printf("[log] Error: file not found at: %s\n", filePath)
		} else {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			fmt.Printf("[log] Error: failed to open file: %v\n", err)
		}
		return
	}
	defer file.Close()

	// 8. Get file info for size
	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	// 9. Set response headers for download
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filepath.Base(fileName)))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	// 10. Copy file content to response
	if _, err := io.Copy(w, file); err != nil {
		fmt.Printf("[log] Error downloading file %s: %v\n", fileName, err)
		return
	}

	fmt.Printf("[log] File downloaded successfully: %s\n", fileName)
}

// Helper function to check if a path is within a base directory
func isPathWithinDirectory(filePath, baseDir string) bool {
	rel, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		return false
	}
	return !filepath.HasPrefix(rel, "..")
}
