package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/base64"
	"encoding/pem"
	"syscall/js"
)

//go:embed rsa/public.pem
var publicKeyBytes []byte

func main() {
	T := js.Global().Get("T")
	if T.IsUndefined() {
		T = js.ValueOf(map[string]interface{}{
			"wasm": map[string]interface{}{},
		})
		js.Global().Set("T", T)
	}
	wasm := T.Get("wasm")
	wasm.Set("encode", js.FuncOf(Encode)) // tinygo 不支持
	// 需要阻塞，否则会抛出 Go program has already exited
	select {}
}

// Encode
func Encode(this js.Value, args []js.Value) interface{} {
	data := args[0].String()
	b, _ := pem.Decode(publicKeyBytes)
	pubKey, _ := x509.ParsePKIXPublicKey(b.Bytes)
	encryData, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(data))
	return base64.StdEncoding.EncodeToString(encryData)
}
