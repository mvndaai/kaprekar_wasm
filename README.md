# kaprekar_wasm
Attempting to calculate Kaprekar's Constant using Go and Web Assembly

Kaprekar's Constant is 6174 and happens if you take a four digit number where all the values are not repeeating, arrange digits largest to smallest and subtract smallest to largest at most 7 times,

## Build

To copy the built in Go wasm file
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

To build the wasm file run:

```bash
GOOS=js GOARCH=wasm go build -o kaprekar.wasm
```

The file for wasm build by go can be copied to this directory
```bash
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```
Or I used jsdelivr.com to make the git repo file into a cdn


## References
https://golangbot.com/webassembly-using-go/
https://brilliant.org/wiki/kaprekars-constant/
https://codepo8.github.io/css-fork-on-github-ribbon/