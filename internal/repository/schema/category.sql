create table categories (
  id varchar not null,
  name varchar not null,
  description varchar,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  constraint categories_pkey primary key (id)
);