#### restricted page ####
GET http://localhost:3000/restrict/main.json
HTTP 401
[Asserts]
header "Cache-Control" == "no-cache"
header "WWW-Authenticate" == "Basic realm=\"Restricted\""

GET http://localhost:3000/restrict/main.json
Authorization: Basic dGVzdDp0ZXN0
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.restrict" == "restrict"
