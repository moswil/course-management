# demo client (confidential type)
resource "keycloak_openid_client" "demo_openid_client" {
  realm_id      = keycloak_realm.demo_realm.id
  client_id     = var.demo_client_name
  client_secret = var.demo_client_secret

  name    = var.demo_client_display_name
  enabled = true

  # extract the rest to variables
  access_type                  = "CONFIDENTIAL"
  standard_flow_enabled        = true
  direct_access_grants_enabled = true
  pkce_code_challenge_method   = "S256"
  full_scope_allowed           = true

  valid_redirect_uris = [
    "http://localhost:8080/oauth2/callback"
  ]

  login_theme = "keycloak"

  extra_config = {
    "key1" = "value1"
    "key2" = "value2"
  }
}

resource "keycloak_openid_audience_protocol_mapper" "demo_client_aud_mapper" {
  realm_id  = keycloak_realm.demo_realm.id
  client_id = keycloak_openid_client.demo_openid_client.id
  name      = var.demo_client_name

  included_custom_audience = var.demo_client_name
}

# roles associated with our demo_openid_client
resource "keycloak_role" "demo_instructor_role" {
  realm_id    = keycloak_realm.demo_realm.id
  client_id   = keycloak_openid_client.demo_openid_client.id
  name        = "instructor"
  description = "Instructor role"
}

resource "keycloak_role" "demo_student_role" {
  realm_id    = keycloak_realm.demo_realm.id
  client_id   = keycloak_openid_client.demo_openid_client.id
  name        = "student"
  description = "Student role"
}

resource "keycloak_role" "admin_role" {
  realm_id    = keycloak_realm.demo_realm.id
  client_id   = keycloak_openid_client.demo_openid_client.id
  name        = "course-admin"
  description = "Admin Role (Instructor and Student)"
  composite_roles = [
    keycloak_role.demo_instructor_role.id,
    keycloak_role.demo_student_role.id
  ]
}

# client with sa enabled
resource "keycloak_openid_client" "demo_sa_openid_client" {
  realm_id      = keycloak_realm.demo_realm.id
  client_id     = "demo-sa-client"
  client_secret = var.demo_client_secret

  name                     = "Demo SA Client"
  enabled                  = true
  access_type              = "CONFIDENTIAL"
  service_accounts_enabled = true
}

# data sources
data "keycloak_openid_client" "realm_management" {
  realm_id  = keycloak_realm.demo_realm.id
  client_id = "realm-management"
}

# use the data source
data "keycloak_role" "manage_users" {
  realm_id  = keycloak_realm.demo_realm.id
  client_id = data.keycloak_openid_client.realm_management.id
  name      = "manage-users"
}

resource "keycloak_openid_client_service_account_role" "demo_sa_manage_users" {
  realm_id                = keycloak_realm.demo_realm.id
  service_account_user_id = keycloak_openid_client.demo_sa_openid_client.service_account_user_id
  client_id               = data.keycloak_openid_client.realm_management.id
  role                    = data.keycloak_role.manage_users.name
}
