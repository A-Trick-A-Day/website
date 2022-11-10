---
title: "HTTP RoundTripper"
date: 2022-11-14T07:49:03+01:00
layout: post
authors: ["Sidorenko Konstantin"]
categories: ["Golang", "Development"]
description: The basics of Go HTTP Transport & RoundTripper.
thumbnail: "assets/images/thumbnail/2022-11-14-golang-http-roundtripper.png"
comments: true
---

In this article, I will discuss what is **Round tripping** in Go, its use cases and different application examples.

From the [Go doc](https://pkg.go.dev/net/http#RoundTripper):

```txt
RoundTripper is an interface representing the ability to execute a single HTTP transaction,
obtaining the Response for a given Request.
```

Basically, it means that you can **get into** what happens between an **HTTP request** and the receipt of a **response**. In simple terms, it's **like "middleware"** but for an **http client**.

Since [http.RoundTripper](https://pkg.go.dev/net/http#RoundTripper) is an interface.
All you have to do to get this functionality is to **implement RoundTrip**:

```go
type MyFirstRoundTripper struct {}

func (m *MyFirstRoundTripper) RoundTrip(r *http.Request)(*Response, error) {
	// wouhou I did my first round tripper
}
```

And **use it as transport** at the **http.Client** level.

Usecases:

- Caching http responses
- Adding appropriate (authorization) headers to the request
- Rate limiting
- Logging
- Metrics
- Monitoring
- Everything

## My first RoundTripper

In this part we will make a simple [RoundTripper](https://pkg.go.dev/net/http#RoundTripper) that **allows us to add a header**.

I took this example because it shows how to **pass variables** to its [RoundTripper](https://pkg.go.dev/net/http#RoundTripper):

```go
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

type BasicAuthTransport struct {
	Password string
	Username string
}

func (b BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprint("Basic ",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(
			b.Username, ":", b.Password)))))
	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	client := http.Client{
		Transport: BasicAuthTransport{
			Username: "a-trick",
			Password: "a-day",
		},
	}
	_, _ = client.Get("http://a-trick-a-day.github.io/")
}
```

Thanks to this **all the requests** made with the **client** will **have this header**.

All the "configuration" of our requests will be done at the **client level** so you **don't have to share** all **this configuration** and **only share the client**.

## Nested Rountripper

In this part, we will see how to **nest several [RoundTripper](https://pkg.go.dev/net/http#RoundTripper)**:

```go
type LogTransport struct {
	Level               string
	DefaultRoundTripper http.RoundTripper
}

func (l LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println("level:", l.Level, " - request: ", req.URL)
	return l.DefaultRoundTripper.RoundTrip(req)
}

func main() {
	client := http.Client{
		// to avoid to be spam by the redirection for the demo
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: LogTransport{
			Level: "INFO",
			DefaultRoundTripper: BasicAuthTransport{
				Username: "a-trick",
				Password: "a-day",
			},
		},
	}
	_, _ = client.Get("http://a-trick-a-day.github.io/")
}
```

Ouput:

```txt
2022/11/08 14:27:26 level: {INFO {a-trick a-day}}  - request:  http://a-trick-a-day.github.io/
```

As you can see, you need to have a **field in your RoundTripper structure** that **contains** the **RoundTripper to inherit**.

{% include alerts/warning.html content='Pay attention to the execution order!' %}

## Response RoundTripper

Here we will see how to use [RoundTripper](https://pkg.go.dev/net/http#RoundTripper) to **manipulate the response**:

```go
type PostRequestTransport struct {
	DefaultRoundTripper http.RoundTripper
}

func (p PostRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// execute the round trip to get the response
	resp, err :=  p.DefaultRoundTripper.RoundTrip(req)
	log.Println("resp:", resp.Header)
	return resp, err
}

func main() {
	client := http.Client{
		// to avoid to be spam by the redirection for the demo
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport:  PostRequestTransport{
			DefaultRoundTripper: LogTransport{
				Level: "INFO",
				DefaultRoundTripper: BasicAuthTransport{
					Username: "a-trick",
					Password: "a-day",
				},
			},
		},
	}
	_, _ = client.Get("http://a-trick-a-day.github.io/")
}
```

```txt
2022/11/08 14:27:26 level: {INFO {a-trick a-day}}  - request:  http://a-trick-a-day.github.io/
2022/11/08 14:27:26 resp: map[Accept-Ranges:[bytes] Age:[148] Connection:[keep-alive] Content-Length:[162] Content-Type:[text/html] Date:[Tue, 08 Nov 2022 14:27:26 GMT] Location:[https://a-trick-a-day.github.io/] Permissions-Policy:[interest-cohort=()] Server:[GitHub.com] Vary:[Accept-Encoding] Via:[1.1 varnish] X-Cache:[HIT] X-Cache-Hits:[1] X-Fastly-Request-Id:[ee289dfdb0c0400e8f0a1fafa8a6712c7fbae831] X-Github-Request-Id:[6148:1350C:A0FFBE:A5A5B6:636A66BA] X-Served-By:[cache-cdg20744-CDG] X-Timer:[S1667917647.738975,VS0,VE1]]
```

In **this use case** you can **implement** for example **cache the response**, **retrieve telemetry headers** etc.

## Resume

Here is the final version with the **different ways** to use the [RoundTripper](https://pkg.go.dev/net/http#RoundTripper), if you have **others implementations do not hesitate to share** them with me so I can update this post:

```go
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

type BasicAuthTransport struct {
	Password string
	Username string
}

func (b BasicAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", fmt.Sprint("Basic ",
		base64.StdEncoding.EncodeToString([]byte(fmt.Sprint(
			b.Username, ":", b.Password)))))
	return http.DefaultTransport.RoundTrip(req)
}

type LogTransport struct {
	Level               string
	DefaultRoundTripper http.RoundTripper
}

func (l LogTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println("level:", l, " - request: ", req.URL)
	return l.DefaultRoundTripper.RoundTrip(req)
}

type PostRequestTransport struct {
	DefaultRoundTripper http.RoundTripper
}

func (p PostRequestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := p.DefaultRoundTripper.RoundTrip(req)
	log.Println("resp:", resp.Header)
	return resp, err
}

func main() {
	client := http.Client{
		// to avoid to be spam by the redirection for the demo
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: PostRequestTransport{
			DefaultRoundTripper: LogTransport{
				Level: "INFO",
				DefaultRoundTripper: BasicAuthTransport{
					Username: "a-trick",
					Password: "a-day",
				},
			},
		},
	}
	_, _ = client.Get("http://a-trick-a-day.github.io/")
}
```

Go playground: <https://go.dev/play/p/s6Kt1FgylE->. (You can't test directly from the go playground because HTTP requests are blocked)

Now **you know the basics** of [RoundTripper](https://pkg.go.dev/net/http#RoundTripper)!

{% include alerts/info.html content='If you need a <strong>simple rest client</strong> don’t hesitate to look at <a href="https://github.com/go-resty/resty">resty</a>.' %}

## References

<https://lanre.wtf/blog/2017/07/24/roundtripper-go/>
