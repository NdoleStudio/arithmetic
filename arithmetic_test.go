package arithmetic

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideIntReturnFloat(t *testing.T) {
	testData := []struct {
		name        string
		numerator   int
		denominator int
		result      float64
	}{
		{
			"whole number (6/3 = 2)",
			6,
			3,
			2,
		},
		{
			"Fractions (1/3 = 0.33*)",
			1,
			3,
			0.3333333333333333,
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.result, DivideIntsReturnFloat(data.numerator, data.denominator))
		})
	}
}

func BenchmarkDivideIntsReturnFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DivideIntsReturnFloat(1, 3)
	}
}

func BenchmarkDivideIntsReturnFloatNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = float64(1) / float64(3)
	}
}

func TestMinInt(t *testing.T) {
	t.Run("min value first", func(t *testing.T) {
		assert.Equal(t, -3, MinInt(-3, 2))
	})

	t.Run("min value last", func(t *testing.T) {
		assert.Equal(t, -231, MinInt(4, -231))
	})
}

func BenchmarkMinInt(b *testing.B) {
	first := 31
	second := 42
	for i := 0; i < b.N; i++ {
		_ = MinInt(first, second)
	}
}

func BenchmarkMathMin(b *testing.B) {
	first := 31
	second := 42
	for i := 0; i < b.N; i++ {
		_ = math.Min(float64(first), float64(second))
	}
}

func TestMinInts(t *testing.T) {
	t.Run("one value", func(t *testing.T) {
		assert.Equal(t, 1, MinInts(1))
	})

	t.Run("two values", func(t *testing.T) {
		assert.Equal(t, 1, MinInts(2, 1))
	})

	t.Run("five values sorted", func(t *testing.T) {
		assert.Equal(t, 1, MinInts(1, 5, 4, 3, 2, 1))
	})

	t.Run("five values unsorted", func(t *testing.T) {
		assert.Equal(t, 1, MinInts(1, 4, 1, 2, 3, 5))
	})
}

func BenchmarkMinInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MinInts(42, 90, 30)
	}
}

func TestMaxInt(t *testing.T) {
	t.Run("max value first", func(t *testing.T) {
		assert.Equal(t, 2, MaxInt(2, -1))
	})

	t.Run("max value last", func(t *testing.T) {
		assert.Equal(t, 4, MaxInt(-4, 4))
	})
}

func BenchmarkMaxInt(b *testing.B) {
	first := 31
	second := 42
	for i := 0; i < b.N; i++ {
		_ = MaxInt(first, second)
	}
}

func BenchmarkMathMax(b *testing.B) {
	first := 31
	second := 42
	for i := 0; i < b.N; i++ {
		_ = math.Max(float64(first), float64(second))
	}
}

func TestMaxInts(t *testing.T) {
	t.Run("one value", func(t *testing.T) {
		assert.Equal(t, 1, MaxInts(1))
	})

	t.Run("two values", func(t *testing.T) {
		assert.Equal(t, 2, MaxInts(2, 1))
	})

	t.Run("five values sorted", func(t *testing.T) {
		assert.Equal(t, 5, MaxInts(1, 5, 4, 3, 2, 1))
	})

	t.Run("five values unsorted", func(t *testing.T) {
		assert.Equal(t, 5, MaxInts(1, 4, 1, 2, 3, 5))
	})
}

func BenchmarkMaxInts(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxInts(42, 90, 30)
	}
}
