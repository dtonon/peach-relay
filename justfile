dev:
  go run .

test:
  jq -c . event-signed.json | nak event ws://localhost:3334