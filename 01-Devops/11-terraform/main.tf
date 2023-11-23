provider "aws" {
    region = "eu-west-3"
}


terraform {
    required_version = ">= 0.12"
    backend "s3" {
        bucket = "myappbucket96"
        key = "myapp/state.tfstate"
        region = "eu-west-3"
    }
}

module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "my-vpc"
  // cidr = "10.0.0.0/16"
  cidr = var.vpc_cidr_block

  // azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  azs = [var.avail_zone]
  // private_subnets for RDS & DATA bases 
  // private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]

  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]
  public_subnet_tags = {
        Name = "$(var.env_prefix)-subnet-1"
}


#   what is is 
#   enable_nat_gateway = true
#   enable_vpn_gateway = true

  tags = {
    Terraform = "true"
    Environment = "dev"
    Name = "$(var.env_prefix)-vpc"
  }
}


// Module Sources 
// tells TF where to find the source code 
// the source code is downloaded in "terraform init" step 
// 
// Recap our previous AWS resource configuration

// create new Security Group instead of using default ones 
// less code than before 


// Introduction to Remote Storage state file 
// Problem: each user/CI server must make sure they always have 
// the laatest state data before running Terraform 
// How do we share a same Terraform state file ?
// 


resource "aws_subnet" "subnet-2" {
  vpc_id            = module.vpc.vpc_id
  cidr_block        = "10.0.2.0/24"   # 替换为你的 CIDR 块
  availability_zone = "eu-west-3a"    # 替换为你的可用区

  tags = {
    Name = "my-new-subnet"
  }
}

// from: https://github.com/terraform-aws-modules/terraform-aws-alb
module "alb" {
  source = "terraform-aws-modules/alb/aws"

  name    = "my-alb"
  // vpc_id  = "vpc-abcde012"
  vpc_id = module.vpc.vpc_id
  // create = false 

  // At least two subnets in two different Availability Zones must be specified
  subnets = [module.vpc.public_subnets[0], aws_subnet.subnet-2.id]

  # Security Group
  security_group_ingress_rules = {
    all_http = {
      from_port   = 80
      to_port     = 80
      ip_protocol = "tcp"
      description = "HTTP web traffic"
      cidr_ipv4   = "0.0.0.0/0"
    }
    # all_https = {
    #   from_port   = 443
    #   to_port     = 443
    #   ip_protocol = "tcp"
    #   description = "HTTPS web traffic"
    #   cidr_ipv4   = "0.0.0.0/0"
    # }
  }
  security_group_egress_rules = {
    all = {
      ip_protocol = "-1"
      cidr_ipv4   = "10.0.0.0/16"
    }
  }

  # access_logs = {
  #   bucket = "my-alb-logs"
  # }

  listeners = {
    ex-http-https-redirect = {
      port     = 80
      protocol = "HTTP"
      redirect = {
        port        = "443"
        protocol    = "HTTPS"
        status_code = "HTTP_301"
      }
    }
  }


// 一直写✍️不对这里的规则，😫
    # // 这个规则有问题
  #   ex-https = {
  #     port            = 443
  #     protocol        = "HTTPS"
  #     // certificate_arn = "arn:aws:iam::123456789012:server-certificate/test_cert-123456789012"

  #     forward = {
  #       // target_group_key = "ex-instance"
  #       // target_group_key = "main"
  #       target_group_arn = "arn:aws:elasticloadbalancing:eu-west-3:255164393737:targetgroup/main-tg/18ddccdff0df3d73"
  #     }
  #   }
  # }

  # 不直接在这里创建 target_groups 
  # target_groups = {
  #   ex-instance = {
  #     name_prefix      = "h1"
  #     protocol         = "HTTP"
  #     port             = 80
  #     target_type      = "instance"
  #   }
  # }

#   target_groups = {
#     main = {
#     name_prefix      = "h1"
#     protocol         = aws_lb_target_group.main.protocol
#     port             = aws_lb_target_group.main.port
#     target_type      = "instance"
#   }
# }

  tags = {
    Environment = "Development"
    Project     = "Example"
  }
}

resource "aws_lb_target_group" "main" {
  name     = "main-tg"
  port     = 80
  protocol = "HTTP"
  vpc_id   = module.vpc.vpc_id
}

