resource "google_compute_instance" "elastic" {
  name         = "elastic"
  machine_type = "n2-standard-2"
  zone         = "asia-northeast1-a"

  metadata = {
    env = "dev"
    app = "elastic"
  }

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  attached_disk {
    source = google_compute_disk.elastic_data.self_link
  }

  network_interface {
    subnetwork = google_compute_subnetwork.primary.name
  }

  tags = ["elastic-fw"]
}

resource "google_compute_disk" "elastic_data" {
  name = "elastic-data"
  type = "pd-ssd"
  zone = "asia-northeast1-a"
  size = "20"
}