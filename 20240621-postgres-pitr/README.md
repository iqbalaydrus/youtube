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
