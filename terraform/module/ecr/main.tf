resource "aws_ecr_repository" "toy-robot-ecr" {
  name                 = "toy-robot-ecr"
  image_tag_mutability = "IMMUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}
