# This is the role under which ECS will execute our task. This role becomes more important
# as we add integrations with other AWS services later on.

# The assume_role_policy field works with the following aws_iam_policy_document to allow
# ECS tasks to assume this role we're creating.
resource "aws_iam_role" "toy-robot-task-execution-role" {
  name               = "toy-robot-task-execution-role"
  assume_role_policy = data.aws_iam_policy_document.ecs-task-assume-role.json
}

data "aws_iam_policy_document" "ecs-task-assume-role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

# Normally we'd prefer not to hardcode an ARN in our Terraform, but since this is
# an AWS-managed policy, it's okay.
data "aws_iam_policy" "ecs-task-execution-role" {
  arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# Attach the above policy to the execution role.
resource "aws_iam_role_policy_attachment" "ecs-task-execution-role" {
  role       = aws_iam_role.toy-robot-task-execution-role.name
  policy_arn = data.aws_iam_policy.ecs-task-execution-role.arn
}


resource "aws_ecs_cluster" "toy-robot" {
  name = "toy-robot"
}

resource "aws_cloudwatch_log_group" "toy-robot-service" {
  name = "/ecs/toy-robot-service"
}

resource "aws_ecs_task_definition" "toy-robot-task-definition" {
    family = "toy-robot"

    execution_role_arn = aws_iam_role.toy-robot-task-execution-role.arn
    container_definitions = <<EOF
  [
    {
      "name": "toy-robot-container",
      "image": "nginx:latest",
      "portMappings": [
        {
          "containerPort": 3000
        }
      ],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-region": "ap-southeast-1",
          "awslogs-group": "/ecs/toy-robot",
          "awslogs-stream-prefix": "ecs"
        }
      }
    }
  ]
  EOF

  # These are the minimum values for Fargate containers.
  cpu = 256
  memory = 512
  requires_compatibilities = ["FARGATE"]

  # This is required for Fargate containers (more on this later).
  network_mode = "awsvpc"
}

resource "aws_ecs_service" "toy-robot-ecs" {
  name            = "toy-robot-ecs"
  task_definition = aws_ecs_task_definition.toy-robot-task-definition.arn
  cluster         = aws_ecs_cluster.toy-robot.id
  launch_type     = "FARGATE"
  desired_count = 1

  network_configuration {
    assign_public_ip = false

    security_groups = [
      module.network.aws_security_group_egress_all_id,
      module.network.aws_security_group_api_ingress_id,
    ]
    
    subnets = [
      module.network.aws_subnet_private_id,
    ]
  }

  load_balancer {
    target_group_arn = module.alb.aws_lb_target_group_arn
    container_name   = "toy-robot-container"
    container_port   = "3000"
  }

}