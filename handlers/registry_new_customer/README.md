# Register New User

## Request

```JSON
{
    "name": "testing",
    "last_name": "main",
    "email": "testing@main-func.com",
    "birthday": "2021-09-27T19:38:57.731173821-05:00",
    "role_id": 3,
    "gender_id": 1
}
```

## Success

```JSON
{
    "code": 201,
    "status": "Created",
    "body": {
        "user_id": 1
    }
}
```

## Unexpected Error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "an error"
    }
}
```

## Validation Error

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "name",
                "validation": "required"
            }
        ]
    }
}
```
