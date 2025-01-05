package testutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIgnoreFields(t *testing.T) {
	type user struct {
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	type pureuser struct {
		Name string
	}

	now := time.Now()

	tests := []struct {
		name     string
		input    any
		fields   []string
		expected any
	}{
		{
			name: "값 타입 구조체 처리",
			input: user{
				Name:      "test",
				CreatedAt: now,
				UpdatedAt: now,
			},
			fields: []string{"CreatedAt", "UpdatedAt"},
			expected: user{
				Name: "test",
			},
		},
		{
			name: "포인터 타입 구조체 처리",
			input: &user{
				Name:      "test",
				CreatedAt: now,
				UpdatedAt: now,
			},
			fields: []string{"CreatedAt", "UpdatedAt"},
			expected: &user{
				Name: "test",
			},
		},
		{
			name: "CreatedAt, UpdatedAt 필드가 없는 User",
			input: &pureuser{
				Name: "name",
			},
			fields: []string{"CreatedAt", "UpdatedAt"},
			expected: &pureuser{
				Name: "name",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IgnoreFields(tt.input, tt.fields...)
			assert.Equal(t, tt.expected, result)
		})
	}
}
