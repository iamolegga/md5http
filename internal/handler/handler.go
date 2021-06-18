package handler

import (
	"fmt"
	"io"
)

//go:generate mockgen -destination=./mock/fetcher.go -package=mock . Fetcher
type Fetcher interface {
	Fetch(url string) ([]byte, error)
}

//go:generate mockgen -destination=./mock/hasher.go -package=mock . Hasher
type Hasher interface {
	Hash([]byte) string
}

//go:generate mockgen -destination=./mock/string_writer.go -package=mock io StringWriter

type Handler struct {
	fetcher Fetcher
	hasher  Hasher
	writer  io.StringWriter
}

func New(
	fetcher Fetcher,
	hasher Hasher,
	writer io.StringWriter,
) *Handler {
	return &Handler{
		fetcher: fetcher,
		hasher:  hasher,
		writer:  writer,
	}
}

func (h *Handler) Handle(url string) error {
	body, err := h.fetcher.Fetch(url)
	if err != nil {
		return fmt.Errorf("unable to fetch `%s`: %w", url, err)
	}

	hash := h.hasher.Hash(body)

	result := fmt.Sprintf("%s %s\n", url, hash)
	_, err = h.writer.WriteString(result)
	if err != nil {
		return fmt.Errorf("unable to log `%s`: %w", result, err)
	}

	return nil
}
