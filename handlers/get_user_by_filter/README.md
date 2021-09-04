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
        "gender_id": 1,
        "name": "testing",
        "last_name": "testing",
        "email": "email@email.com",
        "birthday": "2021-09-02T19:52:00.618373968-05:00",
        "created_at": "2021-09-02T19:52:00.618374028-05:00",
        "updated_at": "0001-01-01T00:00:00Z"
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
                "tag": "user_id",
                "validation": "required"
            }
        ]
    }
}
```
