package profiles

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/London57/profiles/internal/data/entities"
	"github.com/London57/profiles/internal/data/repo"
)


type ProfilesRepo struct {
	pool *pgxpool.Pool
}

func (ProfilesRepo) New(connstring string) (ProfilesRepo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	pool, err := pgxpool.New(ctx, connstring)
	if err != nil {
		return ProfilesRepo{}, err
	}
	return ProfilesRepo{
		pool: pool,
	}, nil
}

func (r ProfilesRepo) CreateProfile(ctx context.Context, profile entities.ProfileEntity) (*entities.ProfileEntity, error) {	
	stmt := fmt.Sprintf(`insert into %s (birthday, email, name, username, password, gender, longitude, latitude) values (?, ?, ?, ?, ?, ?, ?, ?) returning id, birthday, email, name, username, gender, longitude, latitude` , profilesDB)

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

func (r ProfilesRepo) UpdateProfile(ctx context.Context, id uuid.UUID, fields map[string]any) (*entities.ProfileEntity, error) {
	result_string, keys, values := repo.FieldsToExexString(fields)
	stmt := fmt.Sprintf("update %s set %s where id = ? returning %s", profilesDB, result_string, strings.Join(keys, ", ")) // returning updated fields

	updated_profile := entities.ProfileEntity{}
	values = append(values, id)
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
	err := row.Scan(&res)
	if err != nil {
		return uuid.Nil, fmt.Errorf("database error: %w", err)
	}

	return res, nil
}

func AddPreferences(ctx context.Context, fields map[string]any) error {
	stmt := fmt.Sprintf("")
}