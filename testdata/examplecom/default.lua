host = "example.com"

entry.host = "https://example.com"

function handler(next, req)
    print('handle')
    if (req.path == "/favicon.ico") then
        req.path = "/aaa"
    end
    print(req.path)

    res = next(req)
    res.status = 201

    return res
end
