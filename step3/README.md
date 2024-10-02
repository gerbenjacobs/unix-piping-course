# Step 3

How can we make this look more like a Unix tool?

In this step we're going to build a binary for our system (assuming MacOS) and for a Unix-based system too.


## Compiling Go

By default, Go will compile your binaries for the current OS and architecture.

```bash
go build main.go
```

This will create a binary called `main` (or `main.exe` on Windows) in your local folder.

Add the `-o Output` flag to change the name

```bash
go build -o json_converter main.go
```

You will now have a binary called `json_converter`.

### Go install

There's also the `go install` command, which besides building, also places your binary in the `$GOPATH` location.

If you have this path in your OS's `$PATH`, then you can call this binary from anywhere.

_Since I assume most have installed Go via Brew and not set up the paths correctly, we will use our local binary_

### Cross Compiling

This all works well for your system, but not if you're trying to create binary for production, where more often than not
the systems are running Unix.

Go has the ability to cross compile for other OSes and Architectures. You can run `go tool dist list` to see a list of options.

Run `go tool dist banner` to see what you're running. 
For example macOS is called `darwin`, but you might see `amd` or `arm` based on whether 
you have an Intel chip or the newer Apple M1 chip.

_Go, as well as macOS, sometimes has the ability to translate for different architectures. 
So it might still run on wrong architectures._

Building for other architectures is done using the environment variables `GOOS` and `GOARCH`.

```bash
GOOS=windows GOARCH=386 go build -o json_converter.exe
```

Here we are building our binary for Windows machines with Intel chips, we have to manually add the `.exe` suffix.

## Exercise

1. Build your binary and name it `json_converter`
2. Build your binary again, but use `GOOS=linux` and name it `json_converter_unix`

### Run the exercise

Assuming we haven't set up our paths correctly, we will use our local binaries, requiring us to prefix it with `./`.

```bash
echo -n "Hello World" | ./json_converter | jq
```

Let's try it with our Unix version

```bash
echo -n "Hello World" | ./json_converter_unix | jq
```

Notice how you get `exec format error`.

We can test our Unix version on a small Alpine docker image. It doesn't come with `jq` though.

```bash
docker run -v $(pwd)/json_converter_unix:/usr/local/bin/json_converter:z alpine /bin/sh -c 'echo -n "Hello World" | json_converter'
```

We are mounting (`-v`) our binary to the `/usr/local/bin` directory and use `sh` for our pipe example.  
_The use of `sh -c` is required because of I/O issues with running in Docker, outside of this course's scope._

## Did you know..

- The OS and ARCH Go uses can also be found by the unix tool `uname` and `uname -m`
- Running `go run <file>` actually builds a temporary binary before running it