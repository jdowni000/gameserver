# gameserver

Simple web server that hosts information from a json file on port 8080 of local   container/image/or packages for a kubernetes spec.

## Makefile
The makefile will do most features for you automatically  
`make build` will build a Docker Image  
`make push` will push previous built image  
`make test` will runn all tests against the repository  
`make run` will run the local image created in build step  
`make kubernetes` will automatically create a kubernetes spec file  

## Endpoints
Some common endpoints for the simple application with provided json file are;  
`localhost:8080/`  
`localhost:8080/list`  
`localhost:8080/game?id=${1}`  
`localhost:8080/game?id=${2}`  
