import logging
import time

import duckdb

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    con = duckdb.connect("dataset.duckdb")
    logger.info(f"reading csv. elapsed={time.time() - start}")
    start = time.time()
    con.execute("CREATE TABLE dataset(location text, temperature double)")
    con.execute(
        """
INSERT INTO dataset(location, temperature)
SELECT * FROM read_csv(?,
    delim = ';',
    header = false,
    parallel = true,
    names = ['location', 'temperature'],
    dtypes = ['text', 'double']);
""",
        [
            path,
        ],
    )
    logger.info(f"dumping output. elapsed={time.time() - start}")
    start = time.time()
    con.close()
    return start


def process_dataset(start: float) -> float:
    logger.info(f"reading dataset. elapsed={time.time() - start}")
    start = time.time()
    con = duckdb.connect("dataset.duckdb")
    logger.info(f"processing dataset. elapsed: {time.time() - start}")
    start = time.time()
    con.execute(
        """COPY (
    SELECT location,
           avg(temperature) as temperature_mean,
           max(temperature) as temperature_max,
           min(temperature) as temperature_min 
    FROM dataset
    GROUP BY location
    ORDER BY location
    ) TO 'result_dk.parquet' (FORMAT PARQUET)"""
    )
    con.close()
    return start
