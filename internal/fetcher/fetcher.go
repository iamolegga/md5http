package fetcher

import (
	"io"
	"net/http"
)

type Fetcher struct{}

func New() *Fetcher {
	return &Fetcher{}
}

func (f *Fetcher) Fetch(url string) (body []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	body, err = io.ReadAll(resp.Body)
	defer func() {
		closeBodyErr := resp.Body.Close()
		//do not overwrite main error with not so important error
		if err == nil {
			err = closeBodyErr
		}
	}()

	return
}
