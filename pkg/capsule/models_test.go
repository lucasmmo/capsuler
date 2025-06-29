package capsule_test

import (
	"capsuler/pkg/capsule"
	"capsuler/pkg/testify"
	"testing"
	"time"
)

func TestCapsule(t *testing.T) {
	stub := capsule.Builder().
		WithName("capsule_2025").
		WithDescription("capsule to test 2025").
		WithDateToOpen(time.Now()).
		Build()

	t.Run("add a new message to a capsule", func(t *testing.T) {
		testify.AssertNil(t, stub.AddMessage("message to test"))
	})

	t.Run("open a capsule", func(t *testing.T) {
		testify.AssertNil(t, stub.Open())
	})

}
