# brainfuck.go

### Description

[BrainFuck](https://en.wikipedia.org/wiki/Brainfuck) compiler written in [golang](https://github.com/golang/go).


##
### Table of Contents
- [Description](#Description)
- [Installation](#Installation)
- [Usage](#Usage)
- [License](#License)
- [Copyright](#Copyright)
##


### Installation

Build from source

    #! sh
    
    go get -u github.com/ronnathaniel/brainfuck.go
    cd ~/go/src/github.com/ronnathaniel/brainfuck.go
    go build .
    mv brainfuck.go /usr/local/bin/brainfuck

### Usage

BrainFuck source files *should* end with extensions `.b` or `.bf` 

Run `brainfuck`, pass in BF source files as args:

    brainfuck hello.bf
    
creates a `hello` compiled and optimized executable. Run the executable to compute your BF logic.

    ./hello 

### License

This project is provided by the MIT license.

### Copyright

samo c 20

![S4M0](https://www.juxtapoz.com/images/Austin%20McManus/April_2013/9/jux_samo.jpg)
