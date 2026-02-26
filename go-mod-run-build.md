# go mod init
    => it will give helping messages, but we have to pass module name.
    'go mod init example.com/m' to initialize a v0 or v1 module
    'go mod init example.com/m/v2' to initialize a v2 module


# go mod init requires a module name, and you didn't pass one, and go can't infer this automatically.
    => Correct usages:
    go mod init your-module-name

    ## Example:
    go mod init github.com/ankushkumar/my-awesom-app
    => this will set up go.mod file 


# go code coverage & unit-test cases
