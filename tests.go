// Package aeutils provides convenience functions for AppEngine.
//
// Example usage:
// main_test.go:
// func TestMain(m *testing.M) {
//	AppEngineSetup()
//	result := m.Run()
//	AppEngineShutdown()
//
//	os.Exit(result)
// }

package aeutils

import (
	"context"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

var Inst aetest.Instance
var Ctx context.Context

// AppEngineSetup sets up an AppEngine instance for tests
func AppEngineSetup() {
	var err error
	Inst, err = aetest.NewInstance(&aetest.Options{SuppressDevAppServerLog: true, StronglyConsistentDatastore: true})
	if err != nil {
		panic(err)
	}
	req, err := Inst.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}
	Ctx = appengine.NewContext(req)
}

// AppEngineShutdown tears down the AppEngine instance created by AppEngineSetup()
func AppEngineShutdown() {
	Inst.Close()
}
