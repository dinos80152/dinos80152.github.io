# Prometheus Scrape and Metrics

## Scrape

### Relabel

- [[Prometheus] Service Discovery & Relabel@小信豬的原始部落](https://godleon.github.io/blog/Prometheus/Prometheus-Relabel/)

## Metrics

### Counter

> A counter starts at 0, and is incremented. At each scrape Prometheus takes a sample of this state.

> what happens when the process restarts and the counter is reset to 0? <br>
`rate()` will automatically handle this. Any time a counter appears to decrease it'll be treated as though there was a reset to 0 right after the first data point.

#### Query

- rate()
  - returns per second values,
  - `rate(counter[5s])` is `(v5 - v1) / (t5 - t1)`
- increase()
  - returns the total across the time period.
  - `increase(counter[5s])` is exactly `rate(counter[5s]) * 5`
- irate()
  - only looks at the last two samples
  - useful for high precision graphs over short time frames.
  - `irate(counter[5s])` is `(v5 - v4) / (t5 - t4)`
- offset
  - `counter - counter offest 5s`
  - it can't handle reset situation
- resets()
  - how often the counter resets
  - useful for debugging.

#### Reference

- [How does a Prometheus Counter work?@Robust Perception](https://www.robustperception.io/how-does-a-prometheus-counter-work)
- [Do I understand Prometheus's rate vs increase functions correctly?@stackoverflow](https://stackoverflow.com/q/54494394/4810199)

### Useful Query Example

- <https://github.com/infinityworks/prometheus-example-queries/blob/master/README.md>
