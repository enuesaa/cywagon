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
