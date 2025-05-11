site "sampleapp" {
  port = 3000
  host = "localhost:3000"
  dist = "../dist"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    logic = logic.index
  }
}
