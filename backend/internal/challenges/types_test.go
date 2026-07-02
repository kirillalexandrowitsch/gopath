package challenges

import (
	"encoding/json"
	"testing"
)

func TestStatusContractValues(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status Status
		want   string
	}{
		{name: "passed", status: StatusPassed, want: "passed"},
		{name: "failed", status: StatusFailed, want: "failed"},
		{name: "compile error", status: StatusCompileError, want: "compile_error"},
		{name: "timeout", status: StatusTimeout, want: "timeout"},
		{name: "internal error", status: StatusInternalError, want: "internal_error"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := string(test.status); got != test.want {
				t.Fatalf("expected status %q, got %q", test.want, got)
			}
		})
	}
}

func TestResultJSONContract(t *testing.T) {
	t.Parallel()

	result := Result{
		Status:      StatusPassed,
		Stdout:      "ok",
		Stderr:      "",
		TestsPassed: 3,
		TestsTotal:  3,
		DurationMs:  148,
		Hint:        "",
	}

	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("marshal result: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal result map: %v", err)
	}

	expectedFields := []string{
		"status",
		"stdout",
		"stderr",
		"testsPassed",
		"testsTotal",
		"durationMs",
		"hint",
	}
	if len(payload) != len(expectedFields) {
		t.Fatalf("expected %d fields, got %d: %s", len(expectedFields), len(payload), string(data))
	}

	for _, field := range expectedFields {
		if _, ok := payload[field]; !ok {
			t.Fatalf("expected field %q in payload: %s", field, string(data))
		}
	}
	if payload["status"] != string(StatusPassed) {
		t.Fatalf("expected status %q, got %v", StatusPassed, payload["status"])
	}

	var decoded Result
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal result: %v", err)
	}
	if decoded.Status != StatusPassed {
		t.Fatalf("expected decoded status %q, got %q", StatusPassed, decoded.Status)
	}
}
