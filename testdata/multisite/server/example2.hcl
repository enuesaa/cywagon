site "example2" {
  port = 8443
  host = "example2.local:8443"
  dist = "../../sampleapp/dist"

  tlscert = "./example2.local.pem"
  tlskey  = "./example2.local-key.pem"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    logic = logic.index
  }
}
