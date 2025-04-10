 **Terraform configuration** that provisions a **Docker container running Nginx**. 
 Here's a breakdown of each block:

---

###  `terraform` block
```hcl
terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}
```
- Tells Terraform to use the **Docker provider** from `kreuzwerker`.
- `~> 3.0.1` means any version `>= 3.0.1` but `< 4.0.0`.

---

###  `provider "docker"` block
```hcl
provider "docker" {}
```
- Configures the Docker provider. Since it's empty, it uses **defaults**, usually local Docker (via Docker daemon on your system).

---

###  `resource "docker_image" "nginx"`
```hcl
resource "docker_image" "nginx" {
  name         = "nginx:latest"
  keep_locally = false
}
```
- Downloads the `nginx:latest` image from Docker Hub.
- `keep_locally = false`: Once the container is removed, the image will also be cleaned up (not retained locally).

---

### `resource "docker_container" "nginx"`
```hcl
resource "docker_container" "nginx" {
  image = docker_image.nginx.image_id
  name  = "tutorial"
  ports {
    internal = 80
    external = 8000
  }
}
```
- Creates a Docker container named `"tutorial"`.
- Uses the previously pulled Nginx image.
- Maps:
  - Container port `80` (default for Nginx)
  - To host machine's port `8000`  
    → So you can visit [http://localhost:8000](http://localhost:8000) and get the Nginx welcome page.

---

### 🧪 Summary:
You're spinning up an **Nginx web server** inside a Docker container using Terraform. You can access it via port `8000` on your host. 
