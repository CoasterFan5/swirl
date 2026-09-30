# Swirl
Swirl is a full-stack web framework using file-based routing, Go for server-side code, and TypeScript for client-side code.

## Repository layout

- `cmd/swirl/` contains the development server executable.
- `packages/swirl/` contains the TypeScript framework implementation.
- `examples/playground/` is a consumer project used to exercise builds and editor/LSP support. Its generated `.swirl/` declarations intentionally live beside its `src/` directory.
- The root `go.mod` and `go.work` let the framework and playground remain separate Go modules while being developed together.

Run the development server from the repository root with:

```sh
go run ./cmd/swirl -project ./examples/playground
```

## Architecture
Swirl is a full stack web framework that uses file based routing and some black magic. It does not have support for node packages, and you write serverside logic in Golang, but client side logic in typescript. 

To create a page, it needs to be declared in a directory.
```bash
./index.swirl # This will be served on /
``` 
```bash
./about/index.swirl # This will be served on /about
```
See, convient! 
To add serverside code, you create a `server.go`, and thats about as far as I have gotten.

## Swirl Syntax
+page.swirl
```swirl
<script lang="ts">
  let counter = $state(0)
</script>
<effect ref={[counter]}>
  <div>
    The count is {counter.get()}
  </div>
</effect>
<style>
  div {
    color: #000001
  }
</style>
```
