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

## Sign up

### POST

Also to get authenticate token we should use this request


```
/signup
```
and create login and password.

```
{
"Login": "login",
"Password": "password"
}
```
Also data is checked during registration.
```
    login - `^[a-zA-Z0-9!@#$%^&*()_+-=.,:;'"]{6,}$`
    password - `^[a-zA-Z0-9]{3,25}$`
```

## Account usage

### POST

POST shares task ```/dashboard/add```

POST set photo ```/account/photo```

POST set email ```/account/email```

POST set name ```/account/name```


### DELETE

DELETE delete task ```/dashboard/delete```

DELETE delete photo ```/account/photo```

DELETE delete email ```/account/email```

DELETE delete name ```/account/name```


### PATCH



