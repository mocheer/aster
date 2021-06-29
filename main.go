// +build wasm
package main

import (
	_ "embed"
	"syscall/js"

	"github.com/mocheer/pluto/ec"
	"github.com/mocheer/pluto/jsg"
)

//go:embed rsa/public.pem
var publicKeyBytes []byte

func Encode(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	pubKey, _ := ec.RSA_PublicKeyFromBytes(publicKeyBytes) // 解密公匙
	encryData, _ := ec.RSA_Encrypt([]byte(data), pubKey)   // 加密数据
	return jsg.BtoaBytes(encryData)
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
	select {}
}
