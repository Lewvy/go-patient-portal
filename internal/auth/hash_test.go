package auth

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mySecretPassword123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}
	if hashedPassword == password {
		t.Error("Hashed password is the same as the original password")
	}

	if !MatchPassword(password, hashedPassword) {
		t.Error("MatchPassword failed to match the original password with its hash")
	}

	emptyPassword := ""
	hashedEmptyPassword, err := HashPassword(emptyPassword)
	if err != nil {
		t.Fatalf("HashPassword for empty password returned an error: %v", err)
	}
	if !MatchPassword(emptyPassword, hashedEmptyPassword) {
		t.Error("MatchPassword failed for empty password")
	}
}

func TestMatchPassword(t *testing.T) {
	password := "anotherSecretPassword456"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error, cannot proceed with MatchPassword tests: %v", err)
	}

	if !MatchPassword(password, hashedPassword) {
		t.Error("MatchPassword returned false for a correct password and its hash")
	}

	incorrectPassword := "wrongPassword"
	if MatchPassword(incorrectPassword, hashedPassword) {
		t.Error("MatchPassword returned true for an incorrect password")
	}

	invalidHash := "notARealHash"
	if MatchPassword(password, invalidHash) {
		t.Error("MatchPassword returned true for an invalid hash")
	}

	emptyPassword := ""
	hashedEmptyPassword, err := HashPassword(emptyPassword)
	if err != nil {
		t.Fatalf("HashPassword for empty password returned an error: %v", err)
	}
	if !MatchPassword(emptyPassword, hashedEmptyPassword) {
		t.Error("MatchPassword failed for empty password and its hash")
	}

	if MatchPassword(emptyPassword, hashedPassword) {
		t.Error("MatchPassword returned true for an empty password against a non-empty hash")
	}

	if MatchPassword(password, "") {
		t.Error("MatchPassword returned true for a non-empty password against an empty hash")
	}
}
