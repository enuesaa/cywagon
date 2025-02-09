host = "example.com"

entry.host = "https://example.com"

function handler(req)
    print('handle')

    res = invoke(req)
    res.status = 201

    return res
end
