# gql-bigint

[![Go Reference](https://pkg.go.dev/badge/github.com/xplorfin/gql-bigint.svg)](https://pkg.go.dev/github.com/xplorfin/gql-bigint)
[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/gql-bigint)
[![Tests](https://github.com/xplorfin/gql-bigint/actions/workflows/test.yml/badge.svg)](https://github.com/xplorfin/gql-bigint/actions/workflows/test.yml)
[![Linter](https://github.com/xplorfin/gql-bigint/actions/workflows/lint.yml/badge.svg)](https://github.com/xplorfin/gql-bigint/actions/workflows/lint.yml)

gql-bigint contains two scalar types (`BigInt` and `BigUInt`) for use with [gqlgen](https://github.com/99designs/gqlgen) GraphQL servers.

`BigInt` allows GraphQL servers to bypass the 53-bit limit imposed on GraphQL's inbuilt `Int` type, and corresponds to the 
Go type `int64`.

`BigUInt` is very similar to `BigInt`, except it corresponds to the Go type `uint64`, allowing for _slightly_ more headroom
if needed. 

## Usage

Install `gql-bigint` using `go get -u github.com/xplorfin/gql-bigint`.

Make sure the `autobind` entry in your `.gqlgen.yml` file has `"github.com/xplorfin/gql-bigint"` somewhere in it,
like so:

```yaml
autobind:
  - "github.com/xplorfin/gql-bigint"
  [... all your other autobinds ...]
```

Next, make sure the `models` entry in your `.gqlgen.yml` is set up to map the scalar types properly when generating:

```yaml
models:
  [... whatever else you have here ...]
  BigInt:
      model:
        - github.com/xplorfin/gql-bigint.BigInt
  BigUInt:
      model:
        - github.com/xplorfin/gql-bigint.BigUInt
```

Finally, put the following two lines into a `.graphql` file where gqlgen will be able to pick it up (I recommand `scalars.graphql`):

```graphql
scalar BigInt

scalar BigUInt
```

Run `go generate ./...` to regenerate your schema/resolvers, and you're done. 

## Contributing

Open up an issue or PR, and we'll take a look at it!