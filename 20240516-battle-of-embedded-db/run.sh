#!/usr/bin/env bash
psrecord 'python main.py --engine duckdb --stage load' --include-children --plot plot-duckdb-load.png
psrecord 'python main.py --engine sqlite --stage load' --include-children --plot plot-sqlite-load.png
psrecord 'python main.py --engine polars --stage load' --include-children --plot plot-polars-load.png
psrecord 'python main.py --engine pandas --stage load' --include-children --plot plot-pandas-load.png
