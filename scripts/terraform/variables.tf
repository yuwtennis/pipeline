
variable "ingress_allow" {
  type = map(any)
  default = {
    elasticsearch = {
      protocol = "tcp",
      ports    = ["9200"]
    }
  }
}