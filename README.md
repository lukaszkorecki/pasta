Ephemeral (copy) pasta :spaghetti:

Quick and dirty pastebin in which all pastes expire after 5 minutes.

Not sure why.

### Building

`go install github.com/lukaszkorecki/pasta`

### Example

```bash

id=$(curl -X POST -d'body="lol"' localhost:3000/paste)

curl localhost:3000/$id # => lol

# wait 5 minutes


curl localhost:3000/$id # => Gone, 404
```

# Licence

GPLv3
