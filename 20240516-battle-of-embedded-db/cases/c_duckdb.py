import duckdb
import os.path


def load_dataset(path: str):
    if os.path.exists("dataset.duckdb"):
        os.remove("dataset.duckdb")
    con = duckdb.connect("dataset.duckdb")
    con.execute("CREATE TABLE dataset(location text, temperature double)")
    con.execute(
        """
INSERT INTO dataset(location, temperature)
SELECT * FROM read_csv(?,
    delim = ';',
    header = false,
    parallel = true,
    names = ['location', 'temperature'],
    dtypes = ['text', 'double']);
""",
        [
            path,
        ],
    )
    con.close()
