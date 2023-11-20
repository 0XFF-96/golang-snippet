# Terraform 

Class From Udemy: https://www.udemy.com/course/complete-terraform-course-beginner-to-advanced/



1. https://www.udemy.com/course/complete-terraform-course-beginner-to-advanced/



2. sudo yum update (有问题)

    ```
        Could not retrieve mirrorlist https://amazonlinux-2-repos-eu-west-3.s3.dualstack.eu-west-3.amazonaws.com/2/core/latest/x86_64/mirror.list error was


        12: Timeout on https://amazonlinux-2-repos-eu-west-3.s3.dualstack.eu-west-3.amazonaws.com/2/core/latest/x86_64/mirror.list: (28, 'Failed to connect to amazonlinux-2-repos-eu-west-3.s3.dualstack.eu-west-3.amazonaws.com port 443 after 5000 ms: Timeout was reached')


        One of the configured repositories failed (未知),
        and yum doesn't have enough cached data to continue. At this point the only
        safe thing yum can do is fail. There are a few ways to work "fix" this:

        1. Contact the upstream for the repository and get them to fix the problem.

        2. Reconfigure the baseurl/etc. for the repository, to point to a working
        upstream. This is most often useful if you are using a newer
        distribution release than is supported by the repository (and the
        packages for the previous distribution release still work).

        3. Run the command with the repository temporarily disabled
            yum --disablerepo=<repoid> ...

        4. Disable the repository permanently, so yum won't use it by default. Yum
            will then just ignore the repository until you permanently enable it
            again or use --enablerepo for temporary usage:

            yum-config-manager --disable <repoid>
        or
            subscription-manager repos --disable=<repoid>

        5. Configure the failing repository to be skipped, if it is unavailable.
        Note that yum will try to contact the repo. when it runs most commands,
        so will have to try and fail each time (and thus. yum will be be much
        slower). If it is a very temporary problem though, this is often a nice
        compromise:

            yum-config-manager --save --setopt=<repoid>.skip_if_unavailable=true

        Cannot find a valid baseurl for repo: amzn2-core/2/x86_64
    
    ```

3. 卡住在 EC2 不能连接外网的问题上了。 
    1. https://repost.aws/knowledge-center/ec2-connect-internet-gateway
    2. Why can't my Amazon EC2 instance connect to the internet using an internet gateway?

4. 完成 Docker 镜像的相关部署

5. provisioner, 这些命令的执行总是有问题。 在这个方面上，不如 ansible 好用 

6. 学习如何将项目模块化 （ 重新复习 AWS 的工程师证书🧑‍💻）
    6.1 Grup multiple resources into a logical unit 

    6.2 init moudle 
        Initializing the backend...
        Initializing modules...

        Initializing provider plugins...
        - Reusing previous version of hashicorp/aws from the dependency lock file
        - Using previously-installed hashicorp/aws v5.26.0

        Terraform has been successfully initialized!

        You may now begin working with Terraform. Try running "terraform plan" to see
        any changes that are required for your infrastructure. All Terraform commands
        should now work.

        If you ever set or change modules or backend configuration for Terraform,
        rerun this command to reinitialize your working directory. If you forget, other
        commands will detect it and remind you to do so if necessary.

    6.3 How do we access the resources of a child moudle ?

    6.4 Our "webserver" module 
    6.5 We learnt 
        1. how to write a module 
        2. how to use/reference a module


    7.1   Configurate Remote Storage with 
        Do you want to copy existing state to the new backend?
        Pre-existing state was found while migrating the previous "local" backend to the
        newly configured "s3" backend. No existing state was found in the newly
        configured "s3" backend. Do you want to copy this state to the new "s3"
        backend? Enter "yes" to copy and "no" to start with an empty state.

        Enter a value: yes    


### 资源

1. https://www.pulumi.com/ pulumi 之间有什么关系
2. 课程资源，https://gitlab.com/nanuchi/udemy-terraform-learn
3. 课程资源在不通的分支上，https://gitlab.com/nanuchi/terraform-learn/-/blob/feature/deploy-to-ec2-default-components/main.tf?ref_type=heads 。 
4. 