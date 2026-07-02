package challenges

import (
	"context"
	"errors"
)

type Status string

const (
	StatusPassed        Status = "passed"
	StatusFailed        Status = "failed"
	StatusCompileError  Status = "compile_error"
	StatusTimeout       Status = "timeout"
	StatusInternalError Status = "internal_error"
)

var ErrChallengeNotFound = errors.New("challenge not found")

type CodeRequest struct {
	Code string `json:"code"`
}

type Result struct {
	Status      Status `json:"status"`
	Stdout      string `json:"stdout"`
	Stderr      string `json:"stderr"`
	TestsPassed int    `json:"testsPassed"`
	TestsTotal  int    `json:"testsTotal"`
	DurationMs  int64  `json:"durationMs"`
	Hint        string `json:"hint"`
}

type Runner interface {
	Run(ctx context.Context, challengeID string, code string) (Result, error)
	Submit(ctx context.Context, challengeID string, code string) (Result, error)
}
