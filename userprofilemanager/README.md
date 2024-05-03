# User Profile services 

This is a project built with the following requirements:
- Create new user profile
- Check login and return userid
- Update user profile
    - Get user profile


## Getting started
To get started with this project, follow the steps below.

## Requirements and configuration
- Golang 1.21
- Docker
- Docker-compose
- MongoDB 


## Clone the repository

```
cd your_folder
git clone https://gitlab.vnpaycloud.vn/cloud/training/by-user/huyng2/userprofilemanager.git
git checkout dev
```



## Run with docker compose
```
docker-compose up -d
```
project will run on http://localhost:50051


## Run with golang

```
go mod tidy
go run main.go

```
