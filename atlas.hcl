data "external_schema" "gorm" {
   program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./repository/entity/",
    "--dialect", "postgres",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://hms:hms@localhost:5432/new_drones?search_path=public&sslmode=disable"
  url = "postgres://hms:hms@localhost:5432/new_drones?search_path=public&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "local"{
  url="postgres://hms:hms@localhost:5432/drones?search_path=public&sslmode=disable"
}