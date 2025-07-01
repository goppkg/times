times (fork from [times](https://github.com/jing332/times))
==========

[![GoDoc](https://godoc.org/github.com/djherbis/times?status.svg)](https://godoc.org/github.com/djherbis/times)
[![Release](https://img.shields.io/github/release/djherbis/times.svg)](https://github.com/djherbis/times/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](LICENSE.txt)
[![go test](https://github.com/djherbis/times/actions/workflows/go-test.yml/badge.svg)](https://github.com/djherbis/times/actions/workflows/go-test.yml)
[![Coverage Status](https://coveralls.io/repos/djherbis/times/badge.svg?branch=master)](https://coveralls.io/r/djherbis/times?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/djherbis/times)](https://goreportcard.com/report/github.com/djherbis/times)
[![Sourcegraph](https://sourcegraph.com/github.com/djherbis/times/-/badge.svg)](https://sourcegraph.com/github.com/djherbis/times?badge)

Usage
------------

File Times for #golang

Go has a hidden time functions for most platforms, this repo makes them accessible.

```go
package main

import (
 "fmt"
 "log"
 "os"
 "path/filepath"
 "time"

 "github.com/goppgk/times"
)

func main() {
 switch len(os.Args) {
 case 1:
  tempFile()
  fmt.Println()
  tempDir()

 default:
  printTimes(os.Args[1])
 }
}

func tempDir() {
 name, err := os.MkdirTemp("", "")
 if err != nil {
  log.Fatal(err)
 }
 defer os.Remove(name)
 fmt.Println("# DIR: " + name)

 symname := filepath.Join(filepath.Dir(name), "sym-"+filepath.Base(name))
 if err := os.Symlink(name, symname); err != nil {
  log.Fatal(err)
 }
 defer os.Remove(symname)

 newAtime := time.Now().Add(-10 * time.Second)
 newMtime := time.Now().Add(10 * time.Second)
 if err := os.Chtimes(name, newAtime, newMtime); err != nil {
  log.Fatal(err)
 }

 printTimes(symname)
}

func tempFile() {
 f, err := os.CreateTemp("", "")
 if err != nil {
  log.Fatal(err)
 }
 defer os.Remove(f.Name())
 defer f.Close()
 fmt.Println("# FILE: " + f.Name())

 symname := filepath.Join(filepath.Dir(f.Name()), "sym-"+filepath.Base(f.Name()))
 if err := os.Symlink(f.Name(), symname); err != nil {
  log.Fatal(err)
 }
 defer os.Remove(symname)

 newAtime := time.Now().Add(-10 * time.Second)
 newMtime := time.Now().Add(10 * time.Second)
 if err := os.Chtimes(f.Name(), newAtime, newMtime); err != nil {
  log.Fatal(err)
 }

 printTimes(symname)
}

func printTimes(name string) {
 fmt.Println("## Stat:", name)
 printTimespec(times.Stat(name))

 fmt.Println("\n## Lstat:", name)
 printTimespec(times.Lstat(name))
}

func printTimespec(ts times.Timespec, err error) {
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("AccessTime:", ts.AccessTime())
 fmt.Println("ModTime:", ts.ModTime())

 if ts.HasChangeTime() {
  fmt.Println("ChangeTime:", ts.ChangeTime())
 }

 if ts.HasBirthTime() {
  fmt.Println("BirthTime:", ts.BirthTime())
 }
}
```

Supported Times
------------

|  | windows | linux | solaris | dragonfly | nacl | freebsd | darwin | netbsd | openbsd | plan9 | js | aix |
|:-----:|:-------:|:-----:|:-------:|:---------:|:------:|:-------:|:----:|:------:|:-------:|:-----:|:-----:|:-----:|
| atime | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| mtime | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| ctime | ✓* | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |  | ✓ | ✓ |
| btime | ✓ | ✓* |  |  |  | ✓ |  ✓| ✓ |  |  |

* Linux btime requires kernel 4.11 and filesystem support, so HasBirthTime = false.
Use Timespec.HasBirthTime() to check if file has birth time.
Get(FileInfo) never returns btime.
* Windows XP does not have ChangeTime so HasChangeTime = false,
however Vista onward does have ChangeTime so Timespec.HasChangeTime() will
only return false on those platforms when the syscall used to obtain them fails.
* Also note, Get(FileInfo) will now only return values available in FileInfo.Sys(), this means Stat() is required to get ChangeTime on Windows

Installation
------------

```sh
go get -u github.com/goppkg/times
```
