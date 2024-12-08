package domain

import (
	"testing"
)

// 正常系テスト: 正の整数でIDを作成する
func TestID_Normal(t *testing.T) {
	// 正常なIDの生成
	id, err := NewID[User](10)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// 値が正しいか検証
	if id.Value() != 10 {
		t.Errorf("Expected value to be 10, got %d", id.Value())
	}

	// String()のフォーマットを確認
	expectedString := "ID[domain.User](10)"
	if id.String() != expectedString {
		t.Errorf("Expected string representation to be '%s', got '%s'", expectedString, id.String())
	}
}

// 異常系テスト: 負の値を使用した場合
func TestID_ValidationErrors(t *testing.T) {
	tests := []struct {
		value       int
		expectedErr string
	}{
		{-1, "ID value cannot be negative"},   // 負の値
		{-100, "ID value cannot be negative"}, // 負の値
	}

	for _, tc := range tests {
		t.Run(tc.expectedErr, func(t *testing.T) {
			_, err := NewID[User](tc.value)
			if err == nil || err.Error() != tc.expectedErr {
				t.Errorf("Expected error: '%s', got: '%v'", tc.expectedErr, err)
			}
		})
	}
}
