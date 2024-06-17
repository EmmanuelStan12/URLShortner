package context

import "testing"

func TestInitContext(t *testing.T) {
	_, err := InitRootContext()
	if err != nil {
		t.Errorf("Failed to initialize context, %s.", err)
	}
}
