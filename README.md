# TestCase Warung Pintar

---


---

## Requirement

- [Go 1.12 or higher](https://golang.org/dl/)
- [docker](https://www.docker.com/)
- [Gin-Gonic Framework](https://github.com/gin-gonic/gin)

## Overview 

There is 4 endpoint for this repo :

| Endpoint | Method |Description
| ------ | ----------- | ------ |
| /ping/server   | GET |this endpoint used for generate response from server.
| /ping/client | GET |this endpoint used as client and request message to server. Handlebars is the default.
| /ping/ws    | GET |this endpoint used for handling websocket. 
| /index    | GET |this endpoint used for generate idex.html to trigger websocket. 


## How to trial 

- clone this repo and put outside your gopath
```bash
git clone https://github.com/iqbaldp78/test-case-warung-pintar.git
```
- start docker

```shell
cd test-case-warung-pintar

docker build --rm -f "dockerfile" -t test-case-warung-pintar:latest .

docker run -d -it -p 8080:80 test-case-warung-pintar:latest
```

note :
if port 8080 is already used on your computer, try to changing another port, e.g. 8070


<a name="top"></a>
# API Doc


- [ping/server](#ping/server)
- [ping/client](#ping/client)
- [ping/ws](#ping/ws)
- [ping/index](#ping/index)
	


# <a name='ping/server'></a> ping/server
[Back to top](#top)

<p>ping/server</p>

	GET /ping/server?message=ping





### Parameters URL

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  message | String | <p>message will be send</p>|

### Examples

request body {curl} Example usage:

```
curl -i http://localhost:8080//ping/server?message=ping
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



# <a name='ping/client'></a> ping/client
[Back to top](#top)

<p>ping/client</p>

	GET /ping/client?message=ping





### Parameters URL

| Name     | Type       | Description                           |
|:---------|:-----------|:--------------------------------------|
|  message | String | <p>message will be send to server</p>|

### Examples

request body {curl} Example usage:

```
curl -i http://localhost:8080//ping/server?message=ping
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
|  message | String | <p>Response from server</p>|


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


# <a name='ping/ws'></a> ping/ws
[Back to top](#top)

<p>ping/ws</p>

	GET ws://localhost:8080/ping/ws

this websocket endpoint will be trigger when you run ping/index endpoint

# <a name='ping/index'></a> ping/index
[Back to top](#top)

<p>ping/index</p>

	GET http://localhost:8080/ping/index

this endpoint will be show html page
