# awx-consul-inventory

Better Consul Inventory script to be called from AWX

## Building

### Mac/Linux

0) set GOROOT environment variable
1) Install Go and Make
2) make

### Docker

0) set GOROOT environment variable
1) Install Docker, Go and Make
2) make docker


## Running

### Mac/Linux

```
./awx-consul-inventory
```

### Docker

```
docker pull maguec/awx-consul-inventory:latest
docker run -i -t -p 8080:8080 maguec/awx-consul-inventory
```

## Testing

run either the docker container or the raw application binary

```
curl http://localhost:8080/health
```

---
Copyright © 2018, Chris Mague
