# bashenv

Simple pkg to source a bash profile in go.

__example.sh:__

    export STR="Hello World!"

__example usage__

    bashenv.Source("example.sh")

    // The following will return "Hello World!"
    os.Getenv("STR")

## License
MIT
