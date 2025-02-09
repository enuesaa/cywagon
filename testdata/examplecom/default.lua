host = "example.com"

entry.host = "https://example.com"

function handler(req, next)
    print('handle')

    res = next(req)
    -- print(res)
    res.status = 201

    return res
end
