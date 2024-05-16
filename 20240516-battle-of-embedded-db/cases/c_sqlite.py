import csv
import sqlite3
import os.path


def load_dataset(path: str):
    if os.path.exists("dataset.sqlite"):
        os.remove("dataset.sqlite")
    con = sqlite3.connect("dataset.sqlite")
    rows = []
    with open(path, "r") as f:
        reader = csv.reader(f, delimiter=";")
        for row in reader:
            rows.append((row[0], float(row[1])))
    cur = con.cursor()
    cur.execute("CREATE TABLE dataset(location text, temperature double)")
    cur.executemany("INSERT INTO dataset VALUES(?, ?)", rows)
    con.commit()
    con.close()
