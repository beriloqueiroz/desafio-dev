package implements

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/beriloqueiroz/desafio-back/core/pkg"
	"github.com/beriloqueiroz/desafio-back/web_worker/internal/entity"
	"github.com/sony/gobreaker/v2"
	"log/slog"
	"net/http"
)

type WebRestService struct {
	Url string
}

func NewWebRestService(url string) *WebRestService {
	initCircuitBreak()
	return &WebRestService{Url: url}
}

func (ws *WebRestService) Send(ctx context.Context, notifications []entity.Notification) error {
	bodyBytes, err := json.Marshal(notifications)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	reader := bytes.NewReader(bodyBytes)
	return sendWithCircuitBreaker(ws.Url, reader)
}

func send(url string, reader *bytes.Reader) error {
	resp, err := http.DefaultClient.Post(url, "application/json", reader)
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		slog.Error(resp.Status)
		return errors.New(resp.Status)
	}
	return nil
}

func initCircuitBreak() {
	cb = pkg.NewCircuitBreak[[]byte]("SEND POST WEB APP", 5, 0.5)
}

var cb *gobreaker.CircuitBreaker[[]byte]

func sendWithCircuitBreaker(url string, reader *bytes.Reader) error {
	_, err := cb.Execute(func() ([]byte, error) {
		return nil, send(url, reader)
	})
	if err != nil {
		return err
	}
	return nil
}
