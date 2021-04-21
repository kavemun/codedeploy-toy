output "toy-robot-codebuild-artifact" {
  value = aws_s3_bucket.artifact_bucket.bucket
}

output "toy-robot-codebuild-arn" {
  value = aws_s3_bucket.artifact_bucket.arn
}