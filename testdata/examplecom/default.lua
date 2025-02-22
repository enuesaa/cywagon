host = "example.com"

origin.url = "https://example.com"

healthCheck.protocol = "HTTP"
healthCheck.path = "/"
healthCheck.matcher = "200-300"

function handler(next, req)
    if (req.path == "/favicon.ico") then
        req.path = "/aaa"
    end

    res = next(req)
    res.status = 201

    return res
end
