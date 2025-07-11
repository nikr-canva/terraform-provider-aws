---
subcategory: "Direct Connect"
layout: "aws"
page_title: "AWS: aws_dx_hosted_public_virtual_interface"
description: |-
  Provides a Direct Connect hosted public virtual interface resource.
---

# Resource: aws_dx_hosted_public_virtual_interface

Provides a Direct Connect hosted public virtual interface resource. This resource represents the allocator's side of the hosted virtual interface.
A hosted virtual interface is a virtual interface that is owned by another AWS account.

## Example Usage

```terraform
resource "aws_dx_hosted_public_virtual_interface" "foo" {
  connection_id = "dxcon-zzzzzzzz"

  name           = "vif-foo"
  vlan           = 4094
  address_family = "ipv4"
  bgp_asn        = 65352

  customer_address = "175.45.176.1/30"
  amazon_address   = "175.45.176.2/30"

  route_filter_prefixes = [
    "210.52.109.0/24",
    "175.45.176.0/22",
  ]
}
```

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `address_family` - (Required) The address family for the BGP peer. `ipv4 ` or `ipv6`.
* `bgp_asn` - (Required) The autonomous system (AS) number for Border Gateway Protocol (BGP) configuration.
* `connection_id` - (Required) The ID of the Direct Connect connection (or LAG) on which to create the virtual interface.
* `name` - (Required) The name for the virtual interface.
* `owner_account_id` - (Required) The AWS account that will own the new virtual interface.
* `route_filter_prefixes` - (Required) A list of routes to be advertised to the AWS network in this region.
* `vlan` - (Required) The VLAN ID.
* `amazon_address` - (Optional) The IPv4 CIDR address to use to send traffic to Amazon. Required for IPv4 BGP peers.
* `bgp_auth_key` - (Optional) The authentication key for BGP configuration.
* `customer_address` - (Optional) The IPv4 CIDR destination address to which Amazon should send traffic. Required for IPv4 BGP peers.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the virtual interface.
* `arn` - The ARN of the virtual interface.
* `aws_device` - The Direct Connect endpoint on which the virtual interface terminates.

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `create` - (Default `10m`)
- `delete` - (Default `10m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```terraform
import {
  to = aws_dx_hosted_public_virtual_interface.test
  id = "dxvif-33cc44dd"
}
```

Using `terraform import`, import Direct Connect hosted public virtual interfaces using the VIF `id`. For example:

```console
% terraform import aws_dx_hosted_public_virtual_interface.test dxvif-33cc44dd
```
