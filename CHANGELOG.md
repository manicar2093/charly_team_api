## v0.55.0 (2022-01-23)

### Feat

- add get catalogs aws lambda controller (#156)

## v0.54.0 (2022-01-23)

### Feat

- add creation user aws lambda handler (#155)

## v0.53.0 (2022-01-21)

### Feat

- Add create biotest aws lambda controller and implementation (#154)

## v0.52.0 (2022-01-20)

### Feat

- Add biotest update aws lambda controller (#153)

## v0.51.0 (2022-01-19)

### Feat

- add update user handler and aws lambda controller (#152)

## v0.50.1 (2022-01-18)

### Refactor

- Remove mocks path (#151)

## v0.50.0 (2022-01-06)

### Feat

- Add user repository (#149)

## v0.49.0 (2022-01-05)

### Feat

- Add biotest repository (#148)

## v0.48.0 (2022-01-04)

### Feat

- Add total entries pagination (#145)

## v0.47.0 (2021-12-16)

### Feat

- Change way to get data to create pageSort (#144)

## v0.46.0 (2021-12-16)

### Feat

- Add sort queries on paginator (#142)

## v0.45.0 (2021-12-14)

### Feat

- Add TotalPages to Paginator struct (#139)

## v0.44.0 (2021-11-26)

### Feat

- Add user data return on response (#135)

## v0.43.1 (2021-11-26)

### Fix

- Remove lowercase required on serverless (#133)

## v0.43.0 (2021-11-26)

### Feat

- Add ordering on get all user biotest (#131)

## v0.42.0 (2021-11-26)

### Feat

- Add ordering on AllBiotestDetails

## v0.41.0 (2021-11-16)

### Feat

- Add handler to save biotest images (#128)

## v0.40.0 (2021-11-15)

### Feat

- Add cors config into serverles.yml (#126)

## v0.39.0 (2021-11-04)

### Feat

- Add s3 resource form biotest images (#124)

## v0.38.0 (2021-11-04)

### Feat

- Change code to not return last biotest if not found (#123)

## v0.37.1 (2021-11-04)

### Fix

- Add fix to find user by email, name and last name (#121)

## v0.37.0 (2021-10-30)

### Feat

- Add data to cognito lambda (#118)

## v0.36.0 (2021-10-22)

### Feat

- Find user by uuid to get session properties to front

### Refactor

- Remove unused handler

### Fix

- Add int validator to create sql.Int null

## v0.35.1 (2021-10-20)

### Fix

- Add solution to user filter (#111)

## v0.35.0 (2021-10-18)

### Feat

- Add lambda to retreive user info when login (#109)

## v0.34.0 (2021-10-18)

### Feat

- Add new Pagination implementation.
- Add json tags to catalog entities
- Add authflow for cognito

### Fix

- Add fix to emailSubject on serverless yml
- Change catalog name
- Fix error on page has just 1 entry

## v0.33.1 (2021-10-11)

### Fix

- Add biotest fixes for found bugs (#104)

## v0.33.0 (2021-10-11)

### Fix

- Add svg format to avatar Url on user creation
- Change way to save biotest.
- Add changes and test to biotest entity

### Feat

- Add json tags to CustomerValidation struct

## v0.32.0 (2021-10-11)

### Feat

- Modify user filter to saerch by email and names Close#82
- Add filter to return biotests as catalog Close#98
- Add autoload for customer in biotype entity Close #95
- Add customer validations Close  #97
- Add avatar url on user service
- Add serverless prune plugin
- Add env variables for Prod migrations
- Add avatar url

### Refactor

- Remove dead code Close#93
- Add fixes for migrations issues

### Fix

- Add fix to Paginator package Close#96
- Add page_number as float.

## v0.31.0 (2021-10-10)

### Feat

- Add filter to get all user biotest (#92)

## v0.30.0 (2021-10-07)

### Feat

- Add seed fixes and uuid support in User and Biotest (#91)

## v0.29.0 (2021-10-07)

### Feat

- Add default avatar url for users (#90)

## v0.28.1 (2021-09-30)

### Fix

- Change attributes Biotest names (#88)

## v0.28.0 (2021-09-30)

### Feat

- Add data to user table

## v0.27.0 (2021-09-29)

### Feat

- Add registry customer function: (#84)

## v0.26.0 (2021-09-12)

### Feat

- Add biotest comparision search (#71)

## v0.25.0 (2021-09-12)

### Feat

- Change filter implementation (#70)

## v0.24.0 (2021-09-11)

### Feat

- Change search by user_id to user_uuid (#69)

## v0.23.0 (2021-09-11)

### Feat

- Add update user function. (#68)

## v0.22.0 (2021-09-11)

### Feat

- Create update biotest function (#67)

## v0.21.1 (2021-09-11)

### Fix

- Add changes to fix issues (#65)

## v0.21.0 (2021-09-09)

### Feat

- Add biotest search to filters (#60)

## v0.20.0 (2021-09-09)

### Feat

- Add serverless config test (#59)

## v0.19.0 (2021-09-05)

### Feat

- Add UUID user biotest table (#52)

## v0.18.0 (2021-09-05)

### Feat

- many user feats (#50)

## v0.17.0 (2021-09-03)

### Feat

- Add find biotest by (#46)

## v0.16.0 (2021-09-02)

### Feat

- Change user gender to nullable (#44)

## v0.15.0 (2021-09-02)

### Feat

- registry bio test feature (#40)

## v0.14.0 (2021-08-31)

### Feat

- rel orm implementation (#37)

## v0.13.0 (2021-08-30)

### Feat

- errors handling and implementation (#35)

## v0.12.0 (2021-08-29)

### Feat

- create get catalogs function (#32)

## v0.11.0 (2021-08-28)

### Feat

- Add UserId created to create user response (#29)

## v0.10.0 (2021-08-28)

### Feat

- Add registry user function (#27)

## v0.9.0 (2021-08-26)

### Feat

- Add cognito client generator with unit test (#26)

## v0.8.0 (2021-08-24)

### Feat

- Add aws session and env variables (#23)

## v0.7.0 (2021-08-22)

### Feat

- Add validation service for structs. (#20)

## v0.6.1 (2021-08-22)

### Fix

- Change configurations starting. (#19)

## v0.6.0 (2021-08-22)

### Feat

- Add config and connections packages (#18)

## v0.5.0 (2021-08-19)

### Feat

- Add workflow to run test CI (#15)

## v0.4.0 (2021-08-17)

### Feat

- Add new configuration for knex.
- Add sqlite3 dependency on NPM.

## v0.3.0 (2021-08-17)

### Feat

- change db structure (#5)

## v0.2.0 (2021-08-16)

### Feat

- add some features for serverless (#4)

### Fix

- change bump_message [skip ci]
- Change if condition to bump [skip ci]

## v0.1.0 (2021-08-01)

### Feat

- Change target branch in commitizen workflow
- Add commitizen workflow (#3)
- create DB Entities (#2)
- db creation (#1)

## v0.0.1 (2021-07-17)
