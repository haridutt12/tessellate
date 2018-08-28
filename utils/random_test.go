package utils

import (
	"testing"
	"os"
	"log"
)

func TestRandString(t *testing.T) {
	num := 10
	first := RandString(num)
	second := RandString(num)

	log.Println(first)
	log.Println(second)

	if first == second {
		t.Fatalf("The strings should have been random, got the same.")
	}
	if len(first) != 10 {
		t.Fatalf("expected 10, got length %v", len(first))
	}
	if len(second) != 10 {
		t.Fatalf("expected 10, got length %v", len(second))
	}
}

func TestMain(m *testing.M){
	os.Exit(m.Run())
}
