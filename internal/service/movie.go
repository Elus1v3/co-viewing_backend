package service

import (
	"co-viewing/internal/models"
	"context"
)

func (s *Service) AddWatchedMovie(ctx context.Context, movie models.WatchedMovie) error {
	return s.store.AddWatchedMovie(ctx, movie)
}

func (s *Service) GetAllWatchedMovies(ctx context.Context, userId int) ([]models.WatchedMovie, error) {
	return s.store.GetWatchedMoviesById(ctx, userId)
}
