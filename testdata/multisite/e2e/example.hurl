#### example ####
GET https://example.local/a.json
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.a" == "aaa"
