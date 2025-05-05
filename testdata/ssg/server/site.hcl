site "sampleapp" {
  host = "localhost:3000"
  dist = "../dist"

  headers = {
    "Cache-Control" : "no-cache",
  }

  if {
    logic = logic.index
  }
}
