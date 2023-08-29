# Scratch container with ssl, non-root user, and CGO

Build with:

```bash
docker build . -t userdrop
```

Run with:

```bash
docker run --rm --cap-drop=all userdrop
```

## Added configurations on top of normal scratch

1. Binary build with CGO enabled by leveraging static musl that comes default in alpine containers
2. Added ca-certificates package to a well-known path and then copied into scratch container for Go to load
3. Added a basic passwd file to drop to a non-root user
4. Run with all capabilities dropped
5. For image size with an added initial memory footprint and boot-up time, packed the statically built executable with UPX