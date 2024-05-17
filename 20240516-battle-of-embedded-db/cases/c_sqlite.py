import logging
import time

import csv
import sqlite3
import polars as pl

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    con = sqlite3.connect("dataset.sqlite")
    rows = []
    logger.info(f"reading csv. elapsed={time.time() - start}")
    start = time.time()
    with open(path, "r") as f:
        reader = csv.reader(f, delimiter=";")
        for row in reader:
            rows.append((row[0], float(row[1])))
    cur = con.cursor()
    cur.execute("CREATE TABLE dataset(location text, temperature double)")
    logger.info(f"dumping output. elapsed={time.time() - start}")
    start = time.time()
    cur.executemany("INSERT INTO dataset VALUES(?, ?)", rows)
    con.commit()
    con.close()
    return start


def process_dataset(start: float) -> float:
    logger.info(f"processing dataset. elapsed: {time.time() - start}")
    start = time.time()
    df = pl.read_database_uri(
        """
    SELECT location,
           avg(temperature) as temperature_mean,
           max(temperature) as temperature_max,
           min(temperature) as temperature_min 
    FROM dataset
    GROUP BY location""",
        "sqlite://dataset.sqlite",
    )
    logger.info(f"dumping output. elapsed: {time.time() - start}")
    start = time.time()
    df.write_parquet("result_sq.parquet")
    return start
