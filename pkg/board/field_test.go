package board

import "testing"

func TestNewBoardField(t *testing.T) {
	sut := NewBoardField()
	t.Run("should start as water field", func(t *testing.T) {
		sut := NewBoardField()
		expected := true

		if result := sut.isWater; result != expected {
			t.Errorf("For %v, expected isWater=%t got=%t", sut, expected, result)
		}
	})

	t.Run("should not start as ship field", func(t *testing.T) {
		expected := false

		if result := sut.isShip; result != expected {
			t.Errorf("For=%v, expected=%t, got=%t", sut, expected, result)
		}
	})

	t.Run("should not start as hit", func(t *testing.T) {
		expected := false

		if result := sut.isHit; result != expected {
			t.Errorf("For=%v, expected=%t, got=%t", sut, expected, result)
		}
	})

	t.Run("should have fieldChar field with blank string", func(t *testing.T) {
		expected := ""

		if result := sut.fieldChar; result != expected {
			t.Errorf("For=%v, expected=%s, got=%s", sut, expected, result)
		}
	})
}

func TestChangeFieldChar(t *testing.T) {
	t.Run("should change fieldChar to correct value", func(t *testing.T) {
		tests := []struct {
			name     string
			value    string
			expected string
		}{
			{"change value to X", "X", "X"},
			{"change value to W", "W", "W"},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				sut := NewBoardField()

				sut.ChangeFieldChar(test.value)

				if sut.fieldChar != test.expected {
					t.Errorf("For=%v, expected=%s, got=%s", sut, test.expected, sut.fieldChar)
				}
			})
		}
	})
}
