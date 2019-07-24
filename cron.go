package aeutils

import (
	"errors"
	"net/http"
)

var (
	ErrNoCronHeader = errors.New("Cron request does not have the X-Appengine-Cron header")
)

// ValidateCron returns ErrNoCronHeader if the request was not sent by AppEngine Cron.
// It returns nil otherwise.
func ValidateCron(r *http.Request) error {
	if r.Header.Get("X-Appengine-Cron") != "true" {
		return ErrNoCronHeader
	}
	return nil
}
