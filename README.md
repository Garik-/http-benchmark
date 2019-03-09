# node.js vs golang
Benchmark simple http server
## System resources
| | |
--- | ---
Processor|2,2 GHz Intel Core i7
Memory|16 ГБ 2400 MHz DDR4
## Load testing
```
$ ab -n 10000 -kc 100 http://localhost:3000/ > node.txt
$ ab -n 10000 -kc 100 http://localhost:9000/ > go.txt
```
## Results
|Language|Requests per second|
--- | ---
golang|2586.09
node.js|2513.90
