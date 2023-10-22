# Apis

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
