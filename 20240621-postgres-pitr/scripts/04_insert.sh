#!/usr/bin/env bash
set -ex
while true
do
  PGPASSWORD="zalando" psql -h 127.0.0.1 -p 10432 -U postgres -d pitr_demo -f 03_data.sql
  sleep 1
done
