# flare

Ultra-fast Cloudflare IP updater for nginx `real_ip` configuration. Keeps your Cloudflare real IP ranges current without you having to worry about them.

## Install

```bash
curl -sSL https://raw.githubusercontent.com/coalaura/flare/master/install.sh | sudo bash
```

## Usage

```bash
flare       # Update /etc/nginx/flare.conf and reload nginx
flare --dry # Print config without writing
```

## What it does

- Fetches current Cloudflare IP ranges (IPv4 + IPv6)
- Generates nginx `set_real_ip_from` directives
- Only updates if IPs changed (smart diff)
- Automatically reloads nginx
- Runs in ~80ms with zero-allocation optimizations

## Performance
```bash
$ hyperfine './flare' --runs=100
Benchmark 1: ./flare
  Time (mean ± σ):      69.9 ms ±  10.7 ms    [User: 20.8 ms, System: 9.5 ms]
  Range (min … max):    52.8 ms … 101.8 ms    100 runs
```