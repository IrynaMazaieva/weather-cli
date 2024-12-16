# weather-cli

CLI application in Go that accepts a country and a city as input and returns the weather information for the current day. The output will display the results from the API that responded the fastest.

## Installation

To run the project locally you need to have go installed.
Clone the repository and run the following command to install the dependencies:
```bash
go mod download
go install
``` 

## Usage

```bash
weather-cli [country] [city]
```

## To run tests

```bash
go test ./pkg/api
```

## Example

```bash
weather-cli 'Great Britain' London
```
