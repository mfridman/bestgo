# bestgo

`bestgo` is a CLI that pulls live data from https://api.bestofgo.dev. This is an application that scrapes GitHub data for Go repositories that have **50 or more stars ⭐️** once an hour.

The UI is now live at https://bestofgo.dev.

The API is Protobuf-based and implemented with Twirp. You can view the Protobuf docs for the APIServer here:

https://buf.build/mf192/bestofgo/docs/main/api#api.APIService

## Usage 

```bash
go install github.com/mfridman/bestgo@latest
```

```
Usage of bestgo:
  -i string
        grouping interval. Supported: year, quarter, month (default "year")
  -repo string
        full repository name. Example: go-chi/chi (mandatory)
```

Example:

```bash
$ bestgo -repo golang/go -i year
2014 [4454]	|■■■■■■■■■■■■■
2015 [7440]	|■■■■■■■■■■■■■■■■■■■■■■
2016 [9158]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2017 [12205]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2018 [15276]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2019 [16538]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2020 [14117]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
2021 [10785]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■

Repository: golang/go has 89973 ⭐️ stars total
```

### Bonus

Since the API is Protobuf-based, anyone can pull the SDK Client and just use it. Huh? Don't I need to pull the `proto` files, install a bunch of plugins and then locally generate my source code?

**Nope!**

Heh, check out the source code, the interesting bit is *where* the SDK is being fetched from:

```go
import "go.buf.build/demolab/twirp-go/mf192/bestofgo/api"
```

That's right, there are hosted Protobuf files on [buf.build](https://buf.build) as well as `protoc`-based templates (in this case Go + Twirp). Which means code generation is taking place remotely and you don't have to do anything. 

Just `go get` the code, run `go mod tidy` and you're ready to use it, just like this CLI!

If this doesn't excite you, I don't know what will. Maybe this [tweet](https://twitter.com/_mfridman/status/1426677430320783364)

The folks over at https://buf.build are working on some neat stuff, check it out!
