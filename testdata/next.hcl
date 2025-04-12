server {
    port = 8080
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    if {
        path = "/aaa"

        then {
            status = 302
            body = ""
            headers = {
                "Location": "https://example.com",
            }
        }
    }

    if {
        path = "/bbb"

        not {
            headers = {
                "Accept": "application/json",
            }
        }

        then {
            status = 200 
        }
    }

    # path {
    #     pattern = "/aaa"

    #     status = 302
    #     body = ""
    #     headers = {
    #         "Location": "https://example.com",
    #     }
        
    #     # validate {
    #     #     if {
    #     #     }
    #     #     status = 303
    #     # }
    # }
}
