server {
    port = 8080
}

const {
    basicauth = "Basic xx"
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    headers = {
        "Aaa": "aaa",
    }

    if {
        headers_not = {"Authorization": const.basicauth}

        respond {
            status = 400
            headers = {
                "WWW-Authenticate": "Basic realm=\"Restricted\""
            }
        }
    }

    if {
        path = "/oldpage"

        respond {
            status = 302
            body = ""
            headers = {
                "Location": "https://example.com",
            }
        }
    }
}
