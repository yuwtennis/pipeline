resource "google_compute_instance" "elastic"  {
  name = "elastic"
  machine_type = "n2-standard-2"
  zone = "asia-northeast1-a"

  metadata = {
    env = "dev"
    app = "elastic"
  }

  boot_disk {}

  attached_disk {
    source = google_compute_disk.elastic-data.self_link
  }

  network_interface {
    subnetwork = "dev-primary"
  }
}

resource "google_compute_disk" "elastic-data" {
  name = "elastic-data"
  type = "pd-ssd"
  zone = "asia-northeast1-a"
  size = "20"
}

# ToDo Firewall