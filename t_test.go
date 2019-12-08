package expect

import (
	"testing"
)

func TestGetTSimple(t *testing.T) {
	tt := getT()
	if t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}

func indirectGetT() *testing.T {
	return getT()
}

func TestGetTIndirect(t *testing.T) {
	tt := indirectGetT()
	if t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}

func getTWithParam(t *testing.T) *testing.T {
	return getT()
}

func TestGetTWithTestingArgument(t *testing.T) {
	tt := getTWithParam(t)
	if t != tt {
		t.Fatalf("expected: %p, got %p", t, tt)
	}
}
