## API Documentations

##### Login Feature

1. **"api/user/login" => Method Post**
    input : formfield
    ```
    username : [username]
    password : [secret]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": [
                {
                "id": 1,
                "username": "yourusername",
                "email": "your.email@domain.com"
                }
            ]
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "Invalid username or password",
            "data": null
        } 
        ```
2. **"api/user/register" => Method Post**
    input : formfield
    ```
    username : [username]
    password : [secret]
    email    : [email@domain.com]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": [
                {
                "id": 1,
                "username": "yourusername",
                "email": "your.email@domain.com"
                }
            ]
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "Please fill all the fields",
            "data": null
        } 
        ```
3. **"api/project/list" => Method Post**
    input : formfield
    ```
    user_id : [user_id]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": [
                {
                "id": 1,
                "name": "project1",
                "user_id": 1,
                "created_at": "2021-12-27T22:08:18Z",
                "updated_at": "2021-12-27T22:08:18Z"
                },
                {
                "id": 3,
                "name": "project2",
                "user_id": 1,
                "created_at": "2021-12-27T22:08:18Z",
                "updated_at": "2021-12-27T22:08:18Z"
                }]
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "No projects found",
            "data": null
        } 
        ```
4. **"api/project/create" => Method Post**
    input : formfield
    ```
    user_id : [user_id]
    project_name: [your_project_name]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": [
                {
                "id": 1,]
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "Please fill all the fields",
            "data": null
        } 
        ```
5. **"api/project/update" => Method Patch**
    input : formfield
    ```
    project_id : [project_id]
    project_name: [new_project's_name]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": [
                {
                "id": 1,
                }]
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "Please fill all the fields",
            "data": null
        } 
        ```
6. **"api/project/delete" => Method Patch**
    input : formfield
    ```
    project_id : [project_id]
    ```
    response : 
     *  Success ✅
        ```{
            "status": 200,
            "message": "Success",
            "data": []
            }
        ```
     *  Fail ❌
        ``` 
        {
            "status": 400,
            "message": "Please fill all the fields",
            "data": null
        } 
        ```