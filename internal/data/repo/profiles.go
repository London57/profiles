package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/interfaces/repo"
)

type ProfilesRepo struct {
	pool *pgxpool.Pool
	repo.ProfilesRepo
}

func (r ProfilesRepo) Create_profile(ctx context.Context, profile entities.ProfileEntity) (entities.ProfileEntity, error) {
	var (
		conn *pgxpool.Conn
		err error
		waiting time.Duration
	)
	stmt := fmt.Sprintf(`insert into %s values (?, ?, ?, ?, ?, ?)`, profilesDB) 

	for i := 0; i < maxRetries; i++ {
		conn, err = r.pool.Acquire(ctx)
		if err == nil {
			conn.Exec(ctx, stmt, profile.ID, profile.Age, profile.Name, profile.Gender, profile.Longitude, profile.Latitude)
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
			return entities.ProfileEntity{}, ctx.Err()
		case <- time.After(waiting):
			continue
		}
	}
}