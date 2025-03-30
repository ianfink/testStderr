/*
 * Copyright (C) 2025 Ian M. Fink.  All rights reserved.
 *
 * This program is free software:  you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU General Public License for
 * more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program.  If not, please see: https://www.gnu.org/licenses.
 *
 * Tabstop:	4
 */

package main

/*
 * Imports
 */

import (
	"fmt"
	"os"
	"time"
	"math/rand"
)

/**********************************************************************/

/**
 * Name:	RandomIntBetween
 *
 * @brief	Generate a random integer between min and max integers
 *
 * @param	min		a lower bound for the random integer
 * @param	max		an upper bound for the random integer
 *
 * @return	a random integer between the min and max integers
 */

func RandomIntBetween(min, max int) int {
	return min + rand.Intn(max-min)
} /* RandomIntBetween */

/**********************************************************************/

/**
 * Name:	produceOutput
 *
 * @brief	produce output to a file pointer at random times.
 *
 * @param	filePtr		a file pointer to write output
 * @param	outString	a string that credits the file pointer (e.g.,
 *						stderr, stdout, etc.)
 * @param	done		a channel that is used to provide information
 *						when the func is done producing output.
 *
 */

func produceOutput(filePtr *os.File, outString string, done chan bool) {
	var (
		i		int
		theRand	int
	)

	for i=0; i<5; i++ {
		theRand = RandomIntBetween(750, 2000)
		time.Sleep(time.Duration(theRand) * time.Millisecond)
		fmt.Fprintf(filePtr, "This was printd to %s and i = %d\n",
			outString, i)
	}

	done <- true
}

/**********************************************************************/

func main() {
	var (
		ch1			chan bool
		ch2			chan bool
		doneStderr	bool
		doneStdout	bool
	)

	// allocate the channels
	ch1 = make(chan bool)
	ch2 = make(chan bool)

	// initial the initial conditions for the loop
	doneStdout = false
	doneStderr = false

	// seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	go produceOutput(os.Stderr, "stderr", ch1)
	go produceOutput(os.Stdout, "stdout", ch2)

	for !doneStdout || !doneStderr {
		select {
			case doneStdout = <- ch2:
			case doneStderr = <- ch1:
		}
	}

} /* main */

/*
 * End of file:	main.go
 */

