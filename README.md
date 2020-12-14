# Example Terraform Provider

The point of this provider is to use the v2 provider sdk.

Generate client: 
```
cd client/
swagger generate client -f ../example-server/swagger.yml -A pet
```

Generate example-server:
```
cd example-server
swagger generate server -A pet -f ./swagger.yml
```