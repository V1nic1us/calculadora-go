package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("TestSoma", func(t *testing.T) {
		result := soma(1, 2, 3)
		expected := 6
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("TestMultiplica", func(t *testing.T) {
		result := multiplica(10, 10)
		expected := 100
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("TestSubtrai", func(t *testing.T) {
		result := subtrai(5, 10)
		expected := -15
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("TestDivide", func(t *testing.T) {
		result := divide(20)
		expected := float32(0.05)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})
}