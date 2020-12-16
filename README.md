# brainfuck.go

## Description

[BrainFuck](https://en.wikipedia.org/wiki/Brainfuck) interpreter written in [golang](https://github.com/golang/go).


##
### Table of Contents
- [Description](#Description)
- [Installation](#Installation)
- [Usage](#Usage)
- [License](#License)
- [Copyright](#Copyright)
##




## Installation

Pull this module

    go get -u github.com/ronnathaniel/brainfuck.go

Build executable from source

    cd ~/go/src/github.com/ronnathaniel/brainfuck.go
    go build .
    mv brainfuck.go /usr/local/bin/brainfuck

## Usage

This compiler can be used both in-line code and in the command line. Simple integrations allow for Used-From-Anywhere.


### CLI

On the command line side, this interpreter can be stored as a quasi-executable, like `gcc`.


Run `brainfuck`, pass in BF source files as args. Source files *should* end with extensions `.b` or `.bf` 

    brainfuck hello.bf
    
interprets and optimizes `hello.bf` before computing the logic. 
Input is retrieved from stdin and Output pushed to stdout.

Optionally pass in a string from the command line.

    brainfuck -c "++[>+<-]"
    

### Inline 

Compiling within code is easy. The `Exec` function will parse, optimize, compile, and run any string of BrainFuck logic.

The function has a signature of

    func Exec(inputRaw string) 
    
allowing it to be called from anywhere.

    import "github.com/ronnathaniel/brainfuck.go/brainfuck"
    
    brainfuck.Exec(">++++[+>+++>>+<<<-]")
    
Open source files with a supplied `OpenFile` function. Pass it a relative string path.

    fuckRaw := brainfuck.OpenFile("<PATH/TO/SRC>.bf")
    brainfuck.Exec(fuckRaw)

## Configurations

Enable debugging logs by setting the `DEBUG` flag.

    brainfuck.DEBUG = true
    
The default tap size is set to 500. Override it with the `TAPE_SIZE_DEFAULT` flag.

    brainfuck.TAPE_SIZE_DEFAULT = <NEW-SIZE: uint32>
    
### Structure

    .
    ├── LICENSE
    ├── README.md
    ├── brainfuck/   - importable package
    ├── cmd.go       - handles command line
    └── examples/    - example bf src


### Contributions Welcomed :)

Report Issues, create Pull Requests, or email thesamogroup@gmail.com 
### License

This project is provided by the MIT license. All rights reserved.

### Copyright

samo c 20

![S4M0](https://www.juxtapoz.com/images/Austin%20McManus/April_2013/9/jux_samo.jpg)
