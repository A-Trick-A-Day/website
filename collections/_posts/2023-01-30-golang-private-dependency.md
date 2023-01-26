---
title: "Private Dependency"
date: 2023-01-30T07:49:03+01:00
layout: post
authors: ["Sidorenko Konstantin"]
categories: ["Golang", "Development"]
description: How to have private modules.
thumbnail: "assets/images/thumbnail/2023-01-30-golang-private-dependency.png"
comments: true
---

## Introduction to Go Dependencies

These are based on a **module** whose source is in a **git repository**.

The dependency will target to a **commit** or a **tag**.

Example:

```txt
golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect => this target a commit
gopkg.in/yaml.v2 v2.4.0 // indirect => this target a tag
```

When you **add a dependency** to your project you will do this:

```bash
go get -u github.com/gofiber/fiber/v2 # this will use the latest version without @
```

You will get a `go.mod` file like:

```txt
module test

go 1.19

require github.com/gofiber/fiber/v2 v2.41.0

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.44.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
)
```

{% include alerts/info.html content='For more information on go and its dependencies management, check <a href="https://go.dev/doc/modules/managing-dependencies">this page</a>' %}

{% include alerts/warning.html content='If you need to expose your module (or try to always), name it according to its git repository URL.' %}

## How to use private dependency

What if my **dependencies** are in a **private** git repository ?

Let's use **gitlab as a private repository** which is the most common in companies.

Here is the project <https://gitlab.com/a-trick-a-day/golang-private-dependency>, you can't see it because it's a private project.

![gitlab repository](/assets/images/posts/2023-01-30-golang-private-dependency_gitlab-repository.jpg)

It's a simple `main.go`:

```go
package main

import "fmt"

func main() {
	Hello()
}

func Hello() {
	fmt.Println("Hello")
}
```

and its `go.mod`:

```txt
module gitlab.com/a-trick-a-day/golang-private-dependency

go 1.19
```

In my case I will use a [gitlab deploy token](https://docs.gitlab.com/ee/user/project/deploy_tokens/) with the read repository permission.
Depending of your project you can **use** a **personal** or a **group token** with **read repository permission**.

Run this commands:

```bash
git config --global url.https://gitlab+deploy-token-1715684:zaA_MonZ6XskfCVM7ACm@gitlab.com.insteadOf https://gitlab.com # will rewrite the url during remote actions
go env -w GOPRIVATE=gitlab.com # will tell to go it's a private repo
go install gitlab.com/a-trick-a-day/golang-private-dependency@latest # will install the latest version of golang-private-dependency
```

For [GOPRIVATE](https://goproxy.io/docs/GOPRIVATE-env.html), the go command defaults to downloading modules from the public Go module mirror at goproxy.io.
The GOPRIVATE environment variable controls which modules the go command considers to be private (not available publicly) and should therefore not use the proxy or checksum database.

You should be able to run this command now:

```bash
~/go/bin/golang-private-dependency
Hello
```

Now to use this as dependency you can do this:

```bash
➜  go mod init test # to init my module not needed if you already have it
go: creating new go.mod: module test
➜  go get gitlab.com/a-trick-a-day/golang-private-dependency
go: added gitlab.com/a-trick-a-day/golang-private-dependency v0.0.0-20230123145256-d7f205052328
```

{% include alerts/info.html content="Don't forget to do a<b>go mod tidy</b> after updating your dependencies." %}

`go.mod`:

```txt
module test

go 1.18

require gitlab.com/a-trick-a-day/golang-private-dependency v0.0.0-20230123145256-d7f205052328 // indirect
```

{% include alerts/info.html content='Inside your gitlab CI you can use the <b>CI_JOB_TOKEN</b><br/><b>git config --global url.https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com.insteadOf https://gitlab.com</b>' %}

During development, you should be able to:

```go
package main

import (
	"fmt"

	d "gitlab.com/a-trick-a-day/golang-private-dependency"
)

func main() {
	fmt.Println("aa")
	d.Hello()
}
```

In our case it doesn't work because we try to import a main package.

{% include alerts/info.html content='Try to avoid <b>-</b> in the names of go libraries, it should have a <a href="https://go.dev/ref/spec#identifier">valid identifiers</a>.' %}
