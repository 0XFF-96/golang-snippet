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

    // two way to set it up 
    // env variable 

    // export AWS_ACCESS_KEY_ID="AKIATW2H3VEEU44HXBYE"
    // export AWS_SECRET_ACCESS_KEY="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"

    // ---------------------- 错误写法 -------------------
    // export AWS_SECRERT_ACCESS_KEY=6ElDb9uposrg3cv7vnrN5Q6vyoBnGXZ6b3BmVmc0
    // export AWS_ACCESS_KEY=AKIATW2H3VEEU44HXBYE
    // ---------------------- 错误写法 -------------------- 


    # access_key = "AKIATW2H3VEEU44HXBYE"
    # secret_key = "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"
}

variable "cidr_blocks" {
    description = "cidr block for vpc and subnets"

    // A default value makes the variable optional
    // default = "10.0.10.0/24"

    // 1. possible types
    // 2. "type" specifies what value types are accepted!
    // 3. 
    // type = string 
    // type = list(string)
    
    type = list(
        object(
            {
            cidr_block = string
            name = string
            }
        )
    )
}


variable "vpc_cidr_block" {
    description = "vpc cidr block"
}

variable "environment" {
    description = "deployment environment"
}

// first_name, variable_name
// https://aws.amazon.com/cn/what-is/cidr/
// 
resource "aws_vpc" "development-vpc" {
    // cidr_block = "10.0.0.0/16"
    cidr_block = var.cidr_blocks[0].cidr_block
    tags = {
        Name: var.environment
    }
}

resource "aws_subnet" "dev-subnet-1" {
    // create a resource based on something 
    // that haven't been created yet.
    vpc_id = aws_vpc.development-vpc.id 
    // cidr_block = "10.0.10.0/24"
    cidr_block = var.cidr_blocks[1].cidr_block
    availability_zone = "eu-west-3a"
    tags = {
        Name: var.cidr_blocks[1].name
        vpc_env: "dev"
    }
}

data "aws_vpc" "existing_vpc" {
    // critiera 
    // whenever device 
    default = true 

}

resource "aws_subnet" "dev-subnet-2" {
    // create a resource based on something 
    // that haven't been created yet.
    vpc_id = data.aws_vpc.existing_vpc.id
    cidr_block = "172.31.48.0/20"
    availability_zone = var.avail_zone
       tags = {
        Name: "subnet-1-dev"
    }
}

// remove a resouce 
// 1. remove the resource from terraform file 
// 2. declraimnative remove 
//    terraform destory -target , not consistent , 
// 
//  


// 1. Difference between current and desired state 
// 2. 
// 3. 
// 4. 

// terraform destory, 是一个非常危险⚠️的命令
// you don't have to know in which order you need to delete the resources
// 


// Terraform state ! 
// 1. terraform.tfstate  - 保持本地状态，状态存储的机制
// 2. 


# Usage: terraform [global options] state <subcommand> [options] [args]

#   This command has subcommands for advanced state management.

#   These subcommands can be used to slice and dice the Terraform state.
#   This is sometimes necessary in advanced cases. For your safety, all
#   state management commands that modify the state create a timestamped
#   backup of the state prior to making modifications.

#   The structure and output of the commands is specifically tailored to work
#   well with the common Unix utilities such as grep, awk, etc. We recommend
#   using those tools to perform more advanced state tasks.

# Subcommands:
#     list                List resources in the state
#     mv                  Move an item in the state
#     pull                Pull current state and output to stdout
#     push                Update remote state from a local state file
#     replace-provider    Replace provider in the state
#     rm                  Remove instances from the state
#     show                Show a resource in the state

// output 
// 

output "dev-vpc-id" {
    value = aws_vpc.development-vpc.id 
}

output "dev-subnet-id" {
    value = aws_subnet.dev-subnet-1.id 
}


// Variable 
// There way to assign a value to a variable 
// 1. 
// 2. terrafomr command line - passing parameter - 
// 3. variable file 

// Use case for variable 
// Replicate same infrastrucutre 
// for different environments
// 
// terraform apply -var-file terraform-dev.tfvars


// do NOT hardcode credentials !
// checked in into a Git repository !
// 
// 


// Set variable using TF 
// environment variable 
// 
// export TF_VAR_avail_zone="eu-west-3a"
// 

variable avail_zone {}