package helpers

import "testing"
import "net/http"

func TestGetErrorFromHTTPError(t *testing.T) {
	err := NewHTTPError(http.StatusOK, "key", "msg")
	expected := "key: msg"

	if err.Error() != expected {
		t.Errorf(
			"Should return a message with the following pattern: %s",
			expected,
		)
	}
}
