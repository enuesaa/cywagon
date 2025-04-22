server {
    port = 3000
}

const {
    // test:test
    basicauth = "Basic dGVzdDp0ZXN0"
}

# logic "basicauth" {
#     if {
#         path = "/restrict/*"
#         headers_not = {"Authorization": const.basicauth}
#         respond {
#             status = 401
#             headers = {
#                 "WWW-Authenticate": "Basic realm=\"Restricted\""
#             }
#         }
#     }
# }

site "sampleapp" {
    host = "localhost:3000"
    dist = "./dist"

    headers = {
        "Cache-Control": "no-cache",
    }

    if {
        path = "/storage/*"

        rewrite {
            path = "/a.txt"
        }
        respond {
            dist = "../storage"
        }
    }

    # if {
    #     logic = logic.basicauth
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
        path = "/old/*"

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

    if {
        path_not = "/{**/*.*,*.*}"

        rewrite {
            path = "/index.html"
        }
    }
}
