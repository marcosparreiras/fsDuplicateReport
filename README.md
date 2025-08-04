# fsDuplicateReport

A simple CLI tool to scan and report duplicate files in your filesystem.

## Installation 

```sh
git clone git@github.com:marcosparreiras/fsDuplicateReport.git
cd fsDuplicateReport
go build -o fsDuplicateReport
```

## Usage
```sh
./fsDuplicateReport -path '/home/marcos/Documents' -ignore '.go' -ignore 'configurations'
```
