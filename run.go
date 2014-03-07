package main

/*
The MIT License (MIT)

Copyright (c) 2014 Lealen bez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"log"
	"strings"
)

func main() {
	if len(os.Args)<2 {
		fmt.Printf("Usage: %s ./program\n", os.Args[0])
		return
	}
	for {
		log.Printf("!R Start program with command: %s\n", strings.Join(os.Args[1:], " "))
		var cmd *exec.Cmd
		if len(os.Args)>2 {
			cmd = exec.Command(os.Args[1], strings.Join(os.Args[2:], " "))
		} else {
			cmd = exec.Command(os.Args[1])
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println(err)
		}
		err = cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}
		go io.Copy(os.Stdout, stdout) 
		go io.Copy(os.Stderr, stderr) 
		cmd.Wait()
		log.Printf("!R Program with command: %s has stopped working... \n", os.Args[1])
	}
}
