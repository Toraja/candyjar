# Kaniko

## Build image on local PC
You need to:  
1. `docker login` beforehand
2. bind mount  `~/.docker/config.json` with credential (`docker login` adds
   credential).
3. bind mount `.netrc` if your project depends on private repositories.

The below is an example of command
```sh
docker run -it --rm \
    -v ~/.docker/config.json:/kaniko/.docker/config.json \
    -v ~/.netrc:/root/.netrc \
    -v $(pwd):/workspace/ \
    gcr.io/kaniko-project/executor:latest \
        --context dir:///workspace/ \
        --dockerfile relative/path/to/Dockerfile \
        --destination ${IMAGE_NAME}:${IMAGE_TAG}
```

## Image Caching
Cache is basically stored on remote registry and not on local machine.  
That is done by simply specifying `--cache` flag.  

However using `kaniko-warmer`, you can cache base images on local machines.  
First, run the below command to initialise (store) cached image.  
(cache is stored in `/tmp/kaniko/cache` in this case)  
```sh
docker run -it --rm -v /tmp/kaniko/cache:/cache gcr.io/kaniko-project/warmer:latest \
    --image=golang:1.17 \
    --image=alpine:3.15
```

Then run `kaniko-executer` with the cache mounted to cache dir of the
container. (`/cache` unless specified in `--cache-dir`)  
```sh
docker run -it --rm \
    ...
    -v /tmp/kaniko/cache:/cache \
    gcr.io/kaniko-project/executor:latest \
    ...
```
