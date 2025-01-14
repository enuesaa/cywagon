hostname = "example.com"
port = 3000

print('a')

function handler(next, req, res)
    print('handle')
    print(res.status)
    next()

    res.status = 200
    return res
end
