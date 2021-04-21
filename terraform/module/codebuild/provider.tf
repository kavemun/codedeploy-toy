module "ecr" {
  source = "../ecr"
}

module "code-commit" {
  source = "../code-commit"
}

module "s3" {
  source = "../s3"
}