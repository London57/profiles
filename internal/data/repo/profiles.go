package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/London57/profiles/internal/data/entities"
)

type BaseRepo struct {
	pool *pgxpool.Pool
}

type ProfilesRepo struct {
	pool *pgxpool.Pool
}

func (r ProfilesRepo) CreateProfile(ctx context.Context, profile entities.ProfileEntity) (*entities.ProfileEntity, error) {	
	stmt := fmt.Sprintf(`insert into %s (birthday, email, name, username, password, gender, longitude, latitude) values (?, ?, ?, ?, ?, ?, ?, ?)`, profilesDB)

	res := entities.ProfileEntity{}

	birthday := pgtype.Date{
		Time: profile.Birthday,
		Valid: true,
	}
	row := r.pool.QueryRow(ctx, stmt, birthday, profile.Email, profile.Name, profile.Username, profile.Password, profile.Gender, profile.Longitude, profile.Latitude)

	var pgBirthday pgtype.Date
	err := row.Scan(
		&res.ID,
        &pgBirthday,
        &res.Email,
        &res.Name,
        &res.Username, 
        &res.Password,
        &res.Gender,
        &res.Longitude,
        &res.Latitude,
	)

	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	if pgBirthday.Valid {
		res.Birthday = pgBirthday.Time
	}

	return &res, nil
}

func (r ProfilesRepo) UpdateProfile(ctx context.Context, fields map[string]any) (*entities.ProfileEntity, error) {
	values := make([]any, len(fields))
	keys := make([]string, len(fields))
	var result_string strings.Builder

	for k, v := range fields {
		values = append(values, v)
		result_string.WriteString(k)
		result_string.WriteByte('=')
		result_string.WriteByte('?')
		result_string.WriteString(", ")
	}

	values = append(values, fields["id"])
	s := result_string.String()
	result_string.Reset()
	result_string.WriteString(s[:len(s)-2])
	stmt := fmt.Sprintf("update %s set %s where id = ? returning %s", profilesDB, result_string.String(), strings.Join(keys, ", ")) // returning updated fields

	updated_profile := entities.ProfileEntity{}
	row := r.pool.QueryRow(ctx, stmt, values...)

	err := row.Scan(&updated_profile)
	if err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &updated_profile, nil
}

func (r ProfilesRepo) GetProfileIdByEmail(ctx context.Context, email string) (uuid.UUID, error) {
	stmt := fmt.Sprintf("select id from %s where email=?", profilesDB)

	var res uuid.UUID

	row := r.pool.QueryRow(ctx, stmt, email)
	err := row.Scan(res)
	if err != nil {
		return uuid.Nil, fmt.Errorf("database error: %w", err)
	}

	return res, nil
}

