# schop

`schop` is a Command line tool for searching IPv4/IPv6 address from IPv6/IPv4 address.

For example, when you want to know google public DNS's IPv6 address -- that is, you know only IPv4 address but you don't know IPv6 address -- this script is fine for you.

You can search IPv6 address of this host from IPv4 address by executing this script.

```bash
$ schop 8.8.8.8
```

The result is below.

```
{"fqdn":"google-public-dns-a.google.com.","ipv4":"8.8.8.8","ipv6":"2001:4860:4860::8888"}
```

## Author
Kohei Suzuki <jingle@sfc.wide.ad.jp>

## License
MIT
