# To-do API

## Authentication

### POST

To get authenticate token we should use this request
```
/login
```
and send Username and Password, like that:
```
{
"Login": "login",
"Password": "password"
}
```

## Account usage

### POST

POST shares task ```/account/task```

POST set photo ```/account/photo```

POST set email ```/account/email```

POST set name ```/account/name```


### DELETE

DELETE delete task ```/account/task```

DELETE delete photo ```/account/photo```

DELETE delete email ```/account/email```

DELETE delete name ```/account/name```


### PATCH



