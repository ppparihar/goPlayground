# Build Docker Image

```sh
docker compose -f consistent-hashing.compose.yaml build
```

# Run Docker containers

```sh
docker compose -f consistent-hashing.compose.yaml up
```

Now try to hit following url
`http://localhost:8080/`
you will get a response from backend node in consistent hashing
