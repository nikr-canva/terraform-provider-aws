---
subcategory: "API Gateway V2"
layout: "aws"
page_title: "AWS: aws_apigatewayv2_route_response"
description: |-
  Manages an Amazon API Gateway Version 2 route response.
---

# Resource: aws_apigatewayv2_route_response

Manages an Amazon API Gateway Version 2 route response.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

## Example Usage

### Basic

```terraform
resource "aws_apigatewayv2_route_response" "example" {
  api_id             = aws_apigatewayv2_api.example.id
  route_id           = aws_apigatewayv2_route.example.id
  route_response_key = "$default"
}
```

## Enabling Two-Way Communication

For websocket routes that require two-way communication enabled, an `aws_apigatewayv2_route_response` needs to be added to the route with `route_response_key = "$default"`. More information available  is available in [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

You can only define the $default route response for WebSocket APIs. You can use an integration response to manipulate the response from a backend service. For more information, see [Overview of integration responses](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-integration-responses.html#apigateway-websocket-api-integration-response-overview).

## Argument Reference

This resource supports the following arguments:

* `region` - (Optional) Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
* `api_id` - (Required) API identifier.
* `route_id` - (Required) Identifier of the [`aws_apigatewayv2_route`](/docs/providers/aws/r/apigatewayv2_route.html).
* `route_response_key` - (Required) Route response key.
* `model_selection_expression` - (Optional) The [model selection expression](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.html#apigateway-websocket-api-model-selection-expressions) for the route response.
* `response_models` - (Optional) Response models for the route response.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - Route response identifier.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import `aws_apigatewayv2_route_response` using the API identifier, route identifier and route response identifier. For example:

```terraform
import {
  to = aws_apigatewayv2_route_response.example
  id = "aabbccddee/1122334/998877"
}
```

Using `terraform import`, import `aws_apigatewayv2_route_response` using the API identifier, route identifier and route response identifier. For example:

```console
% terraform import aws_apigatewayv2_route_response.example aabbccddee/1122334/998877
```
