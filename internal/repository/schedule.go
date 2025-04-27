package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// ScheduleRepository handles database operations for schedules
type ScheduleRepository struct {
    db *pgxpool.Pool
}

// NewScheduleRepository creates a new ScheduleRepository instance
func NewScheduleRepository(db *pgxpool.Pool) *ScheduleRepository {
    return &ScheduleRepository{db: db}
}