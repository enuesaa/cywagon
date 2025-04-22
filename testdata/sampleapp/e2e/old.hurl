#### old page ####
GET http://localhost:3000/old
HTTP 301
[Asserts]
header "Cache-Control" == "no-cache"
header "Location" == "/"
