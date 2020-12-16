# Example Terraform Provider

Want to check out a provider that uses the v2 Terraform SDK? So did I!

This provider is an interface to an example API that tracks your pets. The API is described in
an [OpenAPI](https://swagger.io/specification/) spec `example-server/swagger.yml`. The server and client are both
generated using the `swagger` tool.

## Requirements

- Terraform 0.13+
- Go 1.15+
- Swagger v0.25.0+

## Build

```shell
make
```

## Try out an example

In shell #1, start the server

```shell
make run
```

In shell #2, use the provider

```shell
cd tf-examples/example-pet
terraform init
terraform plan
terraform apply
terraform destroy
```
