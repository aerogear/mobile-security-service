#!/bin/sh

# Run some pre commit checks on the Go source code. Prevent the commit if any errors are found

# Format the Go code
check_code_format (){
    go fmt $(go list ./... | grep -v /vendor/)
}

# Check all files for errors
check_code_errors (){
    {
        errcheck -ignoretests $(go list ./... | grep -v /vendor/)
    } || {
        exitStatus=$?

        if [ $exitStatus ]; then
            printf "\nErrors found in your code, please fix them and try again."
            exit 1
        fi
    }
}

# Check all files for suspicious constructs
check_go_vet (){
    {
        go vet $(go list ./... | grep -v /vendor/)
    } || {
        exitStatus=$?

        if [ $exitStatus ]; then
            printf "\nIssues found in your code, please fix them and try again."
            exit 1
        fi
    }
}

# Calling the function
check_code_format
check_code_errors
check_go_vet
