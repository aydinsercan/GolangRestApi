## Rest API Service in GOLANG

A rest-api that works with golang as an in-memory key value store

## Usage

Run command below in terminal in project directory.

```shell
go run main.go
```
Docker Container
```shell
docker build -t go-api .
docker run -p 8888:8888 go-api
```

Now, you can send requests to the application via postman according to the following endpoints

+ Get
    + Get returns the value corresponding to 'key' in DB.
    + GET 'localhost:8888/get/1'
+ Add
    + stores the key value in the JSON that comes with the POST request in the inMemDB key value format.
    + POST 'localhost:8888/add?key=sercanaydin'
+ Flush
    + Flush deletes all data in DB.
    + PUT 'localhost:8888/flush'

## License
[MIT](https://choosealicense.com/licenses/mit/)