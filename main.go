// +build wasm
package main

import (
	_ "embed"
	"encoding/base64"
	"syscall/js"
)

//go:embed rsa/public.pem
var publicKeyBytes []byte

func Encode(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	pubKey, _ := RSA_PublicKeyFromBytes(publicKeyBytes) // 解密公匙
	encryData, _ := RSA_Encrypt([]byte(data), pubKey)   // 加密数据

	return base64.StdEncoding.EncodeToString(encryData)
}

func main() {
	T := js.Global().Get("T")
	if T.IsUndefined() {
		T = js.ValueOf(map[string]interface{}{
			"wasm": map[string]interface{}{},
		})
		js.Global().Set("T", T)
	}
	wasm := T.Get("wasm")
	wasm.Set("encode", js.FuncOf(Encode))
	// 需要阻塞，否则会抛出 Go program has already exited
	select {}
}
