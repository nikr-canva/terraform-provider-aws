---
subcategory: "EC2 (Elastic Compute Cloud)"
layout: "aws"
page_title: "AWS: aws_ami_ids"
description: |-
  Provides a list of AMI IDs.
---

# Data Source: aws_ami_ids

Use this data source to get a list of AMI IDs matching the specified criteria.

## Example Usage

```terraform
data "aws_ami_ids" "ubuntu" {
  owners = ["099720109477"]

  filter {
    name   = "name"
    values = ["ubuntu/images/ubuntu-*-*-amd64-server-*"]
  }
}
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `owners` - (Required) List of AMI owners to limit search. At least 1 value must be specified. Valid values: an AWS account ID, `self` (the current account), or an AWS owner alias (e.g., `amazon`, `aws-marketplace`, `microsoft`).
* `executable_users` - (Optional) Limit search to users with *explicit* launch
permission on  the image. Valid items are the numeric account ID or `self`.
* `filter` - (Optional) One or more name/value pairs to filter off of. There
are several valid keys, for a full reference, check out
[describe-images in the AWS CLI reference][1].
* `name_regex` - (Optional) Regex string to apply to the AMI list returned
by AWS. This allows more advanced filtering not supported from the AWS API.
This filtering is done locally on what AWS returns, and could have a performance
impact if the result is large. Combine this with other
options to narrow down the list AWS returns.
* `sort_ascending` - (Optional) Used to sort AMIs by creation time.
If no value is specified, the default value is `false`.
* `include_deprecated` - (Optional) If true, all deprecated AMIs are included in the response.
If false, no deprecated AMIs are included in the response. If no value is specified, the default value is `false`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `ids` is set to the list of AMI IDs, sorted by creation time according to `sort_ascending`.

[1]: http://docs.aws.amazon.com/cli/latest/reference/ec2/describe-images.html

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)
