package pkg

import (
	"path/filepath"
	"strings"
)

func GetContentType(filename string) string {
	ext := filepath.Ext(filename)
	switch ext {
	case ".html":
		return "text/html"
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".woff":
		return "font/woff"
	case ".woff2":
		return "font/woff2"
	case ".ttf":
		return "font/ttf"
	case ".svg":
		return "image/svg+xml"
	case ".png":
		return "image/png"
	case ".jpg":
		return "image/jpeg"
	case ".ico":
		return "image/x-icon"
	case ".json":
		return "application/json"
	case ".xml":
		return "application/xml"
	case ".txt", ".md":
		return "text/plain"
	case ".mp4":
		return "video/mp4"
	case ".webm":
		return "video/webm"
	case ".ogg":
		return "video/ogg"
	default:
		return "text/plain"
	}
}

func SlashIndexFile(filename string) string {
	if filename == "" {
		filename = "index.html"
	} else if filename[len(filename)-1:] == "/" {
		filename += "index.html"
	}
	return filename
}

func ReplaceDoubleSlash(filename string) string {
	filename = strings.ReplaceAll(filename, "//", "/")
	return filename
}
