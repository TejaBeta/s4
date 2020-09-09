# s4 [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens) [![Go Report Card](https://goreportcard.com/badge/github.com/tejabeta/s4)](https://goreportcard.com/report/github.com/tejabeta/s4)

A tiny CLI to serve static websites from **`AWS S3`** Object store bucket with `index.html` as the entry file. 

Could work with private buckets. Prerequisites, make sure you have appropriate `IAM` access to the bucket and `index.html` inside the bucket.

### Installation

Execution of `make build` creates an executable inside `bin` directory insidse same repo.

### Run locally

`s4 -region "myRegion" -bucket "myBucket" -accessKey "myAccessKey" -secretKey "mySecretKey"`


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

### Future improvements
- Ability to use different cloud platforms
- Ability to support dynamic server side scripting

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.