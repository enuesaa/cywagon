host = "example.com"

entry.host = "https://example.com"

function handler(next, req)
    print('handle')

    res = next(req)
    res.status = 201

    return res
end
