# bestgo

`bestgo` is a CLI that pulls aggregated star counts from the best of go API, a website that polls GitHub stars every 1 hour for repositories that have **>=50 stars ⭐️**.

The best of go API is implemented with [connect-go] and the Protobuf definitions are published:

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
        full repository name. Example: go-chi/chi (mandatory)
```

Example:

```bash
$ bestgo -repo golang/go -i year
2014 [4411] |■■■■■■■■
2015 [7366] |■■■■■■■■■■■■■■
2016 [9055] |■■■■■■■■■■■■■■■■■
2017 [12048]|■■■■■■■■■■■■■■■■■■■■■■
2018 [15030]|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2019 [16177]|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2020 [13713]|■■■■■■■■■■■■■■■■■■■■■■■■■
2021 [14529]|■■■■■■■■■■■■■■■■■■■■■■■■■■■
2022 [11847]|■■■■■■■■■■■■■■■■■■■■■■

Repository: golang/go has 104176 ⭐️ stars total
```

### Bonus

Since the API is implemented with an RPC framework and the Protobuf definitions are published to [Buf Schema Registry][buf-build], anyone can pull an SDK client and just start using it.

Huh?! Don't I need to pull the Protobuf files, install plugins and generate the stubs locally?

**Nope!**

The Buf Schema Registry implements a Go module proxy that exposes Generated SDKs. This is super cool because you can continue using `go get ...` and the source code is generated for you.

With just 2 commands:

```bash
go get buf.build/gen/go/mf192/bestofgo/library/go
go get buf.build/gen/go/mf192/bestofgo/library/connect-go
```

We get an Go SDK client that we can import and start using:

```go
import (
      "buf.build/gen/go/mf192/bestofgo/library/connect-go/api/apiconnect"
      "buf.build/gen/go/mf192/bestofgo/library/go/api"
      "buf.build/gen/go/mf192/bestofgo/library/go/datapb"
)
```

[buf-build]: https://buf.build
[connect-go]: https://github.com/bufbuild/connect-go
