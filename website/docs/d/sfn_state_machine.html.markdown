---
subcategory: "SFN (Step Functions)"
layout: "aws"
page_title: "AWS: aws_sfn_state_machine"
description: |-
  Get information on an Amazon Step Function State Machine
---

# Data Source: aws_sfn_state_machine

Use this data source to get the ARN of a State Machine in AWS Step
Function (SFN). By using this data source, you can reference a
state machine without having to hard code the ARNs as input.

## Example Usage

```terraform
data "aws_sfn_state_machine" "example" {
  name = "an_example_sfn_name"
}
```

## Argument Reference

This data source supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `name` - (Required) Friendly name of the state machine to match.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `id` - Set to the ARN of the found state machine, suitable for referencing in other resources that support State Machines.
* `arn` - Set to the arn of the state function.
* `role_arn` - Set to the role_arn used by the state function.
* `definition` - Set to the state machine definition.
* `creation_date` - Date the state machine was created.
* `revision_id` - The revision identifier for the state machine.
* `status` - Set to the current status of the state machine.
