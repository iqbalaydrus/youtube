import logging
import time

import pandas as pd
import numpy as np

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    logger.info(f"reading csv. elapsed={time.time() - start}")
    start = time.time()
    df = pd.read_csv(
        path,
        sep=";",
        names=["location", "temperature"],
        dtype={"location": str, "temperature": np.float64},
        engine="pyarrow",  # the only engine that has multithreading support
    )
    logger.info(f"dumping output. elapsed={time.time() - start}")
    start = time.time()
    df.to_parquet("dataset_pd.parquet")
    return start


def process_dataset(start: float) -> float:
    logger.info(f"reading dataset. elapsed={time.time() - start}")
    start = time.time()
    df = pd.read_parquet("dataset_pd.parquet")
    logger.info(f"processing dataset. elapsed: {time.time() - start}")
    start = time.time()
    df = df.groupby(["location"]).agg(
        temperature_mean=pd.NamedAgg(column="temperature", aggfunc="mean"),
        temperature_max=pd.NamedAgg(column="temperature", aggfunc="max"),
        temperature_min=pd.NamedAgg(column="temperature", aggfunc="min"),
    )
    logger.info(f"dumping output. elapsed: {time.time() - start}")
    start = time.time()
    df.to_parquet("result_pd.parquet")
    return start
