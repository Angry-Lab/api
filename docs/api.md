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
