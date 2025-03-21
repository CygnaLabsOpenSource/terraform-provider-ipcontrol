# Terraform Resource: IPC Address

## Overview

The `ipcontrol_address` data source retrieves information about a device managed by IPControl.

## Parameters

### Required Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `ip_address` | `string` | The IP address of the device. |

### Optional Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `container` | `string` | The name of the container that contains the device. **Required** only if the IP address is not unique. |
| `rawcontainer` | `boolean` | Pass the container parameter through to the API without prefixing.    |

### Computed Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| `id` | `string` | Unique identifier for the device. |
| `ip_address` | `string` | The IP address of the device. |
| `address_type` | `string` | The address type of this interface IP address. |
| `resource_record_flag` | `string` | Resource record status. |
| `aliases` | `list(string)` | Device aliases. |
| `domain_name` | `string` | The domain name of the device. |
| `hostname` | `string` | The host name of the device. |
| `device_type` | `string` | The type of the device. |
| `domain_type` | `string` | The type of the domain. |
| `duid` | `string` | The DHCP unique identifier for IPv6. |
| `interfaces` | `list` | Network interfaces for the device. |

### Interface Configuration

Each interface in the `interfaces` list supports:

| Parameter | Type | Description |
|-----------|------|-------------|
| `address_type` | `list(string)` | Type of IP address (Computed) |
| `ip_address` | `list(string)` | IP addresses for the interface (Computed) |
| `container` | `list(string)` | Container information (Computed) |
| `name` | `string` | Interface name (Computed) |
| `id` | `integer` | Interface identifier (Computed) |


## Example Usage

#### Device with IPv4
```hcl
data "ipcontrol_address" "my_device" {
  ip_address = "192.168.29.111"
}
```
#### Device with IPv6
```hcl
data "ipcontrol_address" "my_device" {
  ip_address = "2001:db8:85a3::3000:0"
}
```