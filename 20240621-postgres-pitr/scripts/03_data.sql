INSERT INTO pitr_data(val)
SELECT random() * 10 AS i
FROM generate_series(0, 4)
