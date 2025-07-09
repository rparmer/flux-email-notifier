package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rparmer/flux-email-notifier/config"
	"github.com/rparmer/flux-email-notifier/email"
	"github.com/rparmer/flux-email-notifier/event"
)

func main() {
	cfg := config.GetConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Unable to read body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to read body"))
			return
		}

		e, err := event.FromJson(b)
		if err != nil {
			fmt.Printf("Unable to unmarshal event: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to unmarshal event"))
			return
		}

		json, err := event.ToJsonIndent(e)
		if err != nil {
			fmt.Printf("Unable to marshal event: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to marshal event"))
			return
		}

		mail := email.New()
		mail.To = email.Contact(cfg.To)
		mail.From = email.Contact(cfg.From)
		mail.Subject = fmt.Sprintf("Flux Alert - Severity: %s", e.Severity)
		mail.Message = string(json)

		if err := mail.Send(); err != nil {
			fmt.Printf("Unable to send email: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unable to send email"))
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	port := cfg.Server.Port
	fmt.Printf("Server running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", 3000), r)
}
