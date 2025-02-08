host = "example.com"

entry.host = "https://example.com"

function handler(next, req, res)
    print('handle')
    print(res.status)
    next()

    res.status = 200
    return res
end
