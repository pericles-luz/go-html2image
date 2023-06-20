package html2image

import (
	"encoding/base64"
	"log"
	"mime"
	"os"
	"path/filepath"
)

func AssetToBase64(path string) string {
	mimeType := GetMimeType(path)
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return ""
	}
	return "data:" + mimeType + ";base64," + base64.StdEncoding.EncodeToString(bytes)
}

func GetMimeType(path string) string {
	if path == "" {
		return ""
	}
	return mime.TypeByExtension(filepath.Ext(path))
}
