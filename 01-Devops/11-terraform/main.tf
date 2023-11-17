# terraform {
#   required_providers {
#     aws = {
#       source  = "hashicorp/aws"
#       version = "~> 5.0"
#     }
#   }
# }

provider "aws" {
    region = "eu-west-3"
}

variable vpc_cidr_block {}
variable avail_zone {}
variable subnet_cidr_block {}
variable env_prefix {}

resource "aws_vpc" "myapp-vpc" {
    // cidr_block = "10.0.0.0/16"
    cidr_block = var.vpc_cidr_block
    tags = {
        Name: "${var.env_prefix}-vpc"
    }
}

resource "aws_subnet" "myapp-subnet-1" {
    // create a resource based on something 
    // that haven't been created yet.
    vpc_id = aws_vpc.myapp-vpc.id 
    // cidr_block = "10.0.10.0/24"
    cidr_block = var.subnet_cidr_block
    availability_zone = var.avail_zone
    tags = {
        Name: "${var.env_prefix}-subnet-1"
    }
}

// 1. 
// Terraform knows 
// in which sequence the components must be created 
// 2. 
// 3. 
// 4.
resource "aws_internet_gateway" "myapp-igw" {
    vpc_id = aws_vpc.myapp-vpc.id 
}


resource "aws_route_table" "myapp-route-table" {
    vpc_id = aws_vpc.myapp-vpc.id 

    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.myapp-igw.id 
    }

    tags = {
        Name "${var.env_prefix}-igw"
    }
}
