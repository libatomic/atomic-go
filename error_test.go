package atomic

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestErrorDetailAndNilID(t *testing.T) {
	var e Error
	body := `{"code":"request_validation","message":"request validation error","detail":{"login":"must be a valid email address"}}`
	if err := json.Unmarshal([]byte(body), &e); err != nil {
		t.Fatal(err)
	}
	msg := e.Error()
	if !strings.Contains(msg, "login: must be a valid email address") || !strings.Contains(msg, "request_validation") {
		t.Fatalf("detail lost: %q", msg)
	}
	var in UserUpdateInput
	if got := idString(in.UserID); got != "" {
		t.Fatalf("nil id rendered as %q", got)
	}
}
