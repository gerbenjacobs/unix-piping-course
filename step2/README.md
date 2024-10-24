# Step 2

In this step we will convert our incoming string to a JSON object and write that to 'standard out' `stdout`.

We will use `jq` afterward to pretty print this.

## How to write to stdout

We've learned to use a `io.Reader` in combination with `stdin`, but for writing we'll need to use a `io.Writer`.

Since we know we're going to use JSON we can have a look at the `encoding/json` [package](https://pkg.go.dev/encoding/json).

You're probably familiar with `json.Marshal` but the package also comes with an `json.Encoder`,
that needs to be initialized with a `io.Writer` interface.   
We can use the `os.Stdout` file descriptor for this, cause as you remember, `os.File` structs adhere to the interface.

## Our JSON model

Let's pretend we're going to make a custom model that has its own JSON schema. This is what we'll output.

A model in Go is basically just a `struct`, but by adding _struct field tags_ we can make it even better
for JSON marshalling.    
You can read about those json tags in the [Marshall documentation](https://pkg.go.dev/encoding/json#Marshal).

```go
type Output struct {
	Message string `json:"message"`
	Length  int    `json:"len"`
}
```

## Exercise

1. Copy your `main.go` from `step1`
2. Create your custom JSON model (_note: you can't call it `Output`, that's already used in `ref.go`_)
3. Remove the `fmt.Printf` and replace it with a `json.Encoder` (see `json.NewEncoder()`)
4. Call `Encode()` on your encoder with your custom model as input (make sure to fill it with data)

### Run the exercise

```bash
echo -n "Hello World" | go run main.go
```

If you have `jq` installed on your machine you can even try

```bash
echo -n "Hello World" | go run main.go | jq
```

## Did you know..

- Calling `Encode()` on our encoder doesn't return anything but a potential error, do you know why?
- What happens with our output JSON if we remove the entire `json:""` tags?
- Read the json tag documentation about `omitempty` and `"-"`, play around with that and watch what happens to your JSON
- (Optional) Oh no! You're using ancient Microsoft services and need `XML`. 
   See if you can use [encoding/xml](https://pkg.go.dev/encoding/xml) instead!