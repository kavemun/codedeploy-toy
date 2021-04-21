resource "aws_lb_target_group" "toy-robot-target-group" {
  name        = "toy-robot-target-group"
  port        = 3000
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = module.network.aws_vpc_app.id

  health_check {
    enabled = true
    path    = "/health"
  }

  depends_on = [aws_alb.toy-robot-alb]
}

resource "aws_alb" "toy-robot-alb" {
  name               = "toy-robot-lb"
  internal           = false
  load_balancer_type = "application"

  subnets = [
    module.network.aws_alb_subnet_public.id,
    module.network.aws_alb_subnet_private.id,
  ]

  security_groups = [
    module.network.aws_security_group_http.id,
    module.network.aws_security_group_https.id,
    module.network.aws_security_group_egress_all.id,
  ]

  depends_on = [module.network.aws_internet_gateway]
}

resource "aws_alb_listener" "toy-robot-http" {
  load_balancer_arn = aws_alb.toy-robot-alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.toy-robot-target-group.arn
  }
}

output "alb_url" {
  value = "http://${aws_alb.toy-robot-alb.dns_name}"
}