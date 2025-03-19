package controllers

import (
	"log/slog"
	"net/http"
	"net_http_middleware/models"
	"path"
	"text/template"
)

func BlogsIndex(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handling BlogsIndex request",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
	)

	blogs := models.BlogsAll()
	fp := path.Join("templates", "blogs", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		slog.Error("Failed to parse template",
			slog.String("file", fp),
			slog.String("error", err.Error()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, blogs); err != nil {
		slog.Error("Failed to execute template",
			slog.String("error", err.Error()),
		)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	slog.Info("Successfully rendered BlogsIndex")
}
