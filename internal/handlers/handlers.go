package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func Def(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../index.html")
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 20)
	file, _, err := r.FormFile("myFile")
	if err != nil {
		log.Printf("error getting file: %v", err)
		http.Error(w, "no file uploaded", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error reading file: %v", err)
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	input := string(data)

	result, err := service.DetectAndConvert(input)
	if err != nil {
		log.Printf("conversion error: %v", err)
		http.Error(w, "conversion failed", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now()
	outputFilename := timestamp.Format("2006-01-02_15-04-05") + ".txt"
	err = os.WriteFile(outputFilename, []byte(result), 0755)
	if err != nil {
		log.Printf("error writing file: %v", err)
		http.Error(w, "failed to save file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(result))
}
