host = "example.com"

entry.workdir = "../kakkofn"
entry.cmd = "pnpm vite preview"
entry.waitForHealthy = 5
entry.host = "http://localhost:4173"

healthCheck.protocol = "HTTP"
healthCheck.method = "GET"
healthCheck.path = "/"

function handler(next, req, res)
    print('handle')
    print(res.status)
    next()

    res.status = 200
    return res
end
