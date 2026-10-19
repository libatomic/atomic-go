package atomic

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/libatomic/atomic/v2/pkg/atomic"
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

func TestUserUpdateWireRoles(t *testing.T) {
	id := atomicID(t, "CfWBScAgZmUzpjfnFb6t2")
	login := "x@example.com"

	// no roles given: the key must be absent, not {"overwrite":false}
	in := &UserUpdateInput{UserID: &id, Login: &login}
	data, err := json.Marshal(&userUpdateWire{UserUpdateInput: in})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	_ = json.Unmarshal(data, &m)
	if _, ok := m["roles"]; ok {
		t.Fatalf("roles must be omitted when empty: %s", data)
	}
	if m["login"] != login {
		t.Fatalf("login lost: %s", data)
	}

	// roles given: a plain array
	w := &userUpdateWire{UserUpdateInput: in, Roles: []string{"member", "admin"}}
	data, _ = json.Marshal(w)
	_ = json.Unmarshal(data, &m)
	arr, ok := m["roles"].([]any)
	if !ok || len(arr) != 2 || arr[1] != "admin" {
		t.Fatalf("roles must be an array: %s", data)
	}
}

func atomicID(t *testing.T, s string) atomic.ID {
	t.Helper()
	id, err := atomic.ParseID(s)
	if err != nil {
		t.Fatal(err)
	}
	return id
}
