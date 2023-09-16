package tokenBucket_test

import (
	"testing"
	"time"

	"github.com/goPlayground/tokenBucket"
)

func TestTokenBucket(t *testing.T) {
	tb := tokenBucket.NewTokenBucket(10, 10)
	if !tb.Consume(5) {
		t.Error("Expected to consume 5 tokens")
	}
	if tb.Consume(10) {
		t.Error("Expected to not consume 10 tokens")
	}
	time.Sleep(time.Second)
	if !tb.Consume(10) {
		t.Error("Expected to consume 10 tokens")
	}
}
