// + build wasm
package main

import (
	_ "embed"
	"encoding/hex"
	"syscall/js"

	"github.com/mocheer/aster/pkg/ec"
)

// // tinygo不支持 go:embed
// //go:embed xx

var Key []byte = []byte("ZIAIFOBQHQCWYYZI")

func main() {
	T := js.Global().Get("T")
	if T.IsUndefined() {
		T = js.ValueOf(map[string]interface{}{
			"wasm": map[string]interface{}{},
		})
		js.Global().Set("T", T)
	}
	wasm := T.Get("wasm")
	wasm.Set("encry", js.FuncOf(Encry))
	wasm.Set("decry", js.FuncOf(Decry))
	wasm.Set("version", js.ValueOf("1.0.0"))
	// 需要阻塞，否则会抛出 Go program has already exited
	select {}
}

// Encry
func Encry(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	encryptCode := ec.AesEncryptCBC([]byte(data), Key)
	return hex.EncodeToString(encryptCode)
}

// Decry
func Decry(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	rawData, _ := hex.DecodeString(data)
	decryptCode := ec.AesDecryptCBC(rawData, Key)
	return string(decryptCode)
}
