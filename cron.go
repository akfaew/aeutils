package aeutils

import (
	"errors"
	"net/http"
)

var (
	ErrNoCronHeader     = errors.New("Cron request does not have the X-Appengine-Cron header")
	ErrNoTasknameHeader = errors.New("Task request does not have the X-Appengine-Taskname header")
)

// ValidateCron returns ErrNoCronHeader if the request was not sent by AppEngine.
// It returns nil otherwise.
func ValidateCron(r *http.Request) error {
	if r.Header.Get("X-Appengine-Cron") == "" {
		return ErrNoCronHeader
	}
	return nil
}

// ValidateTask returns ErrNoTasknameHeader if the request was not sent by AppEngine.
// It returns nil otherwise.
func ValidateTask(r *http.Request) error {
	if r.Header.Get("X-Appengine-Taskname") == "" {
		return ErrNoTasknameHeader
	}
	return nil
}
