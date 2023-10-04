# realm
variable "realm_name" {
  type        = string
  description = "Realm name"
  default     = "demo-realm"
}

variable "realm_display_name" {
  type        = string
  description = "The displayed name"
  default     = "Demo Realm"
}

variable "realm_display_name_html" {
  type        = string
  description = "The displayed name"
  default     = "<b>Demo Realm</b>"
}

# clients
variable "demo_client_name" {
  type        = string
  description = "OpenID Client (name)"
  default     = "demo-client"
}

variable "demo_client_display_name" {
  type        = string
  description = "OpenID Client (name)"
  default     = "Demo Client"
}

variable "demo_client_secret" {
  type        = string
  sensitive   = true
  description = "Demo Client secret"
  default     = "1c4fe703-55f0-485c-a870-8f7480a6c1eb"
}
