logic "index" {
    if {
        path_not_in = ["/**/*.*", "/*.*"]

        rewrite {
            path = "/index.html"
        }
    }
}
