create table preferences(
	profile_id uuid references profiles.id on delete cascade
	birthday date
	longitude real
	latitude real
)

