<p style="text-align: center"><img src="go-faker-logo.png" alt="Social card of FakerPHP"></p>

# goFaker

goFaker is a Go library that generates fake data for you. 

It's heavily inspired by [FakerPHP](https://github.com/FakerPHP/Faker)


## Getting Started
    
### Installation

```shell
go get https://github.com/harmlessprince/goFaker
```

## Basic Usage

Add this import line to the file you're working in:

```Go
import "github.com/harmlessprince/goFaker"
```

###  Crate Fake Data
Use goFaker.MakeFactory() to create and initialize a faker generator, which can generate data by calling methods named after the type of data you want.

```Go
// use the factory to create a Faker\Generator instance
generator :=  goFaker.MakeFactory()

// generate data by calling methods
fmt.Println(generator.Date("", "2020-08-10 15:04:05")) //1991-08-11

fmt.Println(generator.BloodType())                     // 0

fmt.Println(generator.PhoneNumber())                   //  729-988-596

fmt.Println(generator.Address()) // 67 Ville a 96812 ClantonVille

fmt.Println(generator.CountryCode()) // LU

```

### Localization

goFaker.MakeFactory() can take a locale as an argument, to return localized data. If no localized provider is found, the factory falls back to the default data.

```Go
locality := "en_NG"

generator := goFaker.MakeFactory(locality)

fmt.Println(generator.Date("", "2020-08-10 15:04:05")) // 1991-08-11

fmt.Println(generator.BloodType()) // 0

fmt.Println(generator.PhoneNumber()) // +234 905 352 8428

fmt.Println(generator.PostCode()) // 41 865

fmt.Println(generator.CountryCode()) // LU

fmt.Println(generator.Name()) // Titilayo Ebiere Olaoluwa

```

