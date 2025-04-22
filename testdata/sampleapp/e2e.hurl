#### static contents ####
GET http://localhost:3000/a.json
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.a" == "aaa"

#### rewrite index.html for spa app ####
GET http://localhost:3000/
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
header "Content-Type" == "text/html; charset=utf-8"

GET http://localhost:3000/about
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
header "Content-Type" == "text/html; charset=utf-8"

GET http://localhost:3000/about/2
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
header "Content-Type" == "text/html; charset=utf-8"

GET http://localhost:3000/about/2/content
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
header "Content-Type" == "text/html; charset=utf-8"

GET http://localhost:3000/about/2/not-exists.txt
HTTP 404
[Asserts]
header "Cache-Control" == "no-cache"

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

#### old page ####
GET http://localhost:3000/old
HTTP 301
[Asserts]
header "Cache-Control" == "no-cache"
header "Location" == "/"
