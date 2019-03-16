# node.js vs golang
Benchmark simple http server
## System resources
| | |
--- | ---
Processor|2,2 GHz Intel Core i7
Memory|16 ГБ 2400 MHz DDR4
## Load testing
### ab
```
$ ab -n 10000 -kc 100 http://localhost:3000/ > node.txt
$ ab -n 10000 -kc 100 http://localhost:9000/ > go.txt
```
|Language|Requests per second|
--- | ---
golang|2586.09
node.js|2513.90
### Jmeter
100 users

|Language|Samples|Throughput|
--- | --- | --- |
node.js with uWS|32726|894
node.js|32737|895.8
golang|28663|1026.2

## Alternative load testing
### ab
```
$ ab -n 100000 -kc 100 http://localhost:3000/
```
|Language|Data size|RPS|
--- | --- | --- |
node.js|big|6197
go stdlib|big|6181
go fasthttp|big|7183
node.js|small|13124
go stdlib|small|24377
go fasthttp|small|84639

### wrk
```
$ wrk -c 100 -d 10 -t 2 http://localhost:3000/
```
|Language|Data size|RPS|
--- | --- | --- |
node.js|big|9642
go stdlib|big|11520
go fasthttp|big|14071
node.js|small|24677
go stdlib|small|46439
go fasthttp|small|95183
