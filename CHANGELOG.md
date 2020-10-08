# Changelog oraclehelper

## 0.5.0 (October, 8, 2020)

* refactor: newClient return err now if db.Ping not succeed
* refactor: newClient now using godror as driver for Oracle
* dependency: Bumping hashicorp/go-version & terraform-plugin-sdk
* test: Make aware of multitenant is used or not
* bug: connect to db crash if using special character like `@`

## 0.4.2 (January 3, 2020)

* Bug: Need to quote the password if it contain special characters. Issue #30

## 0.4.1 (January 2, 2020)

* Bug: Adding special character and numeric to the password. So create user will not fail if using `ORA12C_STIG_VERIFY_FUNCTION`. Closing issue #30.

## 0.4.0 (October 14, 2019 )

* feature: updating dependency
* refactor: switching to terraform-plugin-sdk/helper/acctest instead of terraform/helper/acctest
* refactor: cleanup trailing whitespace
* test: fixing dependency that table should exists
* feature: adding audit user

## 0.3.1 (February 14, 2019 )

* Switch to goracle package as driver for connect to Oracle

## 0.3.0 (December 25, 2018)

* Update vendoring for go-oci8.
* Switch to use go mod instead of go vendor with go dep utility.
* Adding blockChangeTrackingService API and test for service.

## 0.2.9 (December 1, 2018)

* Adding API for Managing Automated Database Maintenance Tasks
* Adding database API for force logging and flashback on off

## 0.2.8.1 (November 25, 2018)

* Generate random string password when creating a user
* Change how to handle account status

## 0.2.8

* Adding vendoring hashicorp/terraform
* Adding vendoring mattrobenolt/size
* Adding vendoring golang.org/x/crypto
* Adding quota support for user API
* Adding account status support for user API
* Adding profile support for user API
* Adding SchedulerWindowService to oracle helper

## 0.2.7

* Removing ReadSchemaPref as Oracle have no support for that and bug fix SetTabPre.

## 0.2.6

* ReadSchemaPref was using queryTablePref instead of querySchemaPref query

## 0.2.5

* Adding Read / Set dbms_stats pref

## 0.2.4 (November 11, 2018)

* Revoking priviliges that are assignes to a user

## 0.2.3 (November 9, 2018)

* Switching to sha256 for generate hash
* Updating go-oci8 vendoring

## 0.2.2 (November 5, 2018)

* ReadProfile return error if not exists in db
* Bug fixed ReadGrantObjectPrivilege. Did not bind the result to correct variable. Copy Past bug
* First support in grant api to grant & revoke on schema level for tables.

## 0.2.1 (November 2, 2018)

* Making profile & grant support different version of Oracle.

## 0.2.0 (October 23, 2018)

Changes done on flight to SFO for Hashiconf

* Use UPPER function around bind variables
* Making scope as a variable to control the scope of setting the parameter. Memory + Spfile + Both

## 0.1.0 (October 14, 2018)

NOTES:

initial release of db api

* grant
* profile
* parameter
* role
* user
