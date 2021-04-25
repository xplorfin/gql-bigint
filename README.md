# gql-bigint

[![Renovate enabled](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://app.renovatebot.com/dashboard#github/xplorfin/gql-bigint)
[![Test status](https://github.com/xplorfin/gql-bigint/workflows/Tests/badge.svg)](https://github.com/xplorfin/gql-bigint/actions?query=workflow%3ATests)
[![Lint status](https://github.com/xplorfin/gql-bigint/workflows/Lint/badge.svg)](https://github.com/xplorfin/gql-bigint/actions?query=workflow%3ALint)

gql-bigint contains two scalar types (`BigInt` and `BigUInt`) for use with [gqlgen](https://github.com/99designs/gqlgen) GraphQL servers.

`BigInt` allows GraphQL servers to bypass the 53-bit limit imposed on GraphQL's inbuilt `Int` type, and corresponds to the 
Go type `int64`.

`BigUInt` is very similar to `BigInt`, except it corresponds to the Go type `uint64`, allowing for _slightly_ more headroom
if needed. 

## Usage

Install `gql-bigint` using `go get -u github.com/xplorfin/gql-bigint`

## Contributing

Open up an issue or PR, and we'll take a look at it!