// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var stopSignal = make(chan os.Signal, 1)

func init() {
	CLI.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Serve RoadRunner service(s)",
		RunE:  serveHandler,
	})

	signal.Notify(stopSignal, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
}

func serveHandler(cmd *cobra.Command, args []string) error {
	stopped := make(chan interface{})

	go func() {
		<-stopSignal
		Container.Stop()
		close(stopped)
	}()

	if err := Container.Serve(); err != nil {
		return err
	}

	<-stopped
	return nil
}
