# Swirl 
Swirl is truly a framework for the history books, in a positive or negative connotation I will not say. 

## Architecture
Swirl is a full stack web framework that uses file based routing and some black magic. It does not have support for node packages, and you write serverside logic in Golang, but client side logic in typescript. 
To create a page, it needs to be declared in a directory.
```bash
./+page.swirl # This will be served on /
``` 
```bash
./about/page.swirl # This will be served on /about
```
See, convient! 
To add serverside code, you create a `+server.go`, and thats about as far as I have gotten.

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
