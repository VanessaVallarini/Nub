# calculate-tax

## About
Command line program (CLI) that calculates the tax to be paid on profits or losses from operations in the financial stock market.

## Technologies
* Golang 1.18

## Development requirements
* Visual Studio Code
* Golang (https://golang.org/doc/install) 

## Directory Structure
- `models`
     - Contains the structures that represent the objects used.
- `utils`
     - Contains functionality that can be used throughout the program.

## Running
- Synchronize the dependencies:
```bash
make clean 
```
---------------
- Builds application:
```bash
make build 
```
---------------
- Runs application:
```bash
make run
```

## Runs tests
```bash
make test
make test-alt
make test-cover
make test-out
```