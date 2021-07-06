# the below is used as-is. 
openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true,packageName=openapi  --global-property=models -o ../../pkg/connectors/taxonomy_models_codegen -i test4.yaml

# the below is also modified to suit for our purposes 
# below code is the based for implementing the http client 
# openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true,packageName=openapi  --global-property=apis,supportingFiles -o ../../pkg/connectors/taxonomy_client_apis_codegen2 -i test4.yaml

# below code is the based for implementing the http server 
# openapi-generator-cli generate -g go-server  --additional-properties=packageName=openapiserver,serverPort=8081,sourceFolder=openapiserver  --global-property=apis,supportingFiles -o out_go_server -i test4.yaml