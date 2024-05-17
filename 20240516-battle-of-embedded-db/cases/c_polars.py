import logging
import time

import polars as pl

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
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


def process_dataset(start: float) -> float:
    logger.info(f"reading dataset. elapsed={time.time() - start}")
    start = time.time()
    df = pl.read_parquet("dataset_pl.parquet")
    logger.info(f"processing dataset. elapsed: {time.time() - start}")
    start = time.time()
    df = (
        df.groupby("location")
        .agg(
            pl.mean("temperature").name.suffix("_mean"),
            pl.max("temperature").name.suffix("_max"),
            pl.min("temperature").name.suffix("_min"),
        )
        .sort("location")
    )
    logger.info(f"dumping output. elapsed: {time.time() - start}")
    start = time.time()
    df.write_parquet("result_pl.parquet")
    return start
