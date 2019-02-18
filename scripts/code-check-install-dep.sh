#!/bin/sh

echo "Install dependencies required to run the code checks..."

code_check_install_dep (){
    go get -u github.com/kisielk/errcheck
    dep ensure
}

# Calling the function
code_check_install_dep