import logging
import time
import os.path

import pandas as pd
import numpy as np

logger = logging.getLogger()


def load_dataset(start: float, path: str) -> float:
    if os.path.exists("dataset_pd.parquet"):
        os.remove("dataset_pd.parquet")
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
