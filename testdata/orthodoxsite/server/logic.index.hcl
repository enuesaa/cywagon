logic "index" {
  if {
    path = "/"

    rewrite {
      path = "/index.html"
    }
  }

  if {
    path = "/**/*"
    path_not_in = ["/**/*.*", "/*.*"]

    rewrite {
      path = "{path}.html"
    }
  }
}
