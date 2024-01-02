package main

import (
	"github.com/samslhsieh/dike"
)

func main() {
	//d := dike.New(&dike.Options{
	//d := dike.New(&dike.Options{
	//	Out:     os.Stderr,
	//	IsDebug: true,
	//	Format:  dike.Pretty,
	//})

	d := dike.New(nil)

	d.Logger.Info("test_1", "key_1", "value_1", "key_2", "value_2")

	d.Logger.Info("test_2")
}
