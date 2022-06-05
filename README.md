# go-triangle-classification


## Run LocalStack
```
localstack start
```

## Validated table created

```
 aws --profile=localstack --endpoint-url=http://localhost:4566 --region=us-east-1  dynamodb list-tables
```


## Generate unit test with Mockgen

```
mockgen -source=[file].go -destination=[path to mocks package]/[mock_file].go -package mocks
```

e.g

```
 mockgen -source=./usecase/authentication/jwt_token.go -destination=./mocks/jwt_token_mock.go -package mocks
```

## Generate Swagger Documentation

In this project, We use the [Swaggo Gin](https://github.com/swaggo/gin-swagger). The API information will be in the main file and other information will be in the methods of the handler. Any change is needed to run the command below to generate the documentation files:

```
 swag init
```

Could be accessed via:
```url 
    http://localhost:8080/docs/index.html
```