package services

import (
	"fmt"
	"io"
	"testing"

	"github.com/AhmedEnnaime/SnapEvent/internal/db"
	"github.com/AhmedEnnaime/SnapEvent/internal/store"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func setUp(t *testing.T) (*Handler, func(t *testing.T)) {
	w := zerolog.ConsoleWriter{Out: io.Discard}

	l := zerolog.New(w).With().Timestamp().Logger()

	d, err := db.NewTestDbB()
	if err != nil {
		t.Fatal(fmt.Errorf("failed to initialize database: %w", err))
	}

	us := store.NewUserStore(d)
	es := store.NewEventStore(d)

	return New(&l, us, es), func(t *testing.T) {
		err := db.DropTestDB(d)
		if err != nil {
			t.Fatal(fmt.Errorf("failed to clean database: %w", err))
		}
	}
}

// func ctxWithToken(ctx context.Context, token string) context.Context {
// 	scheme := "Token"
// 	md := metadata.Pairs("authorization", fmt.Sprintf("%s %s", scheme, token))
// 	nCtx := m
// }
