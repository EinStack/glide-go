<div align="center">
    <img loading="lazy" src="https://github.com/EinStack/glide-python/blob/main/docs/glide_logo.png?raw=1" alt="Glide Logo" width="200px" height="200px" />
    <h1>Glide Go Client</h1>
    <p>🐿️ An official Go client for <a href="https://github.com/EinStack/glide">Glide, an open reliable fast model gateway</a>.</p>
    <a href="https://discord.gg/pt53Ej7rrc"><img src="https://img.shields.io/discord/1181281407813828710" alt="Discord" /></a>
    <a href="https://glide.einstack.ai/"><img src="https://img.shields.io/badge/build-view-violet%20?style=flat&logo=books&label=docs&link=https%3A%2F%2Fglide.einstack.ai%2F" alt="Glide Docs" /></a>
    <a href="https://artifacthub.io/packages/helm/einstack/glide"><img src="https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/einstack" alt="ArtifactHub" /></a>
	<br/>
	<a href="https://github.com/einstack/glide-go/actions/workflows/build.yaml">
		<img src="https://img.shields.io/github/actions/workflow/status/einstack/glide-go/build.yaml?branch=main&label=build&logo=github&style=flat-square" alt="Github Action" />
	</a>
	<a href="https://pkg.go.dev/github.com/einstack/glide-go"><img src="https://pkg.go.dev/badge/github.com/einstack/glide-go.svg" alt="Go Reference" /></a>
</div>

---

> Glide is under active development right now 🛠️

> Give us a star ⭐ to support the project and watch 👀 our repositories not to miss any update

## Features

...

## Installation

```cmd
go get github.com/einstack/glide-go
```

## Usage

For a full example take a look at [`hello.go`](examples/hello.go).

```go
package main

import (
	"context"
	"log"

	"github.com/einstack/glide-go"
	"github.com/einstack/glide-go/lang"
)

func main() {
	client, err := glide.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	req := lang.NewChatRequest("Hello")
	resp, err := client.Lang.Chat(ctx, "myrouter", req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("response: ", resp.Content())
}
```
