create table pages
(
	id serial constraint pages_pk primary key,
	path varchar not null,
	template_type varchar not null,
	content json not null
);

create unique index pages_path_uindex on pages (path);
