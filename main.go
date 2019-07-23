// Copyright (c) 2019 Kevin Kragenbrink, II
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"github.com/kkragenbrink/slate/domain"
	"github.com/kkragenbrink/slate/infrastructure"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

// Constants containing exit status codes
const (
	ExitUsage = 64
	ExitError = 65
)

func main() {
	// initialize logrus
	infrastructure.InitLogger()
	logrus.Info("starting slate")

	// configure slate
	config, err := domain.NewSlateConfig()
	exitOnError(err, "unable to initialize configuration", ExitUsage)

	// start up
	slate, err := NewSlate(config)
	exitOnError(err, "error starting services", ExitError)

	slate.Start()
	defer slate.Stop()

	// shutting down
	logrus.Info("goodbye")
}

func exitOnError(err error, msg string, code int) {
	if err != nil {
		logrus.WithError(err).Error(msg)
		os.Exit(code)
	}
}
