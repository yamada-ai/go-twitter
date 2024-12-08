package domain

import (
	"errors"
	"fmt"
)

// IDは値オブジェクトで、負の値を許可しません。
type ID[T any] struct {
	value int
}

// NewIDは新しいIDを作成します。
func NewID[T any](value int) (ID[T], error) {
	if value < 0 {
		return ID[T]{}, errors.New("ID value cannot be negative")
	}
	return ID[T]{value: value}, nil
}

// ValueはIDの値を返します。
func (id ID[T]) Value() int {
	return id.value
}

// StringはIDを文字列としてフォーマットします。
func (id ID[T]) String() string {
	return fmt.Sprintf("ID[%T](%d)", *new(T), id.value)
}
