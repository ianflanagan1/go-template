# go-template

The starting point for a Go project. Don't copy it by hand: run `devbox run new` in [go-baseline](../go-baseline), which copies it, renames it and checks it.

It holds one example of each layer, to replace with the project's own code:

| Path | Layer |
|---|---|
| `internal/core/greeting` | Functional core: decides the greeting from a name and an hour |
| `internal/app` | Shell: reads the time through the `Clock` port, calls the core, writes the result |
| `internal/adapters/clock` | Adapter: the system clock |
| `cmd/app` | Composition root: wires the adapter into the shell |

The settings come from go-baseline, which `go.mod` pins as a tool; `devbox run check` fails if the project drifts from them.
