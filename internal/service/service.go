package service

import "co-viewing/internal/store/postgres"

type Service struct {
	store *postgres.Store
}

func NewService(store *postgres.Store) *Service {
	return &Service{store: store}
}
