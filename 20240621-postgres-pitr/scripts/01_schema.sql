CREATE TABLE IF NOT EXISTS pitr_data (
    data_id bigserial primary key,
    ts timestamp default now(),
    val double precision
);
