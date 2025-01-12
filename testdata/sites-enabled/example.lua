hostname = "example.com"
port = 3000

print('a')

function handle(next, req, res)
    print('handle')
    next()

    res.status = 200
    return res
end
