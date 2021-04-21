terraform {
  required_version = ">= 0.12"

  backend "s3" {
    bucket  = "toy-robot-tf-bucket"
    key     = "terraform.tfstate"
    region  = "ap-southeast-1"
    profile = "default"
  }
}
