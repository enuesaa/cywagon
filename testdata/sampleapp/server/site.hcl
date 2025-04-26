site "sampleapp" {
    host = "localhost:3000"
    dist = "../dist"

    headers = {
        "Cache-Control": "no-cache",
    }

    if {
        path = "/storage/*"

        rewrite {
            from_path_pattern = "/storage(/*)"
            path = "%1/index.html"
        }
        respond {
            dist = "../../storage"
        }
    }

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
        path = "/old"

        respond {
            status = 301
            headers = {
                "Location": "/",
            }
        }
    }

    if {
        path_not = "/{**/*.*,*.*}"

        rewrite {
            path = "/index.html"
        }
    }
}
