package save

import (
	"log/slog"
	"net/http"
)

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

type Request struct {
	URL   string `json:"url" validate:"required, url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Alias  string `json:"alias,omitempty"`
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandleFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
