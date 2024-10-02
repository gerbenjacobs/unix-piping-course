# Step 1

In this step we're going to read __incoming__ data and print this out, so we can verify it works.  
We will be using 'standard in' (`stdin`) which is a common [standard stream](https://en.wikipedia.org/wiki/Standard_streams) 
used in Unix-based systems.

For this we will use `echo` as our input provider. We're adding the `-n` flag to prevent newlines from being sent.

## How to read from stdin

Go has the built-in `os` [package](https://pkg.go.dev/os) that comes with variables for the 3-main I/O-streams.

We will use `os.Stdin`, this returns a `*File` [file descriptor](https://en.wikipedia.org/wiki/File_descriptor) which
itself is implementing the `Reader` and `Writer` interfaces.   
A lot of methods in Go will adhere to these interfaces,
this allows you to pick a variety of options when dealing with reading and writing. (Buffers, Scanners etc.)

Since we are in _the first step_, we will not care about buffer sizes, streaming and/or memory overflow.
(These are things you should consider in _production_)

We will use `io.ReadAll` a method from the `io` package that just reads the __entire__ thing it's given,   
since we  are giving it the `Stdin` file descriptor, it will read whatever we inputted on our Unix's standard in.

## Exercise

1. Create a main.go using `io.ReadAll` and `os.Stdin`
2. Use `fmt.Printf` to print whatever you received 
3. (Optional) As an added bonus, output the length of the string too

### Run the exercise

Use the following to test your application:
```bash
echo -n "Hello World" | go run main.go
```

## Did you know..

- When you use `fmt.Print` (and its siblings) underwater Go uses `os.Stdout`
- Try removing the `-n` from our `echo` command and see what effect it has
- What happens when we run our script without input?
- There's also a `ReaderWriter` interface and internally it just embeds both `Reader` and `Writer` interfaces