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

# if {
#     logic = logic.basicauth
# }
