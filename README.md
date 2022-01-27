# ipapi-go

A go library for fetching locations from ip addresses using [ipapi.co](https://ipapi.co/api/?go#location-of-a-specific-ip).

## Usage

```bash
$ go get github.com/meinside/ipapi-go
```

```go
// ...

import "github.com/meinside/ipapi-go"

// ...

    result, err := ipapi.GetLocation("8.8.8.8")

    if err == nil {
        log.Printf("country = %s", result.Country)
    }

// ...
```

