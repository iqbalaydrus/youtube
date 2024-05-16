import logging

import polars as pl
import os.path


def load_dataset(path: str):
    if os.path.exists("dataset_pl.parquet"):
        os.remove("dataset_pl.parquet")
    logger = logging.getLogger()
    logger.info("reading csv")
    df = pl.read_csv(
        path,
        has_header=False,
        separator=";",
        dtypes={"location": pl.String, "temperature": pl.Float64},
        new_columns=["location", "temperature"],
    )
    logger.info("dumping output")
    df.write_parquet("dataset_pl.parquet")
    logger.info("exit")
