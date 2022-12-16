# Backend Service

![Go version](https://img.shields.io/badge/Go-v1.19-lightblue)
![Docker version](https://img.shields.io/badge/docker--version-20.10.21-blue)

Service atau layanan ini akan menangani logic dan Database storage untuk aplikasi Peminjaman Ruangan. Berperan sebagai API Gateway yang menyediakan endpoint-endpoint yang dapat FrontEnd gunakan. 

## API Documentation

Berikut beberapa dokumentasi API yang dapat kami tampilkan:

### Root Endpoint

Sample request

```sh
curl --location --request GET 'https://peminjamanruangan.rtln.xyz'
```

Sample response

```json
{
    "success": true,
    "error": null,
    "msg": "Server listening with version v1.0.0",
    "data": null
}
```

### Register Endpoint

Sample request

```sh
curl --location --request POST 'https://peminjamanruangan.rtln.xyz/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"paulj",
    "password":"passw0rdkuat",
    "real_name": "Paul J"
}'
```

Sample response

```json
{
    "data": null,
    "error": null,
    "msg": "success",
    "success": true
}
```

### Login Endpoint

Sample request

```sh
curl --location --request POST 'https://peminjamanruangan.rtln.xyz/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"paulj",
    "password":"passw0rdkuat"
}'
```

Sample response

```json
{
    "data": {
        "username": "paulj",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfYXQiOjE2NzE0MjUwOTksImlzc3VlZF9hdCI6MTY3MTIwOTA5OSwiaXNzdWVyIjoicGVtaW5qYW1hbi1ydWFuZ2FuIiwicm9sZV9pZCI6Miwicm9sZV9uYW1lIjoidXNlciIsInVzZXJuYW1lIjoicGF1bGoifQ.xWLaB43R_a0W1HBnX0BsXEJzMuBGd4tCyyM6_0nP2Ow"
    },
    "error": null,
    "msg": "success",
    "success": true
}   
```

---

_Note: to get full access of the API you need to request the main contributor via email here at [me@ariebrainware.com](mailto:me@ariebrainware.com)_

### Hosted API URL

https://peminjamanruangan.rtln.xyz

## Contributor

- 1922021 Muhamad Arie

