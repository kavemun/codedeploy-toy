output "toy_robot_repo_name" {
  value = aws_ecr_repository.toy-robot-ecr.name
}

output "toy_robot_repo_arn" {
  value = aws_ecr_repository.toy-robot-ecr.arn
}

output "toy_robot_repo_url" {
  value = aws_ecr_repository.toy-robot-ecr.repository_url
}