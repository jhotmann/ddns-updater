# deSEC

## Configuration

### Example

```json
{
  "settings": [
    {
      "provider": "desec",
      "domain": "dedyn.io",
      "host": "host",
      "token": "token",
      "ip_version": "ipv4",
      "ipv6_suffix": "",
      "provider_ip": false
    }
  ]
}
```

### Compulsory parameters

- `"domain"`
- `"host"` can be `@` for the root domain or a subdomain or a wildcard subdomain (`*`), defaults to `@`
- `"token"` is your token that you can create [here](https://desec.io/tokens)

### Optional parameters

- `"ip_version"` can be `ipv4` (A records), or `ipv6` (AAAA records) or `ipv4 or ipv6` (update one of the two, depending on the public ip found). It defaults to `ipv4 or ipv6`.
- `"ipv6_suffix"` is the IPv6 interface identifier suffix to use. It can be for example `0:0:0:0:72ad:8fbb:a54e:bedd/64`. If left empty, it defaults to no suffix and the raw public IPv6 address obtained is used in the record updating.
- `"provider_ip"` can be set to `true` to let your DNS provider determine your IPv4 address (and/or IPv6 address) automatically when you send an update request, without sending the new IP address detected by the program in the request.

## Domain setup

[desec.io/domains](https://desec.io/domains)
