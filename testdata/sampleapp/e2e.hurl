GET http://localhost:3000/a.json
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.a" == "aaa"

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

GET http://localhost:3000/about/about
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
header "Content-Type" == "text/html; charset=utf-8"

