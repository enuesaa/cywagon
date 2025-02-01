hostname = "example.com"

entry.workdir = "/tmp"
entry.cmd = "go run ."
entry.waitForHealthy = 60

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
