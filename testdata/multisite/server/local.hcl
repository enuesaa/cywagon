site "local" {
  port = 3000
  host = "localhost:3000"
  dist = "../../sampleapp/dist"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    logic = logic.index
  }
}
