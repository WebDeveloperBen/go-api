package lib_test

import (
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/webdeveloperben/go-api/internal/lib"
)

func TestToTimestamptz(t *testing.T) {
	t.Run("Valid time conversion", func(t *testing.T) {
		// Test with a known time
		testTime := time.Date(2025, time.January, 12, 12, 0, 0, 0, time.UTC)
		result := lib.ToTimestamptz(testTime)

		// Assert fields of the result
		assert.Equal(t, testTime, result.Time, "Time field should match the input time")
		assert.True(t, result.Valid, "Valid field should be true")
		assert.Equal(t, pgtype.InfinityModifier(0), result.InfinityModifier, "InfinityModifier should be 0")
	})

	t.Run("Zero time", func(t *testing.T) {
		// Test with the zero time
		zeroTime := time.Time{}
		result := lib.ToTimestamptz(zeroTime)

		// Assert fields of the result
		assert.Equal(t, zeroTime, result.Time, "Time field should match the zero time")
		assert.True(t, result.Valid, "Valid field should be true")
		assert.Equal(t, pgtype.InfinityModifier(0), result.InfinityModifier, "InfinityModifier should be 0")
	})

	t.Run("Future time", func(t *testing.T) {
		// Test with a future time
		futureTime := time.Now().Add(24 * time.Hour)
		result := lib.ToTimestamptz(futureTime)

		// Assert fields of the result
		assert.Equal(t, futureTime, result.Time, "Time field should match the future time")
		assert.True(t, result.Valid, "Valid field should be true")
		assert.Equal(t, pgtype.InfinityModifier(0), result.InfinityModifier, "InfinityModifier should be 0")
	})
}
