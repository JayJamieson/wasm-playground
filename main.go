package main

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

func main() {
	manifest := extism.Manifest{
		AllowedPaths: map[string]string{

			// Here we specifify a host directory data to be linked
			// to the /mnt directory inside the wasm runtime
			"data": "/mnt",
		},
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: "plugin.wasm",
			},
		},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	data := []byte("Hello world, this is written from within our wasm plugin.")
	exit, _, err := plugin.Call("write_file", data)
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}
}
