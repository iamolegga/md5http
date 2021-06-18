package handler_test

import (
	"crypto/md5"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/iamolegga/md5http/internal/handler"
	"github.com/iamolegga/md5http/internal/handler/mock"
	"github.com/iamolegga/md5http/pkg/randbuf"
)

func TestHandler_HandlePositive(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	f := mock.NewMockFetcher(mockCtrl)
	h := mock.NewMockHasher(mockCtrl)
	w := mock.NewMockStringWriter(mockCtrl)

	url := "https://google.com"
	urlBody, _ := randbuf.New(42)
	hashBytes, _ := randbuf.New(md5.Size)
	hash := string(hashBytes)
	result := fmt.Sprintf("%s %s\n", url, hash)

	f.EXPECT().Fetch(url).Return(urlBody, nil).Times(1)
	h.EXPECT().Hash(urlBody).Return(hash).Times(1)
	w.EXPECT().WriteString(result).Return(len(result), nil).Times(1)

	if err := handler.New(f, h, w).Handle(url); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestHandler_HandleNegativeFetch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	f := mock.NewMockFetcher(mockCtrl)
	h := mock.NewMockHasher(mockCtrl)
	w := mock.NewMockStringWriter(mockCtrl)

	url := "https://google.com"
	errFetch := errors.New("fetch mock error")

	f.EXPECT().Fetch(url).Return(nil, errFetch).Times(1)
	h.EXPECT().Hash(gomock.Any()).Times(0)
	w.EXPECT().WriteString(gomock.Any()).Times(0)

	if err := handler.New(f, h, w).Handle(url); !errors.Is(err, errFetch) {
		t.Errorf("error want: %v, got: %v", errFetch, err)
	}
}

func TestHandler_HandleNegativeWrite(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	f := mock.NewMockFetcher(mockCtrl)
	h := mock.NewMockHasher(mockCtrl)
	w := mock.NewMockStringWriter(mockCtrl)

	url := "https://google.com"
	urlBody, _ := randbuf.New(42)
	hashBytes, _ := randbuf.New(md5.Size)
	hash := string(hashBytes)
	result := fmt.Sprintf("%s %s\n", url, hash)
	errWrite := errors.New("write mock error")

	f.EXPECT().Fetch(url).Return(urlBody, nil).Times(1)
	h.EXPECT().Hash(urlBody).Return(hash).Times(1)
	w.EXPECT().WriteString(result).Return(0, errWrite).Times(1)

	if err := handler.New(f, h, w).Handle(url); !errors.Is(err, errWrite) {
		t.Errorf("error want: %v, got: %v", errWrite, err)
	}
}
