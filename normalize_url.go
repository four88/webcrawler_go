package main

import (
	"net/url"
)

func normalizeURL(u string) (string, error) {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return "", err
	}

	result := parsedUrl.Host + parsedUrl.Path

	return result, nil
}
