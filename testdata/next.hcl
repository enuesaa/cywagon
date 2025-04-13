server {
    port = 8080
}

site "sampleapp" {
    host = "sample.example.com"

    if {
        headers = {"Authorization": "Basic xx"}        

        then {
            status = 400
            headers = {
                "WWW-Authenticate": "Basic realm=\"Restricted\""
            }
        }
    }

    if {
        ipaddr = not(in(["", ""]))

        then {
            status = 403
        }
    }

    if {
        path = "/oldpage"

        then {
            status = 302
            body = ""
            headers = {
                "Location": "https://example.com",
            }
        }
    }

    dist = "./sampleapp/dist"
    # status = 302
    # headers = {
    #     "Location": "https://example.com",
    # }
}
