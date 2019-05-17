# ApiEndPoint
ApiEndPoint in GoLang which identify if A URL address is located in some rest end point

1. config the server port (default_port=9087)and the endpoint URL, they both located in main.go after the import section
2. in order to build the project run: go build .\src\app\main.go
3. the rest endpoint API will receive a POST request with string contains domain+path,
it will return string containing the result 
