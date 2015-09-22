// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package config

import (
	"os"
	"path/filepath"

	"github.com/jcelliott/lumber"

	api "github.com/pagodabox/nanobox-api-client"
)

//
var (
	Console *lumber.ConsoleLogger
	Log     *lumber.FileLogger
)

// init
func init() {

	// check for a ~/.nanobox/nanobox.log file and create one if it's not found
	// NOTE: this is handled by the current logger (Lumber) however this may not
	// always be the case, so this is left in as a fallback
	logfile := filepath.Clean(Root + "/nanobox.log")
	// if fi, _ := os.Stat(logfile); fi == nil {
	//  fmt.Printf(stylish.Bullet("Creating %s directory", logfile))
	//  if _, err := os.Create(logfile); err != nil {
	//    return err
	//  }
	// }

	// set the default log level
	loglvl := lumber.INFO

	// check for debug mode and set the appropriate log level
	if os.Args[len(os.Args)-1] == "--debug" {
		loglvl = lumber.DEBUG

		//
		api.Debug = true
	}

	//
	Console = lumber.NewConsoleLogger(loglvl)

	// create a logger
	if Log, err = lumber.NewFileLogger(logfile, loglvl, lumber.ROTATE, 100, 1, 100); err != nil {
		panic(err)
	}
}