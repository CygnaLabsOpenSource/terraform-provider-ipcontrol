# Terraform Data Source: IPC Domain

## Overview

The `ipcontrol_dns_domain` data source retrieves information about a DNS managed by IPControl.

## Parameters

### Required Parameters

| Parameter     | Type     | Description |
| ------------- | -------- | ----------- |
| `domain_name` | `string` | The name of the domain. |

### Computed Parameters

| Parameter             | Type     | Description    |
| --------------------- | -------- | -------------- |
| `id`                  | `string` | The id of the domain.|
| `domain_type`         | `string` | Domain type already defined in IPControl.|
| `contact`             | `string` | The contact email address in dotted format. If not specified, a default contact will be formed by prepending `dnsadmin` to the domain name, e.g., `dnsadmin.example.com.`|
| `delegated`           | `bool`   | Indicates this domain will be associated directly with a zone file. |
| `default_ttl`         | `string` | Default time to live (TTL) for the zone. Ignored if `managed` is false.|
| `derivative`          | `string` | Specify the role of this domain. One of: `STANDARD`, `TEMPLATE`, `CATALOG`, or `ALIAS`. |
| `expire`              | `string` | Zone expire time. Ignored if `managed` is false. |
| `info_template`       | `string` | The name of the information template associated with this domain. |
| `managed`             | `bool`   | Indicates that this domain is fully defined in IPControl. |
| `negative_cache_ttl`  | `string` | Negative cache TTL. Ignored if `managed` is false. |
| `refresh`             | `string` | Zone refresh interval. Ignored if `managed` is false. |
| `retry`               | `string` | Zone retry interval. Ignored if `managed` is false. |
| `reverse`             | `bool`   | Indicates this is a reverse `in-addr.arpa` domain. |
| `serial_number`       | `int`    | Zone serial number. If `0` or not specified, defaults to `1`. Ignored if `managed` is false. |
| `template_domain`     | `string` | Required if `derivative` is `ALIAS`. Refers to the name of the template domain. |
| `user_defined_fields` | `list`   | User-defined fields associated with the domain as per `info_template`. Required if any UDF is marked required.|
| `local_rpz`           | `bool`   | If `true`, this domain will contain RPZ (Response Policy Zone) rules. |
| `description`         | `string` | Description of the domain. |
| `serialformat`        | `string` | The format of the serial number, if overridden at the domain level. |



## Example Usage

#### Domain Data Source
```hcl
data "ipcontrol_dns_domain" "my_dns" {
  domain_name = "com."
}

output "dns" {
  value = data.ipcontrol_dns_domain.my_dns
}

```
