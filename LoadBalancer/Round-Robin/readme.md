# Build Docker Image

```sh
docker compose -f round-robin.compose.yaml build
```

# Run Docker containers

```sh
docker compose -f round-robin.compose.yaml up
```

Now try to hit following url
`http://localhost:8080/`
you will get a response from backend node in round robin
