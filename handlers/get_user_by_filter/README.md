# UserFilters

## OK

```JSON
{
    "code": 200,
    "status": "OK",
    "body": {
        "id": 1,
        "biotype_id": 1,
        "bone_density": 1,
        "role_id": 1,
        "gender_id": 1,
        "user_uuid": "an-uuid",
        "name": "testing",
        "last_name": "testing",
        "email": "testing@email.com",
        "birthday": "2021-09-11T14:43:05.415894617-05:00",
        "created_at": "2021-09-11T14:43:05.415894888-05:00",
        "updated_at": "2021-09-11T14:43:05.415894938-05:00",
    }
}
```

## Not Found

```JSON
{
    "code": 404,
    "status": "Not Found",
    "body": {
        "error": "Record not found"
    }
}
```

## Internal server error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
        "error": "An ordinary error :O"
    }
}
```

## Validation error

```JSON
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "error": [
            {
                "tag": "user_uuid",
                "validation": "required"
            }
        ]
    }
}
```
