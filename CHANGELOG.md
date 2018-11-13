# Changelog oraclehelper

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
