# Step 0

If you're really new to Go, then this step will help you. Otherwise, you can skip it.

In this step we'll create a very basic Go program that outputs "Hello World!"
and learn how to call this from the CLI.

## main.go

To get a Go program running you need to create a file under the `main` package.
Often this will be called `main.go`.

Then the entrypoint for a Go program is created by having a function called `main`,
with no arguments.

```go
package main

func main() {
	
}
```

## Writing output

While this works, it won't show you anything.

Within the `main()` function body add.

```go
fmt.Println("Hello World!")
```

More than likely your IDE will then import the correct package. 
VS Code uses 'gopls' while IntelliJ uses the 'Go language plugin'.

If not, you also have to write an _import statement_. (But honestly, try to find out how to enable Go language support)

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

## Running

Now that you have a proper Go program with an executable main, we can run it from the command line.

Make sure you `cd` into the `step0` folder, we can now use `go run <filename>`.

```shell
go run main.go
```
