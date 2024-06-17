package context

import "testing"

func TestInitContext(t *testing.T) {
	_, err := InitContext()
	if err != nil {
		t.Errorf("Failed to initialize context, %s.", err)
	}
}
