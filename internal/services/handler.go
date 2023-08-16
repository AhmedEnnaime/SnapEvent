package services

import (
	"github.com/AhmedEnnaime/SnapEvent/internal/store"
	"github.com/rs/zerolog"
)

type Handler struct {
	logger *zerolog.Logger
	us     *store.UserStore
	es     *store.EventStore
}

func New(l *zerolog.Logger, us *store.UserStore, es *store.EventStore) *Handler {
	return &Handler{
		logger: l,
		us:     us,
		es:     es,
	}
}
