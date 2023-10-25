# Server

## Run
1. cd repo/cmd/server
2. go run main.go

### Api
1. **GET** `/tasks`
   - Returns an array of tasks

2. **GET** `/task/{id:[0-9]+}`
   - Returns a specific task

3. **POST** `/task`
   - Create task
   - Request Body (JSON):
     ```json
     {
         "desc": "Task Description",
         "isDone": true/false
     }
     ```

4. **PUT** `/task`
   - Update task
   - Request Body (JSON):
     ```json
     {
         "id": <task_id> (number),
         "desc": "<task desc>" (string),
         "isDone": true/false
     }
     ```

5. **DELETE** `/task/{id:[0-9]+}`
   - Delete task

6. **PUT** `/task/{id:[0-9]+}`
   - Toggle task done status

# CLI

### install
1. cd repo
2. go build -o todo_cli ./cmd/cli 
    this will generate todo_cli executable in root location of repo
3. ./todo_cli -h

### Usages
1. Get All todos
    ```./todo_cli get_all```
2. Get todo by id
    ```./todo_cli -id <todo_id> get```
3. Create todo
    ```./todo_cli -title "your title" create```
4. Mark todo as done
    ```./todo_cli -id <todo_id> mark_done```
5. Delete todo
    ```./todo_cli -id <todo_id> delete```

