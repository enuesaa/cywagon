hostname = "example.com"
port = 3000

function handler(req)
    local content = loadfrom("/var/www/html/aaa")
    --- response
    return content
end
