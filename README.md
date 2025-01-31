# 🍑 Peach Relay

A proof of concept relay to manage [NIP-69](https://github.com/nostr-protocol/nips/blob/master/69.md) orders.

## Run it

```shell
git clone git@github.com:dtonon/peach-relay.git
cd peach-relay
go run .
```

## Test it

```shell
jq -c . event-signed.json | nak event ws://localhost:3334
```
