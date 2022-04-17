# User package

This contains all logic to handle user entity.

## Avatar User Updater

Updates user's avatar url to be shown at the app

## User Creator

Here can be created all kind of users. For each one must be consider the unique request validations besides all required data:

### Customer

Must contain:

* GenderID
* BoneDensityID
* BiotypeID

## Update User

If `id` or `uuid` are not in request a `identifier` validation error is returned
