# Example.com

## Configuration

### Example

<!-- UPDATE THIS JSON EXAMPLE -->

```json
{
  "settings": [
    {
      "provider": "example",
      "domain": "domain.com",
      "host": "@",
      "username": "username",
      "password": "password",
      "ip_version": "ipv4",
      "ipv6_suffix": ""
    }
  ]
}
```

### Compulsory parameters

- `"domain"`
- `"host"` is your host and can be a subdomain or `"@"` or the wildcard `"*"`
- `"username"`
- `"password"`

<!-- UPDATE THIS IF NEEDED -->

### Optional parameters

- `"ip_version"` can be `ipv4` (A records), or `ipv6` (AAAA records) or `ipv4 or ipv6` (update one of the two, depending on the public ip found). It defaults to `ipv4 or ipv6`.
- `"ipv6_suffix"` is the IPv6 interface identifier suffix to use. It can be for example `0:0:0:0:72ad:8fbb:a54e:bedd/64`. If left empty, it defaults to no suffix and the raw public IPv6 address obtained is used in the record updating.

<!-- UPDATE THIS IF NEEDED -->

## Domain setup

<!-- FILL THIS UP WITH A FEW NUMBERED STEPS -->