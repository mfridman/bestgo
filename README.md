# bestgo

`bestgo` is a CLI that fetches aggregated star counts from the Best of Go API — a website that
regularly tracks GitHub stars, refreshing every hour for repositories with 50 stars or more ⭐️.

The Best of Go API is implemented with [Connect](https://connectrpc.com/), an RPC framework that
allows you to build APIs with Protobuf. The definitions are published to Buf Schema Registry:

https://buf.build/mf192/bestofgo/docs/main:api

## Usage

```bash
go install github.com/mfridman/bestgo@latest
```

```
Usage of bestgo:
  -i string
        grouping interval. Supported: year, quarter, month, day (default "year")
  -repo string
        full repository name. Example: pressly/goose (mandatory)
```

Example:

```bash
$ bestgo -repo golang/go -i year
2014 [4360] |■■■■■■■■
2015 [7275] |■■■■■■■■■■■■■■
2016 [8944] |■■■■■■■■■■■■■■■■■
2017 [11910]|■■■■■■■■■■■■■■■■■■■■■■
2018 [14851]|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2019 [15944]|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2020 [13464]|■■■■■■■■■■■■■■■■■■■■■■■■■
2021 [14115]|■■■■■■■■■■■■■■■■■■■■■■■■■■■
2022 [14596]|■■■■■■■■■■■■■■■■■■■■■■■■■■■
2023 [10973]|■■■■■■■■■■■■■■■■■■■■■

Repository: golang/go has 116432 ⭐️ stars total
```

### Bonus

Since the API is implemented with an RPC framework and the Protobuf definitions are published to
[Buf Schema Registry](https://buf.build), **anyone can pull an SDK client and just start using it.**

Huh?! Don't I need to pull the Protobuf files, install plugins and generate the stubs locally?

**Nope!**

The Buf Schema Registry implements a Go module proxy that exposes Generated SDKs.

This is really cool because you can now pull an SDK client for any supported language and start
using it without having to install any plugins or generate code locally. This is all done for you by
the Buf Schema Registry.

With just 2 commands:

```bash
go get buf.build/gen/go/mf192/bestofgo/protocolbuffers/go
go get buf.build/gen/go/mf192/bestofgo/connectrpc/go
```

We can now import the SDK client and start using it:

```go
import (
      "buf.build/gen/go/mf192/bestofgo/connectrpc/go/api/apiconnect"
      "buf.build/gen/go/mf192/bestofgo/protocolbuffers/go/api"
      "buf.build/gen/go/mf192/bestofgo/protocolbuffers/go/datapb"
)
```
