# API Documentation

### Auth

```bash
curl --location --request POST 'localhost:6363/v1/auth' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "email@example.com",
    "password": "pasw0rd"
}'
```

### Reset Password Start

```bash
curl --location --request POST 'localhost:6363/v1/auth/reset' \
--header 'Content-Type: application/json' \
--data-raw '{
    "login": "email@example.com"
}'
```

### Reset Password Finish

```bash
curl --location --request PUT 'localhost:6363/v1/auth/reset' \
--header 'Content-Type: application/json' \
--header 'X-Reset-Token: 65b234d6df5d35d2e5e6a9983bf2dfe33600d74d30bde0432c3ea4e9d29e0b45' \
--data-raw '{
    "password": "pasw0rd"
}'
```

### Segments List

```
curl --location --request GET 'localhost:6363/v1/segments' \
--header 'X-Auth-Token: 2c763f5bc41fa64f95b02993bb49969a067d8a1b1f4ce2a7f1963f8e6c06cc33'
```


### Get Segment

```
curl --location --request GET 'localhost:6363/v1/segments/1' \
--header 'X-Auth-Token: 2c763f5bc41fa64f95b02993bb49969a067d8a1b1f4ce2a7f1963f8e6c06cc33'
```

### Create Segment

```
curl --location --request POST 'localhost:6363/v1/segments' \
--header 'X-Auth-Token: 2c763f5bc41fa64f95b02993bb49969a067d8a1b1f4ce2a7f1963f8e6c06cc33' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Segment Name",
    "description": "Segment description",
    "condition": "parcel.weight > 5000"
}'
```

### Update Segment

```
curl --location --request PUT 'localhost:6363/v1/segments' \
--header 'X-Auth-Token: 2c763f5bc41fa64f95b02993bb49969a067d8a1b1f4ce2a7f1963f8e6c06cc33' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "name": "Segment Name updated",
    "description": "Segment description edited",
    "condition": "parcel.weight > 5001"
}'
```
