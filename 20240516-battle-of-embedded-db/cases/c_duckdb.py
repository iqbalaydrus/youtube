import logging
import time
import os.path

import duckdb

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    if os.path.exists("dataset.duckdb"):
        os.remove("dataset.duckdb")
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
