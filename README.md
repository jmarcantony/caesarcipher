# caesarcipher
Go package to encode and decode strings to and from caesar cipher

## Installation
```
go get github.com/jmarcantony/caesarcipher
```

## Examples
### Encode to caeser cipher
```go
toEncode := "caesercipher"
encoded := caesarcipher.Encode(toEncode, 5) // returns "hfjxjwhnumjw"
```

### Decode from caeser cipher
```go
toDecode := "hfjxjwhnumjw"
decoded := caesarcipher.Decode(toDecode, 5) // returns "caesercipher"
```

