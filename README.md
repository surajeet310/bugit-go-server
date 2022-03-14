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
The responses and requests are in `JSON` format . Unique identifiers, in all the tables are of the data type UUID4 . `JWT` (JSON Web Token) is used for authentication .

##### Data Types of columns
	(project_count, member_count, priority, alert) - int
	(name,descp, deadline, created_at, user_name, requestee, email, tech, comment) - string
	(is_taken, is_admin, is_assignee) - bool
	(assignee, assigned_to) - uuid

#### Non-Auth endpoints (Base url/open)

| Endpoint | Method | Request | Response |
|:--------:|:------:|:-------:|:--------:|
|`/register`|`POST`|`{"fname":"","lname":"","email":"","password":""} `|`{"response":"success","result":null}`|
|`/login`|`POST`|`"email":"","password":""}`|`{"response":"success","result":token}`|

#### Auth endpoints (Base url/auth)
1. ***Workspaces***
	##### Response Data Types
	```json
	{"w_id" :  , "name" : "" , "project_count" :  ,"member_count" : } - home
	
	{"w_id" : ,"name" : "","descp" : "","project_count" : ,"member_count" : ,"created_at" : ""} - workspace
	
	{"p_id" : ,"name" : "","task_count" : ,"member_count" : } - project
	
	{"user_id" : ,"is_admin" : ,"is_taken" : ,"user_name" : ""} - workspaceMembers
	
	{"user_id" : ,"is_admin" : ,"is_taken" : false, "user_name" : ""} - allWorkspaceMembers
	
	{"req_id" : ,"w_id" : ,"user_id" : , "requestee" : "", "priority" : } - request
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
	|`/home`|`GET`|`user_id = `|`{"response" : "success","result" : [ home ]}`|
	|`/home/workspace`|`GET`|`workspace_id = `|`{"response" : "success","result" : { workspace, [ project ] }}`|
	|`/workspaceMembers`|`GET`|`workspace_id = & project_id = `|`{"response" : "success","result" : [ workspaceMembers ]}`|
	|`/allWorkspaceMembers`|`GET`|`workspace_id = `|`{"response" : "success","result" : [ allWorkspaceMembers ]}`|
	|`/requests`|`GET`|`user_id = `|`{"response" : "success","result" : [ request ]}`|
	|`/addWorkspace`|`POST`|`addWorkspaceStruct`|`{"response" : "success","result" : null}`|
	|`/addWorkspaceMemberReq`|`POST`|`addWorkspaceMemberReqStruct`|`{"response" : "success","result" : null}`|
	|`/addWorkspaceMember`|`POST`|`addWorkspaceMemberStruct`|`{"response" : "success","result" : null}`|
	|`/makeUserAdmin`|`POST`|`makeUserAdminStruct`|`{"response" : "success","result" : null}`|
	|`/deleteWorkspace`|`DELETE`|`workspace_id = `|`{"response" : "success","result" : null}`|
	|`/removeWorkspaceMember`|`DELETE`|`workspace_id = & user_id = `|`{"response" : "success","result" : null}`|

2. ***Projects***
	##### Response Data Types
	```json
	{"p_id" : ,"name" : "", "descp" : "", "task_count" : ,"member_count" : ,"created_at" : "", "deadline" : "", "tech" : ""} - project
	
	{"t_id" : , "name" : ""} - task
	
	{"user_id" : ,"is_admin" : ,"is_assigned" : ,"user_name" : ""} - projectMembers
	
	{"user_id" : ,"is_admin" : ,"is_assigned" false: ,"user_name" : ""} - allProjectMembers
	```
	
	##### Request Data Types
	```json
	{"w_id" : "", "user_id" : "", "name" : "", "descp" : "", "created_at" : "", "tech" : "", "deadline" : "" } - addProjectStruct
	
	{"p_id" : "", "user_id" : ""} - addProjectMemberStruct
	
	{ "p_id" : "", "user_id" : ""} - makeProjectMemberAdminStruct
	```
	
	| Endpoint | Method | Request Body/Query Params | Response |
	|:--------:|:------:|:--------------------:|:--------:|
	|`/project`|`GET`|`project_id = `|`{"response" : "success","result" : { project, [ task ] } }`|
	|`/projectMembers`|`GET`|`project_id = & workspace_id = `|`{"response" : "success","result" : [projectMembers] }`|
	|`/allProjectMembers`|`GET`|`project_id = `|`{"response" : "success","result" : [allProjectMembers] }`|
	|`/addProject`|`POST`|`addProjectStruct`|`{"response" : "success","result" : null }`|
	|`/addProjectMember`|`POST`|`addProjectMemberStruct`|`{"response" : "success","result" : null }`|
	|`/makeProjectUserAdmin`|`POST`|`makeProjectMemberAdminStruct`|`{"response" : "success","result" : null }`|
	|`/deleteProject`|`DELETE`|`project_id = `|`{"response" : "success","result" : null }`|
	|`/removeProjectMember`|`DELETE`|`user_id = & project_id = `|`{"response" : "success","result" : null }`|

3. ***Tasks***
	##### Response Data Types
	```json
	{"t_id" : , "p_id" : ,"name" : "", "descp" : "", "assignee" : ,"assigned_to" : ,"created_at" : "", "deadline" : "", "tech" : ""} - task
	
	{"tc_id" : , "t_id" : , "user_id" : , "comment" : "", "created_at" : } - taskComment
	```
	
	##### Request Data Types
	```json
	{"p_id" : "", "name" : "", "descp" : "", "assignee" : "", "created_at" : "", "deadline" : "", "tech" : "" } - addTaskStruct
	
	{"t_id" : "", "user_id" : "", "comment" : "", "created_at" : "" } - addCommentStruct
	
	{"t_id" : "", "user_id" : "" } - assignTaskComment
	```
	
	| Endpoint | Method | Request Body/Query Params | Response |
	|:--------:|:------:|:--------------------:|:--------:|
	|`/task`|`GET`|`task_id = `|`{"response" : "success","result" : { task, [ taskComment ] } }`|
	|`/addTask`|`POST`|`addTaskStruct`|`{"response" : "success","result" : null }`|
	|`/addComent`|`POST`|`addCommentStruct`|`{"response" : "success","result" : null }`|
	|`/assignTask`|`POST`|`assignTaskComment`|`{"response" : "success","result" : null }`|
	|`/deleteTask`|`DELETE`|`task_id = `|`{"response" : "success","result" : null }`|

4. ***Users***
	##### Response Data Types
	```json
	{"email" : "", "fname" : "", "lname" : "", "alert" : } - user
	```
	
	##### Request Data Types
	```json
	{"user_id" : "", "password" : "" } - checkPwdStruct
	
	{"user_id" : "", "fname" : "" } - changeFnameStruct
	
	{"user_id" : "", "lname" : "" } - changeLnameStruct
	
	{"user_id" : "", "password" : "" } - changePasswStruct
	```
	
	| Endpoint | Method | Request Body/Query Params | Response |
	|:--------:|:------:|:--------------------:|:--------:|
	|`/user`|`GET`|`user_id = `|`{"response" : "success","result" : user }`|
	|`/getUserId`|`GET`||`{"response" : "success","result" : user_id }`|
	|`/checkPwd`|`POST`|`checkPwdStruct`|`{"response" : "success","result" : null }`|
	|`/changeFname`|`PATCH`|`changeFnameStruct`|`{"response" : "success","result" : null }`|
	|`/changelname`|`PATCH`|`changeLnameStruct`|`{"response" : "success","result" : null }`|
	|`/changePwd`|`PATCH`|`changePasswStruct`|`{"response" : "success","result" : null }`|
	|`/deleteUser`|`DELETE`|`user_id = `|`{"response" : "success","result" : null }`|
