# Bugit-Server
Bugit is an android based application, which provides solutions to manage virtual workspaces, projects and assign and track bugs for the same.<br />
The REST Api is written in [Gin](https://github.com/gin-gonic/gin), a framework of [Golang](https://go.dev/) and is deployed on Heroku. The Api uses RDBMS (Postgresql) to organize and store data.

#  API Documentation

### Base Url
https://bugit-server.herokuapp.com/

### Database schema
There exist 9 relational tables as listed below :
* Users 
* Workspaces
* Projects
* Tasks
* Requests
* Workspace-Members
* Project-Members
* Task-members
* Task-comments

### API endpoints

#### Non-Auth endpoints (Base url/open)

| Endpoint | Method | Request | Response |
|:--------:|:------:|:-------:|:--------:|
|`/register`|`POST`|`{"fname":"","lname":"","email":"","password":""}`|`{"response":"","result":null}`|
|`/login`|`POST`|`{"email":"","password":""}`|`{"response":"","result":token}`|

#### Auth endpoints (Base url/auth)
