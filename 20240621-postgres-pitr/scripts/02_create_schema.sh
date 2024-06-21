#!/usr/bin/env bash
set -ex
PGPASSWORD="zalando" psql -h 127.0.0.1 -p 10432 -U postgres -c "CREATE DATABASE pitr_demo"
PGPASSWORD="zalando" psql -h 127.0.0.1 -p 10432 -U postgres -d pitr_demo -f 01_schema.sql
