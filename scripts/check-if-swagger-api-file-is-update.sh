#!/bin/bash

# return error if the swagger file was not updated
check-if-swagger-api-file-is-update (){
    echo "Checking swagger api documentation ..."
    # create a copy of the original file
    cp api/swagger.yaml api/original.yaml
    # run the make file task to update the swagger file definition
    make build_swagger_api
    # Check if it is updated
    if ! cmp api/swagger.yaml api/original.yaml >/dev/null 2>&1
    then
        echo "The REST API Swagger doc was updated and it was not committed."
        echo "Please, run 'make setup' or 'build_swagger_api' and commit the changes in the file 'api/swagger.yaml'."
        exit 1
    fi

}

# Calling the function
check-if-swagger-api-file-is-update