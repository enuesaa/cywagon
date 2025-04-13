server {
    # port = 8080
}

const {
    success = 200
}
const {
    notfound = 404
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    path {
        # each = ["/aaa", "/bbb"]
        pattern = "/aaa"

        status = const.success
        body = ""
        headers = {
            "Location": "https://example.com",
        }
    }
}
