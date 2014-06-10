package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

const (
	help_text string = `
	Usage: cat [OPTIONS] [FILE]...
	concatenate and print thre content of files
		--help display this help and exit
		--version output version information and exit
	`
	version_text string = `
	cat (go-coreutils) 0.1

	Copyright (C) 2014, The GO-Coreutils Developers.
    This program comes with ABSOLUTELY NO WARRANTY; for details see
    LICENSE. This is free software, and you are welcome to redistribute 
    it under certain conditions in LICENSE.

	`
)

var (
	countNonBlank     = flag.Bool("b", false, "Number the non-blank output lines , starting at 1.")
	numberOutput      = flag.Bool("n", false, "Number the output lines,  starting at 1.")
	squeezeEmpthLines = flag.Bool("s", false, "Squeeze multiple adjacent empty lines, causing the output top be single spaced.")
)

func openFiles(s string) (io.ReadWriteCloser, error) {
	fi, err := os.Stat(s)
	if err != nil {
		return nil, err
	}

	if fi.Mode()&os.ModeSocket != 0 {
		return net.Dial("unix", s)
	}
	return os.Open(s)
}

func dumpLines(w io.Writer, r io.Reader) (n int64, err error) {
	var lastline, line string
	br := bufio.NewReader(r)
	var nr int64
	nr = 0
	for {
		line, err = br.ReadString('\n')
		if err != nil {
			return
		}
		if *squeezeEmpthLines && lastline == "\n" && line == "\n" {
			continue
		}

		if *countNonBlank && line == "\n" || line == "" {
			fmt.Fprint(w, line)

		} else if *countNonBlank || *numberOutput {
			nr++
			fmt.Fprintf(w, "%6d\t%s", nr, line)
		} else {
			fmt.Fprint(w, line)
		}
		lastline = line
	}
	return

}

func main() {
	help := flag.Bool("help", false, help_text)
	version := flag.Bool("version", false, version_text)
	flag.Parse()

	if *help {
		fmt.Println(help_text)
		os.Exit(0)
	}

	if *version {
		fmt.Println(help_text)
		os.Exit(0)
	}

	rcopy := io.Copy
	if *countNonBlank || *numberOutput || *squeezeEmpthLines {
		fmt.Println("countNonBlank=", *countNonBlank)
		fmt.Println("numberOutput=", *numberOutput)
		fmt.Println("squeezeEmpthLines", *squeezeEmpthLines)
		rcopy = dumpLines
	}

	for index, fname := range flag.Args() {
		fmt.Println("index", index, "fname=", fname)
		if fname == "-" {
			rcopy(os.Stdout, os.Stdin)
		} else {
			f, err := openFiles(fname)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			rcopy(os.Stdout, f)
			f.Close()
		}

	}

}
