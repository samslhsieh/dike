package main

import (
	"os"

	"github.com/samslhsieh/dike"
)

func main() {
	d := dike.New(&dike.Options{
		Out:     os.Stderr,
		IsDebug: true,
		Format:  dike.Pretty,
	})

	d.Logger.Info("test", "key", "value")
}
