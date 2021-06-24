rm -rf out_go_server
rm -rf out_go_client
openapi-generator-cli generate -g go-server  -o out_go_server -i test4.yaml
openapi-generator-cli generate -g go  -o out_go_client -i test4.yaml
