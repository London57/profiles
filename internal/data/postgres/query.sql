-- name: CreateProfile :one
INSERT INTO social.profiles (birthday, email, name, username, password, gender, longitude, latitude, phone_number) VALUES (sqlc.arg(birthday)::date, sqlc.arg(email)::varchar(320), sqlc.arg(name)::varchar(50), sqlc.arg(username)::varchar(30), sqlc.arg(password)::varchar(250), sqlc.arg(gender)::smallint, sqlc.narg(longitude)::real, sqlc.narg(latitude)::real, sqlc.narg(phone_number)::varchar(15)) RETURNING *;

-- name: UpdateProfile :one
UPDATE social.profiles
SET
    birthday = COALESCE(sqlc.narg(birthday)::date, birthday),
    name = COALESCE(sqlc.narg(name)::varchar(50), name),
    phone_number = COALESCE(sqlc.narg(phone_number)::varchar(15), phone_number),
    username = COALESCE(sqlc.narg(username)::varchar(30), username),
    longitude = COALESCE(sqlc.narg(longitude)::real, longitude),
    latitude = COALESCE(sqlc.narg(latitude)::real, latitude),
    like_ttl = COALESCE(sqlc.narg(like_ttl)::smallint, like_ttl) RETURNING *;

-- name: GetProfileByEmail :one
SELECT id FROM social.profiles
WHERE email=$1;

-- name: UpdatePreferences :one
UPDATE social.preferences
SET
    age_from = COALESCE(sqlc.narg(age_from)::smallint, latitude),
    age_to = COALESCE(sqlc.narg(age_to)::smallint, latitude),
    radius = COALESCE(sqlc.narg(radius)::int, latitude)
WHERE profile_id=sqlc.arg(profile_id)::uuid RETURNING *;