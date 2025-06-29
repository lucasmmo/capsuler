package testify

import (
	"testing"
)

func AssertNil(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("expected nil error, got %s", err.Error())
	}
}

func AssertNotEmptyAny(t *testing.T, expected []any) {
	if len(expected) == 0 {
		t.Fatalf("expected not empty any error, got 0")
	}
}

func AssertNotNil(t *testing.T, expected any) {
	if expected == nil {
		t.Fatalf("expected not nil error, got nil")
	}
}

func AssertNotEmptyStr(t *testing.T, expected string) {
	if expected == "" {
		t.Fatalf("expected not empty got %s", expected)
	}
}

func AssertEqualsInt(t *testing.T, expected, recieved int) {
	if expected != recieved {
		t.Fatalf("expected %d error, got %d", expected, recieved)
	}
}

func AssertNotEmptyStrSlice(t *testing.T, slice []string) {
	if len(slice) == 0 {
		t.Fatalf("expected empty error, got %d", len(slice))
	}

}

func AssertEqualsErr(t *testing.T, expected error, recieved error) {
	if expected != recieved {
		t.Fatalf("expected %s error, got %s", expected.Error(), recieved.Error())
	}
}

func AssertFalse(t *testing.T, recieved bool) {
	if recieved {
		t.Fatal("expected false error, got true")
	}
}

func AssertTrue(t *testing.T, recieved bool) {
	if !recieved {
		t.Fatal("expected true error, got false")
	}
}
