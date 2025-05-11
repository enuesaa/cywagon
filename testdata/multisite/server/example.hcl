site "example" {
  port = 8443
  host = "example.local:8443"
  dist = "../../sampleapp/dist"

  tlscert = "./example.local.pem"
  tlskey  = "./example.local-key.pem"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    logic = logic.index
  }
}
