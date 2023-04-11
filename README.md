# Randomly-generated JSON

Generates syntactically-correct,
yet semantically-meaningless JSON.

Does not require any non-standard library packages.

## Building and Running

```sh
$ go build rj1.go
$ ./rj1 > random.json
$ jq . random.json
```
