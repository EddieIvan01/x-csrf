# x-csrf
middleware to defend CSRF attack for gin framework

***

## Usage

see [example](https://github.com/EddieIvan01/x-csrf/blob/master/example/example.go)

enable two Middleware in order：

1. SetCSRFToken
2. XCSRF

***

## Custom config

you could change these configs as needed

```go
// TokenLength is the length of csrf token
TokenLength = 16

// TokenKey is the key of csrf token
// it could be in get-query, post-form or header
TokenKey = "X-Csrf-Token"

// TokenCookie is the name of token cookie
TokenCookie = "X-Csrf-Token"

// DefaultExpire is the default expire time of cookie
DefaultExpire = 3600 * 6

// RandomSec is the flag which represents the random-source
// will be changed after each period of time
RandomSec = false

// randSource will be changed every DefaultExpire time
randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

// GenerateToken returns random CSRF token
GenerateToken = func() string {
    result := make([]byte, TokenLength)
    for i := 0; i < TokenLength; i++ {
        result[i] = padding[randSource.Intn(62)]
    }
    return string(result)
}
```

