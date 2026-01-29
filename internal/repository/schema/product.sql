create table public.products (
  id bigint not null,
  name varchar not null,
  price numeric not null,
  stock int2 not null,
  category_id varchar not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  constraint products_pkey primary key (id),
  constraint products_category_id_fkey foreign KEY (category_id) references categories (id) on update CASCADE on delete CASCADE
);