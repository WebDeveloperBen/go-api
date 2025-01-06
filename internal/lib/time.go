package lib

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// ToTimestamptz converts a time.Time value to a pgtype.Timestamptz
func ToTimestamptz(t time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{
		Time:             t,
		Valid:            true,
		InfinityModifier: 0, // 0 indicates a normal timestamp
	}
}
