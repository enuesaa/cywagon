host = "example.com"

entry.workdir = "../kakkofn"
entry.cmd = "pnpm vite preview"
entry.waitForHealthy = 5
entry.host = "http://localhost:4173"

healthCheck.protocol = "HTTP"
healthCheck.method = "GET"
healthCheck.path = "/"

function handler(req)
    print('handle')

    res = req.invoke()
    res.status = 200

    return res
end
