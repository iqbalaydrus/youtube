# How To Run

## Generating Measurements
- You must have Java 21 on your system, verify by running `java -version` on your command line
- Clone Official 1 Billion Row Challenge [Repo](https://github.com/gunnarmorling/1brc)
- Run `./mvnw clean verify`
- Run `./create_measurements.sh 100000000`
- Copy `measurements.txt` to `youtube/20240516-battle-of-embedded-db/` path

## Running The Benchmark
- python 3.12 is recommended
- Install the requirements `pip install -r requirements.txt`
- Run the benchmark `./run.sh`