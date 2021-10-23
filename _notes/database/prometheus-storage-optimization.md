# Prometheus Storage Optimization

## Storage Size Evaluation

```math
needed_disk_space = retention_time_seconds * ingested_samples_per_second * bytes_per_sample (average 1-2 bytes per sample)
```

## Data Structure

The current block for incoming samples is kept in memory and is not fully persisted.

- ./data
  - 01BKGTZQ1SYQJTR4PB43C8PD98 (Ingested samples are grouped into blocks of two hours)
    - chunks (containing all the time series samples for that window of time)
      - 000001 (segment files of up to 512MB)
    - tombstones (When series are deleted via the API, deletion records are stored here)
    - index (indexes metric names and labels to time series in the chunks directory)
    - meta.json
  - chunks_head
    - 000001
  - wal (keep at least two hours of raw data; replayed when the Prometheus server restarts; has not yet been compacted)
    - 000000002 (stored in the wal directory in 128MB segments;retain a minimum of three write-ahead log files)
    - checkpoint.00000001
      - 00000000

> take up to two hours to remove expired blocks

## Find largest volume metrics

- web: <http://your.prometheus.host:9090/tsdb-status>
- query: `topk(10, count({__name__=~".+"}) by (__name__))`

## Delete metrics

<https://prometheus.io/docs/prometheus/latest/querying/api/#delete-series>

```sh
$ curl -X POST \
  -g 'http://your.prometheus.host:9090/api/v1/admin/tsdb/delete_series? \
  match[]=up& \
  match[]=http_request{status="200"}& \
  start=1595521900& \
  end=1595521990'
```

## Use relabel to reduce metrics

1 target could have N metrics

- reduce target by `relabel_configs`
- reduce metrics from target by `source_labels` in `metric_relabel_configs`
- reduct label from metrics by directly use `metric_relabel_configs`

```yaml
scrape_configs:
  - job_name: 'node'
    ec2_sd_configs:
      - region: us-west-2
        port: 8080
        # only discover tag Name is myapp
        filters:
          - name: "tag:name"
            values:
              - "myapp"
    relabel_configs:
      # Drop all metrics from windows server
      - source_labels: [__meta_ec2_platform]
        regex: windows
        action: drop
    metric_relabel_configs:
    # keep only http_request with label status 20X metric
    - source_labels: [__name__, status]
      regex: 'http_request;20.'
      action: keep
    # drop all source ip label
    - regex: 'source_ip'
      action: labeldrop
```

> Care must be taken with `labeldrop` and `labelkeep` to ensure that metrics are still uniquely labeled once the labels are removed.

### Reference

- [STORAGE@Prometheus](https://prometheus.io/docs/prometheus/latest/storage/)
- [TSDB format@Prometheus](https://github.com/prometheus/prometheus/blob/release-2.29/tsdb/docs/format/README.md)
- [Dropping metrics at scrape time with Prometheus@Robust Perception](https://www.robustperception.io/dropping-metrics-at-scrape-time-with-prometheus)
- [relabel_configs vs metric_relabel_configs@Robust Perception](https://www.robustperception.io/relabel_configs-vs-metric_relabel_configs)
- [Prometheus relabeling tricks@Quiq](https://medium.com/quiq-blog/prometheus-relabeling-tricks-6ae62c56cbda)
- [relabel_config@Prometheus](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config)
