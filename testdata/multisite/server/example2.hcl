site "example2" {
  port = 443
  host = "example2.local"
  dist = "../../sampleapp/dist"

  tlscert = "./example2.local.pem"
  tlskey  = "./example2.local-key.pem"

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
