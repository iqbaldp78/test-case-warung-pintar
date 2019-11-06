# TestCase Warung Pintar

---


---

## Requirement

- [Go 1.12 or higher](https://golang.org/dl/)
- [docker](https://www.docker.com/)
- [Gin-Gonic Framework](https://github.com/gin-gonic/gin)

## Overview 
This Repo uses websocket for implementing retrive message with long live connection. The websocket uses `filewatch` txt file to listen and show the responses based on browser.


There is 4 endpoint :

| Endpoint | Method |Description
| ------ | ----------- | ------ |
| /sample   | GET |this endpoint used for generate response from server and save to txt file.
| /show| GET |this endpoint used as show all response.
| /ws    | GET |this endpoint used for handling websocket. 
| /index    | GET |this endpoint used for generate index.html to show the response. 


## How to trial 

- clone this repo and put outside your gopath
```bash
git clone https://github.com/iqbaldp78/test-case-warung-pintar.git
```
- using docker

```shell
cd test-case-warung-pintar

docker build --rm -f "Dockerfile" -t test-case-warung-pintar:latest .

docker run -d -it -p 8080:8080 test-case-warung-pintar:latest
```

- run manually
```shell
cd test-case-warung-pintar

go run main.go
```

note :
if port 8080 is already used on your computer, try to changing another port, e.g. 8070


if you want to try websocket

1. open your browser and go to http://localhost:8080/index
2. hit /sample endpoint
3. and the browser will showing the response realtime

<a name="top"></a>
# API Doc

- [sample](#sample)
- [show](#show)
- [ws](#ws)
- [index](#index)
	


# <a name='sample'></a> sample
[Back to top](#top)

<p>sample</p>
this endpoint used for generate response from server and save to txt file (output_response.txt).

	GET /sample?message=ping





### Parameters URL

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  message | String (Required) | <p>message will be send to response</p>|

### Examples

request body {curl} Example usage:

```
curl -i http://localhost:8080/sample?message=ping
```


### Success Response

Success-Response:

```
HTTP/1.1 200 Ok
{
    "message": "ping"
}
```

### Success 200

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  message | String | <p>output from server based on parameter URL</p>|


### Error Response

Error-Response:

```
HTTP/1.1 400 Bad Request
{
	 "message": "Paramater Not Suitable",
	 "code": "0XX"
}
```
Error-Response:

```
HTTP/1.1 500 Internal Server Error
{
	 "message": "Something wrong when get data",
	 "code": "0XX"
}
```



# <a name='show'></a> show
[Back to top](#top)

<p>show</p>
this endpoint used as show all response from txt file (output_response.txt).

	GET /show






### Examples

request body {curl} Example usage:

```
curl -i http://localhost:8080/show
```


### Success Response

Success-Response:

```
HTTP/1.1 200 Ok
{
    "total": 1,
    "data": [
        "ping"
    ]
}
```

### Success 200

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  total | Int | <p>Total data response</p>|
|  data | []String | <p>Show all responses</p>|


### Error Response

Error-Response:

```
HTTP/1.1 400 Bad Request
{
	 "message": "Paramater Not Suitable",
	 "code": "0XX"
}
```
Error-Response:

```
HTTP/1.1 500 Internal Server Error
{
	 "message": "Something wrong when get data",
	 "code": "0XX"
}
```


# <a name='ws'></a> ws
[Back to top](#top)

<p>ws</p>

	GET ws://localhost:8080/ws

this websocket endpoint will be trigger when you run ping/index endpoint

# <a name='index'></a> index
[Back to top](#top)

<p>index</p>

	GET http://localhost:8080/index

this endpoint will be show html page
