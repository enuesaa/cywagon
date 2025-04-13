server {
    port = 8080
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    headers = {
        "Aaa": "aaa",
    }

    if {
        headers_not = {"Authorization": "Basic xx"}        

        respond {
            status = 400
            headers = {
                "WWW-Authenticate": "Basic realm=\"Restricted\""
            }
        }
    }

    if {        
        ipaddr_not_in = [""]

        respond {
            status = 403
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
