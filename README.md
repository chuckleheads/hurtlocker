# Hurtlocker

gRPC server with built in swagger support. This is far from being anywhere near functional.

## Making changes

#### Generating go code from protobufs
*** You need to do this for each proto file you change.
```
make protobufs
 ```

#### Generating swagger defs
```
make swagger
```

If you just run `make` by itself, it will run both of the above commands.
