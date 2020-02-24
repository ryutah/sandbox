# cloud-run-helloworld

## Performance

### Default

```console
ab -n 1000 -c 10 'https://[CLOUD_RUN_SERVICE].run.app/'
```

```txt
Server Software:        Google
Server Hostname:        cloud-run-started-hbl5nasmva-an.a.run.app
Server Port:            443
SSL/TLS Protocol:       TLSv1.2,ECDHE-RSA-CHACHA20-POLY1305,2048,256
Server Temp Key:        ECDH X25519 253 bits
TLS Server Name:        cloud-run-started-hbl5nasmva-an.a.run.app

Document Path:          /
Document Length:        3 bytes

Concurrency Level:      10
Time taken for tests:   11.148 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      381008 bytes
HTML transferred:       3000 bytes
Requests per second:    89.70 [#/sec] (mean)
Time per request:       111.477 [ms] (mean)
Time per request:       11.148 [ms] (mean, across all concurrent requests)
Transfer rate:          33.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       68   89  12.2     87     171
Processing:    14   21   6.6     19     149
Waiting:       13   20   5.6     19     104
Total:         87  109  14.3    107     307

Percentage of the requests served within a certain time (ms)
  50%    107
  66%    110
  75%    113
  80%    114
  90%    120
  95%    129
  98%    148
  99%    177
 100%    307 (longest request)
```

## With Stackdriver Trace

```txt
Server Software:        Google
Server Hostname:        cloud-run-started-hbl5nasmva-an.a.run.app
Server Port:            443
SSL/TLS Protocol:       TLSv1.2,ECDHE-RSA-CHACHA20-POLY1305,2048,256
Server Temp Key:        ECDH X25519 253 bits
TLS Server Name:        cloud-run-started-hbl5nasmva-an.a.run.app

Document Path:          /
Document Length:        3 bytes

Concurrency Level:      10
Time taken for tests:   11.490 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      381004 bytes
HTML transferred:       3000 bytes
Requests per second:    87.03 [#/sec] (mean)
Time per request:       114.898 [ms] (mean)
Time per request:       11.490 [ms] (mean, across all concurrent requests)
Transfer rate:          32.38 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       66   91  14.1     87     163
Processing:    14   23   9.4     20      94
Waiting:       13   22   9.3     20      94
Total:         83  113  19.8    108     233

Percentage of the requests served within a certain time (ms)
  50%    108
  66%    113
  75%    117
  80%    119
  90%    131
  95%    154
  98%    182
  99%    207
 100%    233 (longest request)
```

## With Stackdriver Trace And Profiler

```txt
Server Software:        Google
Server Hostname:        cloud-run-started-hbl5nasmva-an.a.run.app
Server Port:            443
SSL/TLS Protocol:       TLSv1.2,ECDHE-RSA-CHACHA20-POLY1305,2048,256
Server Temp Key:        ECDH X25519 253 bits
TLS Server Name:        cloud-run-started-hbl5nasmva-an.a.run.app

Document Path:          /
Document Length:        3 bytes

Concurrency Level:      10
Time taken for tests:   14.532 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      381068 bytes
HTML transferred:       3000 bytes
Requests per second:    68.81 [#/sec] (mean)
Time per request:       145.323 [ms] (mean)
Time per request:       14.532 [ms] (mean, across all concurrent requests)
Transfer rate:          25.61 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:       66   87  18.1     84     209
Processing:    13   53 173.8     21    1784
Waiting:       13   53 173.8     20    1784
Total:         82  140 173.7    105    1869

Percentage of the requests served within a certain time (ms)
  50%    105
  66%    110
  75%    114
  80%    118
  90%    139
  95%    221
  98%    913
  99%   1226
 100%   1869 (longest request)
```

## Local Server

Run local server with [skaffold](https://github.com/GoogleContainerTools/skaffold).

```console
script/serve.sh
```

It can belows

1. see tracing with zipkin
    - Open `localhost:9411`
1. see profiling with `go tool pprof`
    - Open CPU Profiler

        ```console
        go tool pprof -http=":8081" http://localhost:6060/debug/pprof/profile
        ```

    - Open Heap Profiler

        ```console
        go tool pprof -http=":8081" http://localhost:6060/debug/pprof/heap
        ```

### Tips

#### Cleanup images after shutdown local server

```console
docker image rm $(docker image ls --filter 'dangling=true' -q)
```
