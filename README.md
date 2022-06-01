# go-triangle-classification


## Run LocalStack
```
localstack start
```

## Validated table created

```
 aws --profile=localstack --endpoint-url=http://localhost:4566 --region=us-east-1  dynamodb list-tables
```


## Genereted unit test with Mockgen

```
mockgen -source=[file].go -destination=[path to mocks package]/[mock_file].go -package mocks
```
e.g
```
 mockgen -source=./usecase/authentication/jwt_token.go -destination=./mocks/jwt_token_mock.go -package mocks
```
