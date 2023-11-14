package pkg

import "testing"

func TestGetContentType(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"test.html", "text/html"},
		{"test.css", "text/css"},
		{"test.js", "application/javascript"},
		{"test.woff", "font/woff"},
		{"test.woff2", "font/woff2"},
		{"test.ttf", "font/ttf"},
		{"test.svg", "image/svg+xml"},
		{"test.png", "image/png"},
		{"test.jpg", "image/jpeg"},
		{"test.ico", "image/x-icon"},
		{"test.json", "application/json"},
		{"test.xml", "application/xml"},
		{"test.txt", "text/plain"},
		{"test.md", "text/plain"},
		{"test.mp4", "video/mp4"},
		{"test.webm", "video/webm"},
		{"test.ogg", "video/ogg"},
		{"test.unknown", "text/plain"},
		{"", "text/plain"},
		{"test", "text/plain"},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			if got := GetContentType(test.filename); got != test.expected {
				t.Errorf("GetContentType(%s) = %v, want %v", test.filename, got, test.expected)
			}
		})
	}
}

func TestSlashIndexFile(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"folder/", "folder/index.html"},
		{"folder/file", "folder/file"},
		{"", "index.html"},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			if got := SlashIndexFile(test.filename); got != test.expected {
				t.Errorf("SlashIndexFile(%s) = %v, want %v", test.filename, got, test.expected)
			}
		})
	}
}
