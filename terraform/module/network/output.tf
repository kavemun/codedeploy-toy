output "aws_subnet_private_id" {
  value = aws_subnet.private.id
}

output "aws_security_group_api_ingress_id" {
  value = aws_security_group.api-ingress.id
}

output "aws_security_group_egress_all_id" {
  value = aws_security_group.egress-all.id
}

output "aws_alb_subnet_public" {
    value = aws_subnet.public
}

output "aws_alb_subnet_private" {
    value = aws_subnet.private
}

output "aws_security_group_http" {
    value = aws_security_group.http
}

output "aws_security_group_https" {
    value = aws_security_group.https
}

output "aws_security_group_egress_all" {
    value = aws_security_group.egress-all
}

output "aws_internet_gateway" {
  value = aws_internet_gateway.igw
}


output "aws_vpc_app" {
  value = aws_vpc.app-vpc
}