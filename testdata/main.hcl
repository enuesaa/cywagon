server {
    # port = 8080
}

const "status" {
    success = 200
    notfound = 404
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    path {
        each = ["/aaa", "/bbb"]
        pattern = each.value

        status = const.status.success
        body = ""
        headers = {
            "Location": "https://example.com",
        }
    }
}
