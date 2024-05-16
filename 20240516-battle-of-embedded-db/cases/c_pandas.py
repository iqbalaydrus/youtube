import pandas as pd
import numpy as np
import os.path


def load_dataset(path: str):
    if os.path.exists("dataset_pd.parquet"):
        os.remove("dataset_pd.parquet")
    df = pd.read_csv(
        path,
        sep=";",
        names=["location", "temperature"],
        dtype={"location": str, "temperature": np.float64},
        engine="pyarrow",  # the only engine that has multithreading support
    )
    df.to_parquet("dataset_pd.parquet")
