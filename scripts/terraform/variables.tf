
variable "ingress_allow" {
  type = map(any)
  default = {
    elasticsearch = {
      protocol = "tcp",
      ports    = ["9200"]
    }
    ssh = {
      protocol = "tcp"
      ports    = ["22"]
    }
  }
}

variable "egress_allow" {
  type = map(any)
  default = {
    http = {
      protocol = "tcp"
      ports    = ["80", "443"]
    }
  }
}

variable "project" {
  type = string
}
variable "region" {
  type = string
}

variable "es_version" {
  type = string
  default = "7.17.7"
}