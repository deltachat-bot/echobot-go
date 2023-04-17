#  deltabot-cli-go template

![Latest release](https://img.shields.io/github/v/tag/deltachat-bot/echobot-go?label=release)
[![CI](https://github.com/deltachat-bot/echobot-go/actions/workflows/ci.yml/badge.svg)](https://github.com/deltachat-bot/echobot-go/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-22.2%25-red)
[![Go Report Card](https://goreportcard.com/badge/github.com/deltachat-bot/echobot-go)](https://goreportcard.com/report/github.com/deltachat-bot/echobot-go)

This is a template project using the [deltabot-cli-go](https://github.com/deltachat-bot/deltabot-cli-go) library.

## Install

Binary releases can be found at: https://github.com/deltachat-bot/echobot-go/releases

To install from source:

```sh
go install github.com/deltachat-bot/echobot-go@latest
```

### Installing deltachat-rpc-server

This program depends on a standalone Delta Chat RPC server `deltachat-rpc-server` program that must be
available in your `PATH`. For installation instructions check:
https://github.com/deltachat/deltachat-core-rust/tree/master/deltachat-rpc-server

## Running the bot

Configure the bot:

```sh
echobot init bot@example.com PASSWORD
```

Start listening to incoming messages:

```sh
echobot serve
```

Run `echobot --help` to see all available options.
