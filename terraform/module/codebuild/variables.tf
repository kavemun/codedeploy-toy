# Source repo name and branch

variable "family" {
  description = "Family of the Task Definition"
  default     = "toy-robot"
}

variable "aws_region" {
  description = "The AWS region to create things in."
  default     = "ap-southeast-1"
}