# Rate limiting Server


## Quickstart

```bash
make run

for ((i=1; i<=70; i++)); do curl localhost:8080; done;
```

## Testing

```bash
make test
```

## Algorithms

See [wiki](https://en.wikipedia.org/wiki/Token_bucket)


## TODO

* IP Mask
* Middleware
