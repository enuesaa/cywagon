host = "example.com"

origin.url = "http://localhost:4173"
origin.workdir = "./testdata/sampleapp"
origin.cmd = "pnpm install && pnpm build && pnpm preview"
origin.waitForHealthy = 5

healthCheck.protocol = "HTTP"
healthCheck.method = "GET"
healthCheck.path = "/"

function handler(next, req)
    res = next(req)
    -- res.status = 201

    return res
end
