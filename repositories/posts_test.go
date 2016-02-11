package repositories_test

import (
	"testing"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func TestMyFunction(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("Failed to create instance: %v", err)
	}
	defer inst.Close()

	req1, err := inst.NewRequest("GET", "/gophers", nil)
	if err != nil {
		t.Fatalf("Failed to create req1: %v", err)
	}
	c1 := appengine.NewContext(req1)

	req2, err := inst.NewRequest("GET", "/herons", nil)
	if err != nil {
		t.Fatalf("Failed to create req2: %v", err)
	}
	c2 := appengine.NewContext(req2)

	// Run code and tests with *http.Request req1 and req2,
	// and context.Context c1 and c2.
}
