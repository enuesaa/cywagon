host = "example.com"

entry.host = "https://example.com"

function handler(req)
    print('handle')

    res = req.invoke()
    res.status = 201

    return res
end
