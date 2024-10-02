# Unix Piping Course with Go

This course will show you how to make a Go binary that acts as a unix tool.

We will read data from `stdin`, convert it to our custom JSON model and push it through to `stdout`, completing the cycle.

_This is a beginner to intermediate course for both Go and Unix, some things will be explained, but if you're unsure
you should be able to autonomously search for the right information._

## Background

The [Unix philosophy](https://en.wikipedia.org/wiki/Unix_philosophy) is to have tools only do one thing and by 
reading and writing from `stdout` these tools can be combined to create pipelines of work that needs to be done.

An example you might have seen already is using `curl` to query a JSON endpoint and using `jq` to pretty-print
this to your screen.

```bash
curl -s https://dummyjson.com/test | jq
```

A lot of these small, do one thing well, tools are written in C, but there's no reason we can't make our own in Go!

## Requirements

- You'll need to have Go installed. (See [Go install](https://go.dev/doc/install))
- You'll need to run this in a Unix-based system (assuming macOS)
- (Optional) Docker
- (Optional) We will use `jq` as well

## Steps

This course is divided into steps, navigate to `step1` and make sure to read the individual READMEs.

The README explains what needs to be done and what the exercise is, it also contains background information.
If something is not covered don't be afraid to search for more information on Wikipedia, pkg.go.dev or StackOverflow.

Each step also contains a `ref.go` which is a reference and/or the answer for that step.
Reading this before you do the exercise is considered cheating! ;)