import argparse
import logging
import time

handler = logging.StreamHandler()
handler.setLevel(logging.INFO)
handler.setFormatter(logging.Formatter("%(asctime)s %(message)s"))
logger = logging.getLogger()
logger.setLevel(logging.INFO)
logger.addHandler(handler)


def main():
    logger.info("starting")
    total_start = time.time()
    start = time.time()
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--engine",
        choices=["duckdb", "sqlite", "pandas", "polars"],
        required=True,
    )
    parser.add_argument(
        "--stage",
        choices=["load", "process"],
        required=True,
    )
    args = parser.parse_args()
    dataset_path = "measurements.txt"
    if args.engine == "duckdb":
        if args.stage == "load":
            from cases import c_duckdb as engine

            start = engine.load_dataset(start, dataset_path)
        elif args.stage == "process":
            from cases import c_duckdb as engine

            start = engine.process_dataset(start)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "sqlite":
        if args.stage == "load":
            from cases import c_sqlite as engine

            start = engine.load_dataset(start, dataset_path)
        elif args.stage == "process":
            from cases import c_sqlite as engine

            start = engine.process_dataset(start)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "pandas":
        if args.stage == "load":
            from cases import c_pandas as engine

            start = engine.load_dataset(start, dataset_path)
        elif args.stage == "process":
            from cases import c_pandas as engine

            start = engine.process_dataset(start)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "polars":
        if args.stage == "load":
            from cases import c_polars as engine

            start = engine.load_dataset(start, dataset_path)
        elif args.stage == "process":
            from cases import c_polars as engine

            start = engine.process_dataset(start)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    else:
        raise ValueError(f"Unknown engine {args.engine}")
    logger.info(
        f"exit. elapsed={time.time() - start} total={time.time() - total_start}"
    )


if __name__ == "__main__":
    main()
