site "sampleapp" {
  port = 3000
  host = "localhost:3000"
  dist = "../dist"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    path = "/storage/*"

    rewrite {
      path = "/{dir2:}" # Example: "/{dir2}", "/{:dir2}", "/{last}", "{path}"
    }
    respond {
      dist = "../../storage"
    }
  }

  if {
    logic = logic.basicauth
  }

  if {
    path = "/old"

    respond {
      status = 301
      headers = {
        "Location" : "/",
      }
    }
  }

  if {
    logic = logic.index
  }
}
