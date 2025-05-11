#### example ####
GET https://example.local:8443/a.json
HTTP 200
[Asserts]
header "Cache-Control" == "no-cache"
jsonpath "$.a" == "aaa"
