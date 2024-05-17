import logging
import time
import os.path

import polars as pl

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    if os.path.exists("dataset_pl.parquet"):
        os.remove("dataset_pl.parquet")
    logger.info(f"reading csv. elapsed={time.time() - start}")
    start = time.time()
    df = pl.read_csv(
        path,
        has_header=False,
        separator=";",
        dtypes={"location": pl.String, "temperature": pl.Float64},
        new_columns=["location", "temperature"],
    )
    logger.info(f"dumping output. elapsed: {time.time() - start}")
    start = time.time()
    df.write_parquet("dataset_pl.parquet")
    return start
