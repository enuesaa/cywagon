server {
    port = 3000
}

const {
    basicauth = "Basic xx"
}

site "sampleapp" {
    host = "localhost:3000"
    dist = "./testdata/sampleapp/dist"

    headers = {
        "Aaa": "aaa",
    }

    if {
        path = "/restrict"
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

    if {
        path = "/aaa"

        respond {
            status = 299
        }
    }
}
