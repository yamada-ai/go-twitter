package domain

import (
	"errors"
	"testing"
)

// TestUser_Normal: 正常系テスト
func TestUser_Normal(t *testing.T) {
	// テストケースID: TC001
	user, err := NewUser(1, "validUser", "pass123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if user.id.Value() != 1 {
		t.Errorf("Expected ID to be 1, got %v", user.id)
	}

	if user.username != "validUser" {
		t.Errorf("Expected Username to be 'validUser', got %v", user.username)
	}

	if user.password != "pass123" {
		t.Errorf("Expected Password to be 'pass123', got %v", user.password)
	}
}

func TestUser_IDValidationError(t *testing.T) {
	_, err := NewUser(-1, "validUser", "password123")
	if err == nil || err.Error() != "ID value cannot be negative" {
		t.Errorf("Expected error for negative ID, got %v", err)
	}
}

// TestUser_ValidationErrors: 異常系テスト
func TestUser_ValidationErrors(t *testing.T) {
	tests := []struct {
		testCaseID string
		id         int
		username   string
		password   string
		expected   error
	}{
		{"TC002", 1, "", "pass123", errors.New("username must be at least 3 characters long")},
		{"TC003", 1, "a", "pass123", errors.New("username must be at least 3 characters long")},
		{"TC004", 1, "validUser", "", errors.New("password must be at least 6 characters long")},
		{"TC005", 1, "validUser", "123", errors.New("password must be at least 6 characters long")},
	}

	for _, tc := range tests {
		t.Run(tc.testCaseID, func(t *testing.T) {
			_, err := NewUser(tc.id, tc.username, tc.password)
			if err == nil || err.Error() != tc.expected.Error() {
				t.Errorf("[%s] Expected error: %v, got: %v", tc.testCaseID, tc.expected, err)
			}
		})
	}
}
