# `olsgo`

Ordinary-least-squares bivariate regression in Go.

# Use

To use what's already created:

1. First install

```sh
go get github.com/w-decker/olsgo@latest
```

2. Then import like you would any other library
```go
package main

import "github.com/w-decker/olsgo"
```
3. Go to the [example](/example/main.go) to see how to execute `olsgo` and get some output.

# Output

1. `.txt` file
```txt
OLSGO Output 
Intercept: 46.0917        B1: 0.1461         Pearson's r: 0.9471         Residual variance: 0.1030         
```

2. Some plots (more coming later)


![[Raw data]](/example/plot.png)


