package extensions

import (
	"testing"

	"github.com/uptrace/bunrouter"
)

func TestLoadingExtensions(t *testing.T) {
	router := bunrouter.New()
	_, err := LoadExtensions(router)
	if err != nil {
		t.Fatalf("%v", err)
	}
}
