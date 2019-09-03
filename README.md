# gocleanarch
Generator for clean architecture implementation in Golang.
The clean architecture implementation adopted from [here](https://github.com/bxcodec/go-clean-arch). There are 4 main components:
1. Model
2. Repository
3. Usecase
4. Delivery

## Installation
To install the cli, run:

```bash
go get github.com/martinock/gocleanarch
```

## Usage
There 4 main generate commands for this cli:
These files will be written into `<YOUR_APP_DIRECTORY>/internal` make sure you don't have files/folders with the same name in the folder

1. Model
To generate golang model, use:
```bash
gocleanarch model
```

2. Repository
To generate golang repository, use:
```bash
gocleanarch repo
```

3. Usecase
To generate golang usecase, use:
```bash
gocleanarch usecase
```

4. Delivery
To generate golang delivery, use:
```bash
gocleanarch delivery
```
