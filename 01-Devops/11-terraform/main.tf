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

  //azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
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

