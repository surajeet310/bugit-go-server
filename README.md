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
The responses and requests are in `JSON` format . ID's used in the tables are of the data type UUID4 .

#### Non-Auth endpoints (Base url/open)

| Endpoint | Method | Request | Response |
|:--------:|:------:|:-------:|:--------:|
|`/register`|`POST`|`{"fname":"","lname":"","email":"","password":""} `|`{"response":"success","result":null}`|
|`/login`|`POST`|`"email":"","password":""}`|`{"response":"success","result":token}`|

#### Auth endpoints (Base url/auth)
1. ***Workspaces***
	##### Response Data Types
	```json
	{"w_id" : "" , "name" : "" , "project_count" : "" ,"member_count" : ""} - home
	
	{"w_id" : "","name" : "","descp" : "","project_count" : "","member_count" : "","created_at" : ""} - workspace
	
	{"p_id" : "","name" : "","task_count" : "","member_count" : ""} - projects
	
	{"user_id" : "","is_admin" : "","is_taken" : "","user_name" : ""} - workspaceMembers
	
	{"user_id" : "","is_admin" : "","is_taken" : false,"user_name" : ""} - allWorkspaceMembers
	
	```
	##### Request Data Types
	```json
	{"user_id" : "","name" : "","descp" : "", "created_at" : ""} - addWorkspaceStruct
	
	{"w_id" : "","email" : "","requestee_id" : ""} - addWorkspaceMemberReqStruct
	
	{"req_id" : ""} - addWorkspaceMemberStruct
	
	{"w_id" : "", "user_id" : ""} - makeUserAdminStruct
	```

	| Endpoint | Method | Request Body/Query Params | Response |
	|:--------:|:------:|:--------------------:|:--------:|
	|`/home`|`GET`|`user_id`|`{"response":"success","result":[ home ]}`|
	|`/home/workspace`|`GET`|`workspace_id = `|`{"response":"success","result":{ workspace, [ projects] }}`|
	|`/workspaceMembers`|`GET`|`workspace_id = & project_id = `|`{"response":"success","result":[ workspaceMembers ]}`|
	|`/allWorkspaceMembers`|`GET`|`workspace_id = `|`{"response":"success","result":[ allWorkspaceMembers ]}`|
	|`/addWorkspace`|`POST`|`addWorkspaceStruct`|`{"response":"success","result":null}`|
	|`/addWorkspaceMemberReq`|`POST`|`addWorkspaceMemberReqStruct`|`{"response":"success","result":null}`|
	|`/addWorkspaceMember`|`POST`|`addWorkspaceMemberStruct`|`{"response":"success","result":null}`|
	|`/makeUserAdmin`|`POST`|`makeUserAdminStruct`|`{"response":"success","result":null}`|
	|`/deleteWorkspace`|`DELETE`|`workspace_id = `|`{"response":"success","result":null}`|
	|`/removeWorkspaceMember`|`DELETE`|`workspace_id = & user_id = `|`{"response":"success","result":null}`|
