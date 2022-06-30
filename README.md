## Example Usage

### Build wasm

```
tinygo build -o bin/fib.wasm -target wasi cmd/wasm/main.go
```

### Use wasm

```go
package main

import (
	"fmt"
	"io/ioutil"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
	"go.uber.org/zap"
)

var (
	log, _ = zap.NewDevelopment()
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("bin/fib.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	wasiEnv, err := wasmer.NewWasiStateBuilder("fib").Finalize()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Instantiates the module
	importObject, err := wasiEnv.GenerateImportObject(store, module)
	if err != nil {
		log.Fatal(err.Error())
	}

	instance, err := wasmer.NewInstance(module, importObject)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Gets the `Fibonacci` exported function from the WebAssembly instance.
	fibFunc, err := instance.Exports.GetFunction("Fibonacci")
	if err != nil {
		log.Fatal(err.Error())
	}

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	result, err := fibFunc(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(result) // 89
}
```
