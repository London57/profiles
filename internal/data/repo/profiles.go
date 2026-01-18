package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/London57/profiles/internal/data/entities"
)

type ProfilesRepo struct {
	pool *pgxpool.Pool
}



func (r ProfilesRepo) Create_profile(ctx context.Context, profile entities.ProfileEntity) (*entities.ProfileEntity, error) {
	var (
		conn *pgxpool.Conn
		err error
		waiting time.Duration
	)
	
	stmt := fmt.Sprintf(`insert into %s values (?, ?, ?, ?, ?)`, profilesDB)
	res := entities.ProfileEntity{}

	for i := 0; i < maxRetries; i++ {
		conn, err = r.pool.Acquire(ctx)
		if err == nil {
			defer conn.Release()
			row := conn.QueryRow(ctx, stmt, profile.Age, profile.Name, profile.Gender, profile.Longitude, profile.Latitude)
			err := row.Scan(&res)
			if err != nil {
				return &res, fmt.Errorf("database error: %w", err)
			}
		}
		stats := r.pool.Stat()
		if stats.AcquiredConns() >= stats.MaxConns() {
			waiting = time.Duration(i*i)*100*time.Millisecond
		}
		if waiting > 1*time.Second {
			waiting = 1*time.Second
		}
		
		select {
		case <- ctx.Done():
			conn.Release()
			return entities.ProfileEntity{}, ctx.Err()
		case <- time.After(waiting):
			continue
		}
	}
	return 
}