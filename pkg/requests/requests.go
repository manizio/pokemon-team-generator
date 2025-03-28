package requests

import (
	"io"
	"net/http"
)

func MakeRequest(
	method, url string,
	data io.Reader,
) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}
