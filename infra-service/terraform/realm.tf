resource "keycloak_realm" "demo_realm" {
  realm             = var.realm_name
  enabled           = true
  display_name      = var.realm_display_name
  display_name_html = var.realm_display_name_html

  login_theme   = "keycloak"
  account_theme = "keycloak.v3"
  admin_theme   = "keycloak.v2"
  email_theme   = "keycloak"

  access_code_lifespan    = "1h"
  access_token_lifespan   = "1h"
  refresh_token_max_reuse = 1 # after using a refresh token, get another one

  ssl_required    = "external"
  password_policy = "upperCase(1) and length(8) and forceExpiredPasswordChange(365) and notUsername"

  attributes = {
    middleName  = "Middle name"
    phoneNumber = "Phone number"
  }

  smtp_server {
    host = "smtp.example.com"
    from = "example@example.com"

    auth {
      username = "tom"
      password = "password"
    }
  }

  internationalization {
    supported_locales = [
      "en",
      "de",
      "es"
    ]
    default_locale = "en"
  }

  security_defenses {
    headers {
      x_frame_options                     = "DENY"
      content_security_policy             = "frame-src 'self'; frame-ancestors 'self'; object-src 'none';"
      content_security_policy_report_only = ""
      x_content_type_options              = "nosniff"
      x_robots_tag                        = "none"
      x_xss_protection                    = "1; mode=block"
      strict_transport_security           = "max-age=31536000; includeSubDomains"
    }
    brute_force_detection {
      permanent_lockout                = false
      max_login_failures               = 30
      wait_increment_seconds           = 60
      quick_login_check_milli_seconds  = 1000
      minimum_quick_login_wait_seconds = 60
      max_failure_wait_seconds         = 900
      failure_reset_time_seconds       = 43200
    }
  }

  web_authn_policy {
    relying_party_entity_name = "Example"
    relying_party_id          = "keycloak.example.com"
    signature_algorithms      = ["ES256", "RS256"]
  }
}
