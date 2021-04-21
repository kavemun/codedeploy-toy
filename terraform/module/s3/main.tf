resource "aws_s3_bucket" "input_file_bucket" {
    bucket = "${module.common.account_id}-toy-robot-input"
    acl = "private"
}

resource "aws_s3_bucket" "artifact_bucket" {
    bucket = "${module.common.account_id}-toy-robot-codebuiild-artifact"
    acl = "private"
}