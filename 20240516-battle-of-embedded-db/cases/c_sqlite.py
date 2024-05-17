import logging
import time
import os.path

import csv
import sqlite3

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    if os.path.exists("dataset.sqlite"):
        os.remove("dataset.sqlite")
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
