# Kaniko

## Build image local PC
You need to:  
1. `docker login` beforehand
2. bind mount  `~/.docker/config.json` with credential (`docker login` adds
   credential).
3. bind mount `.netrc` if your project depends on private repositories.

The below is an example of command
```sh
docker run -it -v ~/.docker/config.json:/kaniko/.docker/config.json -v ~/.netrc:/root/.netrc -v $(pwd):/workspace/ --rm gcr.io/kaniko-project/executor:latest \
    --context dir:///workspace/ \
    --dockerfile /workspace/path/to/Dockerfile \
    --destination ${IMAGE_NAME}:${IMAGE_TAG}
```

## Image Caching
Specifying `--cache` and `--cache-dir` flag still pushes cache to remote.  
Hope someday it will be possible to store intermediate image on local machine.
