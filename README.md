version
-------

A small package for exposing your Go library's [semver](http://semver.org/) to others consuming it.

[![Build Status](https://drone.io/github.com/chuckpreslar/version/status.png)](https://drone.io/github.com/chuckpreslar/version/latest)

## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/version

## Example Usage

```go
package main

import (
  "github.com/chuckpreslar/version"
)

// Version.Major()   => 1
// Version.Minor()   => 5
// Version.Patch()   => 2
// Version.Label()   => "rc1"
// Version.String()  => "1.5.2-rc1"
var Version = version.Establish(1, 5, 2, "rc1")
```

## Documentation

View godoc's or visit [godoc.org](http://godoc.org/github.com/chuckpreslar/version).

    $ godoc version

## License

> The MIT License (MIT)

> Copyright (c) 2016 Chuck Preslar

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
