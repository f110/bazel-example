package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"golang.org/x/xerrors"

	"github.com/f110/bazel-example/internal/server"
)

func gocon20(args []string) error {
	localFile := ""
	fs := flag.NewFlagSet("gocon20", flag.ContinueOnError)
	fs.StringVar(&localFile, "local-file", "", "Local file path")
	if err := fs.Parse(args); err != nil {
		return xerrors.Errorf(": %w", err)
	}

	s := server.New(localFile)

	notifyChan := make(chan os.Signal)
	go func() {
		select {
		case <-notifyChan:
			if err := s.Stop(); err != nil {
				fmt.Fprintf(os.Stderr, "%+v\n", err)
			}
		}
	}()
	signal.Notify(notifyChan, os.Interrupt)

	if err := s.Start(); err != nil {
		return xerrors.Errorf(": %w", err)
	}
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := gocon20(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
	}
}
