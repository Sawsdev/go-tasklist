# Task Tracker / TODO list
---

Task tracker project created for [roadmap.sh](https://roadmap.sh/) project [task tracker](https://roadmap.sh/projects/task-tracker) the details of the project can be found there

---

## How to run the project
Clone the project and move into the project folder
~~~
https://github.com/Sawsdev/go-tasklist.git
cd go-tasklist
~~~

Build the binary and start the program
~~~
go build -o tasktracker
./tasktracker
~~~

the program will start in the console and you can interact with using the following commands

### add :
Add a task to the list 
```
add "Make coffee"

```

---
### delete :
Delete a task from the list 
```
delete 1
```

---
### list :
List all tasks
```
list

# You can filter the task by status
list todo
list in-progress
list done
```

---
### update :
Update a task 
```
update 1 "Make coffee with milk"
```

---
### mark-done :
Mark a task as done 
```
mark-done 1
```

---
### mark-in-progress :
Mark a task as in progress 
```
mark-progress 1
```

---
### terminate :
Terminate the program: 
```
terminate
```

---
### commands :
List all available commands
```
commands
```


