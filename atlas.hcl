data "external_schema" "beego" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-beego",
    "load",
    "--path", "./models",
    "--dialect", "postgres"
  ]
}

env "beego" {
  src = data.external_schema.beego.url
  dev ="postgres://postgres:Dev@123@localhost:5432"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}