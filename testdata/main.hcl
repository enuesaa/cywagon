server {
    # port = 8080
}

def "status" {
    value = 200
    # notfound = 404
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    path {
        pattern = "/aaa"

        status = def.status
        body = ""
        headers = {
            "Location": "https://example.com",
        }
        
        # validate {
        #     if {
        #     }
        #     status = 303
        # }
    }
}
