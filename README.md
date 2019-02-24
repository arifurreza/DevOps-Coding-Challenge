# Web Application with Round Robin Load Balancing & Mongo DB

1. [Goal](#Goal)
2. [Design](#architecture)
3. [Installation](#installation)
4. [Deployment](#deployment)
5. [Delete](#delete)


## Goal
  - Design a simple Web Application - 2 Application servers(here GoLang) and 1 Web Server as a Load Balancer(NGINX).
  - Sending a HTTP request to the Web Server should return the below response in a round-robin load balancing mechanism:  
  `Hello FASHION CLOUD , I'm served from <application node hostname>!`
  - Following architecture has been implemented in AWS platform by CloudFormation service.

## Design
Below figure represents the high level architecture[1]

![Architectural Diagram](https://github.com/arifurreza/DevOps-Coding-Challenge/blob/master/architecture.png)



Specs:
- Instance Type: t2.micro
- AMI: Amazon Linux AMI
- Application Server - GoLang
- Web Server - NGINX
- Database - Mongo

## Installation
1. Should have admin access to AWS Console
2. Create a common key pair i.e., pem file (here ec2-keypair.pem) for accessing both Application Servers and Web Server over SSH.
3. As I will be deploying CloudFormation template via AWS CLI, I will be creating user with programmatic admin access to AWS services. If you are deploying CloudFormation Templates via AWS console, please skip this step.
4. [Install](https://docs.aws.amazon.com/cli/latest/userguide/awscli-install-windows.html) and [Configure](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html) AWS CLI in your local machine.

## Deployment

Clone the following git repo to your local machine.

`git clone https://github.com/arifurreza/DevOps-Coding-Challenge.git`

Execute below command via AWS CLI inside your locally cloned git repository to create a simple Web Application using AWS CloudFormation.

`aws cloudformation create-stack --stack-name web_app --template-body file://ec2_launch.yml --profile demo ap-south-1`

Things achieved by executing above commands are:
- Procures 5 EC2 Instances, Installs GoLang and runs the code over port 80 on either of the servers. Thus we have successfully procured and deployed our Application Servers.
- Procures a single EC2 Instances, Installs NGINX and runs the webserver over port 80. Thus we have successfully procured and deployed our NGINX webserver.
- Procures 3 EC2 Instances, Installs MongoDB and runs the MongoDB server with 27017 port and create admin database as well as admin user and some necessary configuration

Finally, configure your NGINX Web Server to listen request over 2 Application Server via round-robin load balancing mechanism.
Login to your NGINX Web Server via SSH and edit `/etc/nginx/conf.d/load_blancer.conf`

```
upstream backend {
   server app_server1_IP;
   server app_server2_IP;
}

server {
   listen 80;
   server_name webserver_IP;

   location / {
      proxy_pass http://backend;
   }
}
```

Replace `app_server1_IP` by Application Server 1 IP, `app_server2_IP` by Application Server 2 IP and `webserver_IP` by NGINX Web Server IP.

Restart your nginx web server for your changes to reflect.
`service nginx restart`

## Delete

You can delete the stack in by executing the below command

`aws cloudformation delete-stack --stack-name web_app`

