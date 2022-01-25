# Get Biotest Comparition Data

## Request

```JSON
{
    "user_uuid": "fd1772041825434fa1771f05a3df1903"
}
```

## Ok Response

If not exists, last_biotest will not be in the answere

```JSON
{
    "code": 200,
    "status": "OK",
    "body": {
        "first_biotest": {
            // biotest
        },
        "last_biotest": {
            // biotest
        },
        "all_biotests_details": [
            // biotest details
        ]
    }
}
```

## BadRequest Response

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

## Internal server error

```JSON
{
    "code": 500,
    "status": "Internal Server Error",
    "body": {
       "error": "ordinary error"
    }
}
```
