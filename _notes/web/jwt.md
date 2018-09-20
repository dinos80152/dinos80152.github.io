# JWT - JSON Web Token

## Purpose

Store data in client side, which could be read by can't be modified.

* Authentication
* Authorization

## Structure

base64(header).base64(payload).signature

`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9`**.**`eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0`**.**`YaLyoBs8z5Va7YsIQaC6uEZDw8GZHBiV_2hIUSVQYUs`

### Header

```json
{
    "alg": "HS256", // algorithm
    "typ": "JWT" // type
}
```

### Payload

```json
{
    "iss": "Dino Lai", // issuer
    "sub": "dinos80152@gmail.com", // subject
    "aud": "dinolai.com", // audience
    "exp": 1537357262, // expiration time
    "iat": 1537357162, // issued at
    "userId": 80152 // custom field
}
```

### Signature

Encrypt by algorithm defined in header

```sh
HmacSHA256(base64(header)+"."+base64(payload), $secret)
```

## Flow

```mermaid
sequenceDiagram
client->>server: login
server->>server: authenticate
server->>server: get user id
opt Generate JWT
    server->>server: get JWT header
    note right of server: {"alg": "HS256", "typ": "JWT"}
    server->>server: base64 encode JWT header
    note right of server: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
    server->>server: put user id in JWT payload
    note right of server: {"iss": "Dino Lai", "sub": "dinos80152@gmail.com", "aud": "dinolai.com", "exp": 1537357262, "iat": 1537357162 ,"userId": 80152}
    server->>server: base64 encode JWT payload
    note right of server: eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0
    server->>server: generate signature
    note right of server:  lNSYE_dZuNPCjCf9ybMfIDiUJ4CXFZCqOn5zpJ5oqPY = HS256(base64(header)+"."+base64(payload), secret)
    server->>server: put it all together by [header].[payload].[signature]
    note right of server: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0.YaLyoBs8z5Va7YsIQaC6uEZDw8GZHBiV_2hIUSVQYUs
end
alt cookie
    server->>client: send cookie with JWT
    note over client, server: set-cookie: jwt=xxx, Http-only, max-age=...
    else header
    server->>client: response with header
    note over client, server: Authorization: Bearer <jwt>
end
client->>server: request with JWT
opt verify JWT
    server->>server: check signature to prevent data tamper
    note right of server: generate signature by header and payload, is the same as request signature?
    server->>server: check expiration
    note right of server: check exp field in payload
    server->>server: check owner
    note right of server: check aud field in payload
end
alt is OK
    server->>server: read user id from JWT
else is Fail
    server->>client: 401 UNAUTHORIZED
end
```

## Comparison

| | JWT | Cookie | Session |
|-|-----|--------|---------|
|Side|Client|Client|Server|
|Visible| ○ | ○ | ❌ |
|Tamper| ❌ | ○ | ❌ |
|Identify| ○ | ❌ | ○ |
|additional resource|spend computing for en/decode, encrypt| | diskIO or network IO |

## Reference

* [JSON Web Tokens - jwt.io](https://jwt.io/)
* [JSON Web Token - 在Web应用间安全地传递信息](http://blog.leapoahead.com/2015/09/06/understanding-jwt/)
* [八幅漫画理解使用JSON Web Token设计单点登录系统](http://blog.leapoahead.com/2015/09/07/user-authentication-with-jwt/)