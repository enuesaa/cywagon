host = "example.com"

entry.host = "https://example.com"
entry.secure = true

healthCheck.protocol = "HTTP"
healthCheck.path = "/"

function handler(next, req)
    if (req.path == "/favicon.ico") then
        req.path = "/aaa"
    end

    res = next(req)
    res.status = 201

    return res
end
