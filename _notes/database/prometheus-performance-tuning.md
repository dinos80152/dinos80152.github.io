# Prometheus Performance Tuning

## Profiling

### Overview

<http://your.prometheus.host:9090/debug/pprof>

### Go Profiling

1. go pprof <http://your.prometheus.host:9090/debug/pprof/heap>
1. go pprof web

## System

### Adjust nofile

1. find the current ulimit number of open file

    ```sh
    ulimit -a
    ```

1. increase open files (file descriptor) limit

    ```sh
    ulimit -n 102400
    ```

## Startup Flag

adjust query flag

- `query.timeout`
- `query.max-samples`
- `query.max-concurrency`

## Federation

## Thanos

### Reference

[Why does Prometheus consume so much memory?@stackoverflow](https://stackoverflow.com/a/56142573/4810199)
