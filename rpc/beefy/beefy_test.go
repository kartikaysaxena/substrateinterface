package beefy

import (
	"os"
	"testing"

	"github.com/kartikaysaxena/susbtrateinterface/client"
	"github.com/kartikaysaxena/susbtrateinterface/config"
)

var testBeefy Beefy

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	testBeefy = NewBeefy(cl)
	os.Exit(m.Run())
}
