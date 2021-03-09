# esi-hw1


### Task1
```bash
> *Create a simple backend service in Golang that exposes a REST interface for a TODO app with the following functionality:
> * [x] List ToDos
> * [x] Add ToDo
> * [x] Remove ToDo
> * [x] Mark ToDo as completed
> * [x] Add a crazy, artsy, or funny feature that makes your TODO app unique 
> * (e.g. colorization or importance levels, creation date, emoji counters, ... anything else not too complex is welcome here) - please document what you picked!
 
```

### CMD Run application
```bash
go run . 
```
### Curl manual testing
```bash
# get homepage
curl http://localhost:8000

# list all ToDos
curl http://localhost:8000/todos

# get single ToDo (here 1)
curl http://localhost:8000/todo/1

# add new ToDo
curl -X POST -d \
'{"Title":"Swimming","Description":"Training","Date":"2021-03-09","Time":"10:00 pm","HighPriority":false,"Completed":true}' \
http://localhost:8000/todo

# delete ToDo with index 1
curl -X DELETE http://localhost:8000/todo/1

# Mark ToDo as completed with index 1
curl -X PUT http://localhost:8000/todo/1

# Unique feature update Todo to have high priority with index 1
curl -X PUT http://localhost:8000/todos/1
```
