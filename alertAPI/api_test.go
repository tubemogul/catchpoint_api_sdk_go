package alertsAPI

import "testing"

func TestGetAlertReturnCode(t *testing.T) {
	a := new(Alert)
	if a.getAlertReturnCode(0) != 1 {
		t.Error("expected: 1")
	}
	if a.getAlertReturnCode(1) != 2 {
		t.Error("expected: 2")
	}
	if a.getAlertReturnCode(2) != 0 {
		t.Error("expected: 0")
	}
	if a.getAlertReturnCode(3) != 0 {
		t.Error("expected: 0")
	}
}
