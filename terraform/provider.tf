provider "aws" {
  region  = var.region
  shared_credentials_file = "~/.aws/credentials"
  profile = "default"
}

module "sqs" {
  source = "./module/sqs"
}

module "ecs" {
  source = "./module/ecs"
}

module "codebuild" {
  source = "./module/codebuild"
}

module "common" {
  source = "./module/common"
}
