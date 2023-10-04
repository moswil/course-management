terraform {
  required_providers {
    keycloak = {
      source  = "mrparkers/keycloak"
      version = "4.3.1"
    }
  }
}

provider "keycloak" {
  client_id     = "shirwil-dev"
  client_secret = "6AWk2FYvyYFNAfvaqqnY00WeWgGk5ygK"
  url           = "http://dev.keycloak.shirwil.com:8180"
}
