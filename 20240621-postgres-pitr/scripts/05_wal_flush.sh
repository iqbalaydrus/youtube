#!/usr/bin/env bash
# This script will force WAL to switch to a new file, so that our WAL-E can back it up to minio
set -ex
PGPASSWORD="zalando" psql -h 127.0.0.1 -p 10432 -U postgres -d pitr_demo -c "SELECT pg_switch_wal();"
