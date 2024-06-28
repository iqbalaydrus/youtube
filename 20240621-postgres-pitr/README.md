# How To Run

## 1. Minio Setup
Minio is our on-premise S3 bucket. Used for spilo's backup and WAL archiving.
We have to first create the bucket using the minio's web console named `postgres-backup`.
To start minio simply execute:
```shell
docker compose up -d minio
```
Then create the bucket by typing `http://127.0.0.1:10001` and put `admin` and `password` as its
credentials.

## 2. Create postgres1 Service
We'll simulate table drop on this instance. You can create the service by invoking:
```shell
docker compose up postgres1
```
After starting the service, check minio web console, and wait until the wal and basebackup
process is finished.

## 3. postgres1 Service Preparation
First, we'll create the database and table schema on postgres1
```shell
cd scripts
./02_create_schema.sh
```
Then we'll continuously insert data to the table
```shell
./04_insert.sh
```
And just let the insert happen until we stop it later. Take notes of the timestamp you want to restore to.

## 4. Drop the Table
```shell
PGPASSWORD="zalando" psql -h 127.0.0.1 -p 10432 -U postgres -d pitr_demo -c 'DROP TABLE pitr_data'
```

## 5. Flush the WAL
To make sure the table drop changes are sent to minio, we need to rotate the wal.
```shell
./05_wal_flush.sh
```

## 6. Edit docker-compose.yml File
Put the timestamp you get from the 3rd step into `CLONE_TARGET_TIME` field in the yaml file.
Don't forget to also put timezone in ISO8601 format. e.g: `2024-06-28T01:02:03+07:00`

## 7. Start postgres2
Now we'll restore the database to a state before the drop table happened.
```shell
cd ..
docker compose up postgres2
```

## 8. PROFIT!