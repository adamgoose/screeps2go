# screeps2go

screeps2go is an SDK Generator and Golang Runtime for screeps.com!

Well, it intends to be eventually. In the meantime, have a look. ;)

## Building

```bash
GOOS=js GOARCH=wasm go build -o example/example.wasm ./example
```

## Running

Include the following modules in your branch:

- `./example/main.js` as `main`
- `./example/example.wasm` as `wasmbin`
- `./screeps2go.js` as `screeps2go`

## Runtime

The "runtime" simply refers to how the SDK implements a game tick. Essentially,
your Go application will look something like this:

```go
package main

import . "github.com/adamgoose/screeps2go/runtime"

func main() {
	Run(func(g *Game) any {
	  // Do something with the game
	  return nil
	})
}
```

It's very sloppily generated in `gen.go` from the various types defined in the
aptly named `types.go`. In theory, the generator and type definitions can be
expanded to support the rest of the Screeps API.

## Attribution

I copy/pasted some constant definitions from an early Golang implementation by
[Robalian](https://github.com/Robalian/goScreeps/tree/master).
