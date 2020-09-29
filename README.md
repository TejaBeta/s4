# s4 [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) [![Go Report Card](https://goreportcard.com/badge/github.com/tejabeta/s4)](https://goreportcard.com/report/github.com/tejabeta/s4)

A CLI tool that serves as a middleware to Cloud provisioned object storage services such a **`AWS S3`**. Currently supports AWS S3, with the following features such as serving a static website with `index.html` as the entry file from the object storage.

Could work with private buckets. Prerequisites, make sure you have appropriate `IAM` access to the bucket and `index.html` inside the bucket.

### Installation

Execution of `make build` creates an executable inside `bin` directory insidse same repo.

### Run locally

`s4 static --region="myRegion" --bucket="myBucket" --accessKey="myAccessKey" --secretKey="mySecretKey"`


### CLI Options

command | type | default value| Definition
--------|------|--------------|------------
`isAWS` | `bool`  | `true`  | A boolean flag to pick a platform
`bucket` | `string` | "" | Bucket name from where tool has to read
`accessKey` | `string` | "" | IAM Access Key for the tool to read a private bucket
`secretKey` | `string` | "" | IAM Secret Key for the tool to read a private bucket
`region`  | `string` | "" | AWS Region where the bucket resides 
`address` | `string` | "127.0.0.1:3000" | Local host address pages are served
`autoUpdate` | `bool` | `true` | A boolean flag to enable to disable periodic checks bucket checks. By default it is 15 mins. 

## Main Menu

```

$s4 --help

A tiny CLI tool to that acts as a middleware to build
services making use of AWS S3 object store as a backend. 

Currently supports hosting a static website from private AWS S3
object store with pointing to index.html. And also supports the
hosting a private PyPi server with S3 object store as package storage.

Usage:
  s4 [flags]
  s4 [command]

Available Commands:
  fileserver  Serves the purpose of fileserver from a specific handler
  help        Help about any command
  static      Static serves a website from a specific handler

Flags:
      --accessKey string   AWS access key
      --address string     address:port to serve the s3 content (default "127.0.0.1:3000")
      --autoUpdate         Bool to auto update (default true)
      --bucket string      S3 bucket name
      --config string      config file (default is $HOME/.s4.yaml)
  -h, --help               help for s4
      --isAWS              Bool to pick a platform (default true)
      --localDir string    Local directory to sync and serve (default "./local")
      --region string      AWS Region Bucket resides
      --secretKey string   AWS secret key

Use "s4 [command] --help" for more information about a command.

```

### Future improvements
- Ability to use different cloud platforms
- Ability to support dynamic server side scripting
- Ability to support other features such as S3 back private package repositories
- Ability to archieve and pull logs in regular intervals
- Ability to support as a backend storage system to containers

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.