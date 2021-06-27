rm -rf out_go_server
rm -rf out_go_client


openapi-generator-cli generate -g go-server  --additional-properties=packageName=openapiserver,serverPort=8081,sourceFolder=openapiserver  --global-property=apis,supportingFiles -o out_go_server -i test5.yaml


# openapi-generator-cli generate -g go-echo-server  --additional-properties=packageName=openapiserver,serverPort=8081,sourceFolder=openapiserver  --global-property=apis,supportingFiles -o out_go_echo_server -i test5.yaml


openapi-generator-cli generate -g go  -o out_go_client -i test5.yaml
