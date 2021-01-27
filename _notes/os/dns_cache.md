# DNS Cache

## Cache Policy on different OS

### Windows

<https://superuser.com/a/80799>

> - if DNS zone TTL < MaxCacheEntryTtlLimit, then DNS TTL is used
> - if DNS zone TTL > MaxCacheEntryTtlLimit, then MaxCacheEntryTtlLimit is used

### Linux

<https://stackoverflow.com/a/11021207/4810199>
> On Linux (and probably most Unix), there is no OS-level DNS caching unless nscd is installed and running.

### macOS

<https://apple.stackexchange.com/questions/93949/how-long-does-the-dns-cache-last-in-mac-os-x>
> How long a DNS entry is cached typically depends on the TTL of that DNS record, which is configured by the DNS admin of the relevant hostname.
