
packages {
  development = ["buf@1.29.0", "go@1.21.6", "gotools@0.7.0", "delve@1.21.2", "api-linter@1.63.2"]
  runtime     = ["cacert@3.95"]
}

gomodule {
  name       = "bsf-apis"
  src        = "../."
  doCheck    = false
}
