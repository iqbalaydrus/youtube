# How To Run
- Run `docker compose up -d` command
- Install & Run [vegeta](https://github.com/tsenart/vegeta) command to generate traffic to our demo instance and let it run
```shell
vegeta attack -targets vegeta/target.txt -duration 0 -output /dev/null
# or if you're feeling adventurous, generate 1000 request/second
vegeta attack -targets vegeta/target.txt -duration 0 -rate 1000/s -output /dev/null
```
- Open grafana dashboard on your browser `http://127.0.0.1:8878`
- Don't forget to stop vegeta command after done exploring