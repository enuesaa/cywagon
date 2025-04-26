#### rewrite for storage ####
GET http://localhost:3000/storage/aaa/a.json
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.a" == "aaa"

GET http://localhost:3000/storage/aaa/b.json
HTTP 404
[Asserts]
header "Cache-Control" == "no-cache"
