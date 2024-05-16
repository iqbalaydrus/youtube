import argparse
import logging

handler = logging.StreamHandler()
handler.setLevel(logging.INFO)
handler.setFormatter(logging.Formatter("%(asctime)s %(message)s"))
logging.getLogger().setLevel(logging.INFO)
logging.getLogger().addHandler(handler)


def main():
    logging.getLogger().info("starting")
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--engine",
        choices=["duckdb", "sqlite", "pandas", "polars"],
        required=True,
    )
    parser.add_argument(
        "--stage",
        choices=["load"],
        required=True,
    )
    args = parser.parse_args()
    dataset_path = "measurements.txt"
    if args.engine == "duckdb":
        if args.stage == "load":
            from cases import c_duckdb as engine

            engine.load_dataset(dataset_path)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "sqlite":
        if args.stage == "load":
            from cases import c_sqlite as engine

            engine.load_dataset(dataset_path)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "pandas":
        if args.stage == "load":
            from cases import c_pandas as engine

            engine.load_dataset(dataset_path)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    elif args.engine == "polars":
        if args.stage == "load":
            from cases import c_polars as engine

            engine.load_dataset(dataset_path)
        else:
            raise ValueError(f"Unknown stage: {args.stage}")
    else:
        raise ValueError(f"Unknown engine {args.engine}")


if __name__ == "__main__":
    main()
