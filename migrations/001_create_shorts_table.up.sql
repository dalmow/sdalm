create table shorts (
	short_id char(6) primary key,
	alias char(12),
	original_url varchar(255) not null,
	expires_at timestamp,
	created_at timestamp not null default current_timestamp
);

create unique index idx_shorts_alias_unique on shorts (alias) where alias is not null;