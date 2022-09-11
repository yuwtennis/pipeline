
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

variable "project" {
  type = string
}
variable "region" {
  type = string
}