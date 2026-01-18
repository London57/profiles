create table preferences(
	profile_id uuid references profiles.id on delete cascade
	age smallint
	longitude real
	latitude real
)

