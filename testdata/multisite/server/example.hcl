site "example" {
  port = 3000
  host = "example.local:3000"
  dist = "../../sampleapp/dist"

  tlscert = "./example.local.pem"
  tlskey  = "./example.local-key.pem"

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
