# goit

Get a OIDC token from your terminal.

[![hydrun CI](https://github.com/pojntfx/goit/actions/workflows/hydrun.yaml/badge.svg)](https://github.com/pojntfx/goit/actions/workflows/hydrun.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/pojntfx/goit.svg)](https://pkg.go.dev/github.com/pojntfx/goit)
[![Matrix](https://img.shields.io/matrix/goit:matrix.org)](https://matrix.to/#/#goit:matrix.org?via=matrix.org)
[![Binary Downloads](https://img.shields.io/github/downloads/pojntfx/goit/total?label=binary%20downloads)](https://github.com/pojntfx/goit/releases)

## Overview

`goit` is a tool to quickly get a OpenID Connect token from your terminal, e.g. to use it for API authentication.

## Installation

Static binaries are available on [GitHub releases](https://github.com/pojntfx/goit/releases).

On Linux, you can install them like so:

```shell
$ curl -L -o /tmp/goit "https://github.com/pojntfx/goit/releases/latest/download/goit.linux-$(uname -m)"
$ sudo install /tmp/goit /usr/local/bin
```

On macOS, you can use the following:

```shell
$ curl -L -o /tmp/goit "https://github.com/pojntfx/goit/releases/latest/download/goit.darwin-$(uname -m)"
$ sudo install /tmp/goit /usr/local/bin
```

On Windows, the following should work (using PowerShell as administrator):

```shell
PS> Invoke-WebRequest https://github.com/pojntfx/goit/releases/latest/download/goit.windows-x86_64.exe -OutFile \Windows\System32\goit.exe
```

You can find binaries for more operating systems and architectures on [GitHub releases](https://github.com/pojntfx/goit/releases).

## Usage

You can get the OIDC token by running `goit`; if you have a browser available, it will launch it for you to authorize, and if not will print an authorization URL in the console:

```shell
$ goit --oidc-client-id='Ab7OLrQibhXUzKHGWYDFieLa2KqZmFzb' --oidc-issuer='https://pojntfx.eu.auth0.com/' --oidc-redirect-url='http://localhost:11337'
2022/04/24 13:54:44 Successfully got the following OIDC access token:
ayJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlFtSXdJNVprRjasfasdfadf9.eyJpc3MiOiJodHRwczovL3Bvam50ZnguZXUuYXV0aDAuY29tLyIsInN1YiI6ImF1dGgwfDYyMzgxZTRkN2I3ZGE5MDA2YmQ1NzA5MCIsImF1ZCI6IkFiN09MclFpYmhYVXpLSEdXWURGaWVMYTJLcVptRnpiIiwiaWF0IjoxNjUwODAxMjg0LCJleHAiOjE2NTA4MzcyODR9.h7mXPs0gWQVZao_TyZu7VN1mUReJYBitnMKfgVHE6cLjN-TQX3tRwrasdfsdf-patqRWlBKDZZVuul1WZnArYHHywbzSYMThUfGRlif_QVNNzH2wNVuF_9rjcrVHDcjGZ0Wdta28tPRjSTfKjB6OAeK_8a8gxbk9XJ_tk8bYZfN-4HvyRl_Xrrd-xlMAZm60euNmzd9f3w
```

It is also possible to export the token to a environment variable like so:

```shell
$ export OIDC_CLIENT_ID='Ab7OLrQibhXUzKHGWYDFieLa2KqZmFzb' OIDC_ISSUER='https://pojntfx.eu.auth0.com/' OIDC_REDIRECT_URL='http://localhost:11337'
$ export API_KEY="$(goit)"
2022/04/24 13:57:40 Successfully got the following OIDC access token:
$ echo "${API_KEY}"
ayJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlFtSXdJNVprRjasfasdfadf9.eyJpc3MiOiJodHRwczovL3Bvam50ZnguZXUuYXV0aDAuY29tLyIsInN1YiI6ImF1dGgwfDYyMzgxZTRkN2I3ZGE5MDA2YmQ1NzA5MCIsImF1ZCI6IkFiN09MclFpYmhYVXpLSEdXWURGaWVMYTJLcVptRnpiIiwiaWF0IjoxNjUwODAxMjg0LCJleHAiOjE2NTA4MzcyODR9.h7mXPs0gWQVZao_TyZu7VN1mUReJYBitnMKfgVHE6cLjN-TQX3tRwrasdfsdf-patqRWlBKDZZVuul1WZnArYHHywbzSYMThUfGRlif_QVNNzH2wNVuF_9rjcrVHDcjGZ0Wdta28tPRjSTfKjB6OAeK_8a8gxbk9XJ_tk8bYZfN-4HvyRl_Xrrd-xlMAZm60euNmzd9f3w
```

🚀 **That's it!** We hope you enjoy using goit.

Be sure to check out the [reference](#reference) for more information.

## Reference

### Command Line Arguments

```shell
$ goit --help
Get a OIDC token from your terminal.

Find more information at:
https://github.com/pojntfx/goit

Usage:
  goit [flags]

Flags:
  -h, --help                       help for goit
      --oidc-client-id string      OIDC Client ID (i.e. myoidcclientid) (can also be set using the OIDC_CLIENT_ID env variable)
      --oidc-issuer string         OIDC Issuer (i.e. https://pojntfx.eu.auth0.com/) (can also be set using the OIDC_ISSUER env variable)
      --oidc-redirect-url string   OIDC Redirect URL (default "http://localhost:11337")
  -v, --verbose                    Enable verbose logging
```

### Environment Variables

All command line arguments described above can also be set using environment variables; for example, to set `--oidc-client-id` to `myclientid` with an environment variable, use `OIDC_CLIENT_ID=myclientid`.

## Acknowledgements

- [coreos/go-oidc](https://github.com/coreos/go-oidc) provides the OpenID Connect library for goit.

## Contributing

To contribute, please use the [GitHub flow](https://guides.github.com/introduction/flow/) and follow our [Code of Conduct](./CODE_OF_CONDUCT.md).

To build goit locally, run:

```shell
$ git clone https://github.com/pojntfx/goit.git
$ cd goit
$ make depend
$ make
$ out/goit
```

Have any questions or need help? Chat with us [on Matrix](https://matrix.to/#/#goit:matrix.org?via=matrix.org)!

## License

goit (c) 2023 Felicitas Pojtinger and contributors

SPDX-License-Identifier: AGPL-3.0
