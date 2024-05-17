#!/usr/bin/env bash
rm -f dataset.sqlite* dataset.duckdb* dataset_*.parquet result_*.parquet plot-*-load.png plot-*-process.png
psrecord 'python main.py --engine duckdb --stage load' --include-children --plot plot-duckdb-load.png
psrecord 'python main.py --engine sqlite --stage load' --include-children --plot plot-sqlite-load.png
psrecord 'python main.py --engine polars --stage load' --include-children --plot plot-polars-load.png
psrecord 'python main.py --engine pandas --stage load' --include-children --plot plot-pandas-load.png
psrecord 'python main.py --engine duckdb --stage process' --include-children --plot plot-duckdb-process.png
psrecord 'python main.py --engine sqlite --stage process' --include-children --plot plot-sqlite-process.png
psrecord 'python main.py --engine polars --stage process' --include-children --plot plot-polars-process.png
psrecord 'python main.py --engine pandas --stage process' --include-children --plot plot-pandas-process.png
