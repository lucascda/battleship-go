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
