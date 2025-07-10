# Terraform Resource: IPC Domain

## Overview

The `ipcontrol_dns` resource manages domains in IPControl.

## Parameters

### Required Parameters

| Parameter     | Type     | Description             |
| ------------- | -------- | ----------------------- |
| `domain_name` | `string` | The name of the domain. |

### Optional Parameters

| Parameter             | Type     | Default     | Description    |
| --------------------- | -------- | ----------- | -------------- |
| `id`                  | `string` | **Computed**| The id of the domain.|
| `domain_type`         | `string` | `Default`   | Domain type already defined in IPControl.|
| `contact`             | `string` | **Computed**| The contact email address in dotted format. If not specified, a default contact will be formed by prepending `dnsadmin` to the domain name, e.g., `dnsadmin.example.com.`|
| `delegated`           | `bool`   | `true`      | Indicates this domain will be associated directly with a zone file. |
| `default_ttl`         | `string` | `86400`     | Default time to live (TTL) for the zone. Ignored if `managed` is false.|
| `derivative`          | `string` | `STANDARD`  | Specify the role of this domain. One of: `STANDARD`, `TEMPLATE`, `CATALOG`, or `ALIAS`. |
| `expire`              | `string` | `604800`    | Zone expire time. Ignored if `managed` is false. |
| `info_template`       | `string` | -           | The name of the information template associated with this domain. |
| `managed`             | `bool`   | `true`      | Indicates that this domain is fully defined in IPControl. |
| `negative_cache_ttl`  | `string` | `86400`     | Negative cache TTL. Ignored if `managed` is false. |
| `refresh`             | `string` | `10800`     | Zone refresh interval. Ignored if `managed` is false. |
| `retry`               | `string` | `3600`      | Zone retry interval. Ignored if `managed` is false. |
| `reverse`             | `bool`   | `false`     | Indicates this is a reverse `in-addr.arpa` domain. |
| `serial_number`       | `int`    | `1`         | Zone serial number. If `0` or not specified, defaults to `1`. Ignored if `managed` is false. |
| `template_domain`     | `string` | -           | Required if `derivative` is `ALIAS`. Refers to the name of the template domain. |
| `user_defined_fields` | `list`   | -           | User-defined fields associated with the domain as per `info_template`. Required if any UDF is marked required.|
| `local_rpz`           | `bool`   | `false`     | If `true`, this domain will contain RPZ (Response Policy Zone) rules. |
| `description`         | `string` | -           | Description of the domain. |
| `serialformat`        | `string` | **Computed**| The format of the serial number, if overridden at the domain level. |


## Example Usage

#### Domain Resource Record IPv4

```hcl
resource "ipcontrol_dns" "example" {
  domain_name         = "example.com"
  domain_type         = "Default"
  contact             = "admin.example.com."
  delegated           = true
  default_ttl         = "86400"
  derivative          = "STANDARD"
  expire              = "604800"
  info_template       = "Cisco"
  user_defined_fields = [
    "gatewayip=20.0.0.1",
    "DHCPServerip=192.0.0.1"
  ]
  managed             = true
  negative_cache_ttl  = "86400"
  refresh             = "10800"
  retry               = "3600"
  reverse             = false
  serial_number       = 1
  local_rpz           = false
  description         = "Main domain for production"
}
```