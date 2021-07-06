# Code generated due to the below command is used as-is WITHOUT any modification
openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true,packageName=openapi  --global-property=models -o ../../pkg/connectors/taxonomy_models_codegen -i m4d-policy-manager-taxonomy.yaml

# Code generated due to the below 2 commands is MODIFIED to suit our purposes

# below code is the based for implementing the http client 
# openapi-generator-cli generate -g go --additional-properties=prependFormOrBodyParameters=true,packageName=openapi  --global-property=apis,supportingFiles -o ../../pkg/connectors/taxonomy_client_apis_codegen2 -i m4d-policy-manager-taxonomy.yaml

# below code is the based for implementing the http server 
# openapi-generator-cli generate -g go-server  --additional-properties=packageName=openapiserver,serverPort=8081,sourceFolder=openapiserver  --global-property=apis,supportingFiles -o out_go_server -i m4d-policy-manager-taxonomy.yaml