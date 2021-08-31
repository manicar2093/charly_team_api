## BadRequest Validation errors

```
{
    "code": 400,
    "status": "Bad Request",
    "body": {
        "errors": [
            {
                "tag": "biotype",
                "validation": "required"
            }
        ],
        "message": "Request body does not satisfy needs. Please check documentation"
    }
}
```

## OK. Get Catalogs.

```
{
    "code": 200,
    "status": "OK",
    "body": {
        "biotype": [
            {
                "ID": 1,
                "Description": "biotype1",
                "CreatedAt": "2021-08-28T21:20:34.606958517-05:00"
            },
            {
                "ID": 2,
                "Description": "biotype2",
                "CreatedAt": "2021-08-28T21:20:34.606958637-05:00"
            },
            {
                "ID": 3,
                "Description": "biotype3",
                "CreatedAt": "2021-08-28T21:20:34.606958747-05:00"
            },
            {
                "ID": 4,
                "Description": "biotype4",
                "CreatedAt": "2021-08-28T21:20:34.606958867-05:00"
            }
        ],
        "roles": [
            {
                "ID": 1,
                "Description": "role1",
                "CreatedAt": "2021-08-28T21:20:34.606959959-05:00"
            },
            {
                "ID": 2,
                "Description": "role2",
                "CreatedAt": "2021-08-28T21:20:34.60696007-05:00"
            },
            {
                "ID": 3,
                "Description": "role3",
                "CreatedAt": "2021-08-28T21:20:34.60696019-05:00"
            },
            {
                "ID": 4,
                "Description": "role4",
                "CreatedAt": "2021-08-28T21:20:34.60696031-05:00"
            }
        ]
    }
}
```

## NotFound Catalog does not exists

```
{
    "code": 404,
    "status": "Not Found",
    "body": {
        "errors": "catalog 'no_exists' does not exists",
        "message": "Request body does not satisfy needs. Please check documentation"
    }
}
```