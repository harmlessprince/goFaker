<p style="text-align: center"><img src="go-faker-logo.png" alt="Social card of FakerPHP"></p>

# goFaker

goFaker is a Go library that generates fake data for you. 

It's heavily inspired by [FakerPHP](https://github.com/FakerPHP/Faker)


## Getting Started
    
### Installation

```shell
go get https://github.com/harmlessprince/goFaker
```

## Quick Start

Add this import line to the file you're working in:

```Go
import "github.com/harmlessprince/goFaker"
```

Use goFaker.MakeFactory(locality) to create and initialize a faker generator, which can generate data by calling methods named after the type of data you want.

```Go
// use the factory to create a Faker\Generator instance
generator :=  goFaker.MakeFactory(locality)
// generate data by calling methods
fmt.Println(generator.GenerateDate("", "2020-08-10 15:04:05"))

```