package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/urfave/cli/v2"
)

var (
	counter     int
	counterLock sync.Mutex
)

var done chan struct{}

func main() {

	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "prefix",
			Value:   "",
			Usage:   "Prefix of the address",
			EnvVars: []string{"LEGACY_COMPAT_LANG"},
		},
		&cli.StringFlag{
			Name:    "suffix",
			Value:   "",
			Usage:   "Suffix of the address",
			EnvVars: []string{"LEGACY_COMPAT_LANG"},
		},
	}

	app := &cli.App{
		Flags: flags,
		Name:  "EVM Address Generator",
		Usage: "Generate EVM addresses with a given prefix and suffix",
		Action: func(cCtx *cli.Context) error {
			prefix := cCtx.String("prefix")
			suffix := cCtx.String("suffix")

			multiThreading(prefix, suffix)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}

// Multithreaded address generator
func multiThreading(prefix string, suffix string) {
	done = make(chan struct{})
	start := time.Now() // Record the start time

	thread := uint32(runtime.GOMAXPROCS(0))

	for i := uint32(0); i < thread; i++ {
		go mine(prefix, suffix)
	}
	mine(prefix, suffix)

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("Total execution time: %s\n", elapsed)
}

func mine(prefix string, suffix string) {
	for i := 0; true; i++ {
		mne, err := hdwallet.NewMnemonic(128)
		if err != nil {
			log.Fatal(err)
		}

		// create the wallet from mnemonic
		wallet, err := hdwallet.NewFromMnemonic(mne)
		if err != nil {
			log.Fatal(err)
		}

		path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
		acc, err := wallet.Derive(path, false)
		if err != nil {
			log.Fatal(err)
		}

		privateKey, err := wallet.PrivateKeyHex(acc)
		if err != nil {
			log.Fatal(err)
		}

		incrementCounter()

		if strings.HasPrefix(acc.Address.Hex(), prefix) && strings.HasSuffix(acc.Address.Hex(), suffix) {
			log.Println("\n Your account's mnemonic: ", mne,
				"\n Private key: ", privateKey,
				"\n Address: ", acc.Address.Hex())

			close(done) // Signal all Goroutines to stop gracefully
			return
		}

		select {
		case <-done:
			return // Stop the Goroutine if signaled
		default:
		}

		fmt.Println("Processed accounts: ", getCounter())
	}
}

// Helper functions for thread-safe counter operations
func incrementCounter() {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
}

func getCounter() int {
	counterLock.Lock()
	defer counterLock.Unlock()
	return counter
}
