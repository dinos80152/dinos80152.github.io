# JWT - JSON Web Token

## Purpose

Store data in client side, which could be read by can't be modified.

* Authentication
* Authorization

## Structure

`base64(header).base64(payload).signature`

> eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9**.**eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0**.**YaLyoBs8z5Va7YsIQaC6uEZDw8GZHBiV_2hIUSVQYUs

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
Client->>Auth Server: login
Auth Server->>Auth Server: authenticate
Auth Server->>Auth Server: get user id
opt Generate JWT
    Auth Server->>Auth Server: get JWT header
    note right of Auth Server: {"alg": "HS256", "typ": "JWT"}
    Auth Server->>Auth Server: base64 encode JWT header
    note right of Auth Server: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
    Auth Server->>Auth Server: put user id in JWT payload
    note right of Auth Server: {"iss": "Dino Lai", "sub": "dinos80152@gmail.com", "aud": "dinolai.com", "exp": 1537357262, "iat": 1537357162 ,"userId": 80152}
    Auth Server->>Auth Server: base64 encode JWT payload
    note right of Auth Server: eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0
    Auth Server->>Auth Server: generate signature: HS256(base64(header)+"."+base64(payload), secret)
    note right of Auth Server:  lNSYE_dZuNPCjCf9ybMfIDiUJ4CXFZCqOn5zpJ5oqPY
    Auth Server->>Auth Server: put it all together by [header].[payload].[signature]
    note right of Auth Server: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.<br/>eyJpc3MiOiJEaW5vIExhaSIsInN1YiI6ImRpbm9zODAxNTJAZ21haWwuY29tIiwiYXVkIjoiZGlub2xhaS5jb20iLCJleHAiOjE1MzczNTcyNjIsImlhdCI6MTUzNzM1NzE2MiwidXNlcklkIjo4MDE1Mn0.<br/>YaLyoBs8z5Va7YsIQaC6uEZDw8GZHBiV_2hIUSVQYUs
end
alt cookie
    Auth Server->>Client: send cookie with JWT
    note over Client, Auth Server: set-cookie: jwt=xxx, Http-only, max-age=...
else header
    Auth Server->>Client: response with header
    note over Client, Auth Server: Authorization: Bearer <jwt>
end
Client->>Application Server: request with JWT
opt verify JWT
    Application Server->>Application Server: check signature to prevent data tamper
    note left of Application Server: generate signature by header and payload,<br/> is the same as request signature?
    Application Server->>Application Server: check expiration
    note right of Application Server: check exp field in payload
    Application Server->>Application Server: check owner
    note right of Application Server: check aud field in payload
end
alt is Fail
    Application Server->>Client: 401 UNAUTHORIZED
else is OK
    Application Server->>Application Server: read user id from JWT
end
```

## Comparison

|     Comparison      |                  JWT                   | Cookie |       Session        |
| ------------------- | -------------------------------------- | ------ | -------------------- |
| Side                | Client                                 | Client | Server               |
| Visible             | ✓                                      | ✓      | ☓                    |
| Tamper              | ☓                                      | ✓      | ☓                    |
| Identify            | ✓                                      | ☓      | ✓                    |
| additional resource | spend computing for en/decode, encrypt |        | diskIO or network IO |

## Reference

* [JSON Web Tokens - jwt.io](https://jwt.io/)
* [JSON Web Token - 在Web应用间安全地传递信息](http://blog.leapoahead.com/2015/09/06/understanding-jwt/)
* [八幅漫画理解使用JSON Web Token设计单点登录系统](http://blog.leapoahead.com/2015/09/07/user-authentication-with-jwt/)