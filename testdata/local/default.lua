host = "example.com"

origin.url = "http://localhost:4173"
origin.workdir = "../kakkofn"
origin.cmd = "pnpm vite preview"
origin.waitForHealthy = 5

healthCheck.protocol = "HTTP"
healthCheck.method = "GET"
healthCheck.path = "/"

function handler(next, req)
    res = next(req)
    -- res.status = 201

    return res
end
