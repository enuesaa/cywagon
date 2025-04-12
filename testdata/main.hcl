server {
    port = 8080
}

site "sampleapp" {
    host = "sample.example.com"
    dist = "./sampleapp/dist"

    path {
        pattern = "/aaa"

        status = 302
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

    # path {
    #     pattern = "/**"

    # }
}

# site "sampleapp" {
#     host = "sample.example.com"
# }
