# md5http

![status](https://badgen.net/github/checks/iamolegga/md5http)

Tool that logs md5 of passed urls.

## Usage

Simply pass list of urls:

```shell
md5http http://google.com https://yandex.ru http://yahoo.co.uk 
```

Result:

```shell
http://google.com 79f71c2cdedf29d86fe700f56ee29ac3
https://yandex.ru 3adf250272cab94c30253e0b669ff755
http://yahoo.co.uk 34307f306afa6cad2223717162f82a29
```

It requests urls in parallel (10 requests by default).
You can change it with `-parallel` flag: 

```shell
md5http -parallel 1 http://google.com https://yandex.ru http://yahoo.co.uk
```

## Install

```shell
go install github.com/iamolegga/md5http/cmd/md5http@latest
```

## Development

Main tool is build with zero dependencies, but for generating mock objects for testing [gomock](https://github.com/golang/mock/) was used (required only for changing mock objects).
