# Ethereum Address Generator

## Overview

The Ethereum Address Generator is a Go program designed to generate Ethereum addresses with a specified prefix and suffix. It leverages the `go-ethereum-hdwallet` library to handle Ethereum HD wallets and mnemonics. This project utilizes the Go programming language (Golang) to achieve performance through several key features of the language.

## Features

- Generate Ethereum addresses with custom prefixes and suffixes.
- Utilizes concurrent Goroutines for efficient address generation.
- Concurrency with Goroutines:

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Installation

Clone the repository:

```bash
git clone https://github.com/your-username/ethereum-address-generator.git
cd ethereum-address-generator
```

## Features

Generate Ethereum addresses with custom prefixes and suffixes.
Utilizes concurrent Goroutines for efficient address generation.
Getting Started
Prerequisites
Go installed on your machine.
Installation
Clone the repository:

```bash
git clone https://github.com/your-username/ethereum-address-generator.git
cd ethereum-address-generator
```

## Build and run the program:

```bash
go build -o eth-address-generator
./eth-address-generator
```

## Usage

```bash
./eth-address-generator [flags]
```

### Command-line Flags

--prefix: Prefix of the address.

--suffix: Suffix of the address.

#### Example

Generate Ethereum addresses with a specific prefix and suffix:

```bash
./eth-address-generator --prefix=0x68 --suffix=68
```

## Contribute

Contributions are welcome! If you find a bug or have a feature request, please open an issue.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
The project utilizes the go-ethereum-hdwallet library.
Contact
For inquiries, please contact nguyentruongkhang22@gmail.com.
