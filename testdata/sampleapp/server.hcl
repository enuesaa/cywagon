server {
    port = 3000
}

const {
    // test:test
    basicauth = "Basic dGVzdDp0ZXN0"
}

site "sampleapp" {
    host = "localhost:3000"
    dist = "./dist"

    headers = {
        "Aaa": "aaa",
    }

    # if {
    #     path_in = ["/[^.]", "/*/[^.]"]
    #     rewrite {
    #         path = "{path}/index.html"
    #         append = "index.html"
    #         //dist 
    #     }
    # }

    if {
        path = "/restrict/*"
        headers_not = {"Authorization": const.basicauth}

        respond {
            status = 401
            headers = {
                "WWW-Authenticate": "Basic realm=\"Restricted\""
            }
        }
    }

    if {
        path = "/oldpage"

        respond {
            status = 302
            headers = {
                "Location": "/",
            }
        }
    }

    if {
        path_in = ["/aaa", "/aaa/*"]

        respond {
            status = 299
        }
    }
}
