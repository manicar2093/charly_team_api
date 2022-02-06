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

## NotFound Response

```JSON
{
    "code": 404,
    "status": "Not Found",
    "body": {
        "error": "BiotestComparitionData with identifier 214dde36ed1a437baff06f515633048c not found"
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
