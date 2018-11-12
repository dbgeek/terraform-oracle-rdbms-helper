package oraclehelper

import (
	"log"
)

/*
	DBMS_STATS.SET_SCHEMA_PREFS(pname	=> '',
								ownname	=> '',
								pvalue 	=> '')
	DBMS_STATS.SET_TABLE_PREFS(pname=>'',
                            	ownname =>'',
								tabname => '',
								pvalue 	=> '')
DBMS_STATS.SET_GLOBAL_PREFS (
    pname     IN   VARCHAR2,
    pvalue    IN   VARCHAR2);

	DBMS_STATS.GET_PREFS (
		pname     IN   VARCHAR2,
		ownname   IN   VARCHAR2 DEFAULT NULL,
		tabname   IN   VARCHAR2 DEFAULT NULL)
	RETURN VARCHAR2;
*/
const (
	queryTablePref = `
SELECT 
	DBMS_STATS.GET_PREFS (:1,:2,:3) AS pvalue 
FROM dual
`
	querySchemaPref = `
SELECT 
	DBMS_STATS.GET_PREFS (:1,:2) AS pvalue 
FROM dual
`
	queryGlobalPref = `
SELECT 
	DBMS_STATS.GET_PREFS (:1) AS pvalue 
FROM dual
`
	SetGlobalPref = `
BEGIN
	DBMS_STATS.SET_GLOBAL_PREFS (
		pname     => :1,
		pvalue    => :2
	);
END;
`
	SetSchemaPref = `
BEGIN
	DBMS_STATS.SET_SCHEMA_PREFS(
		pname	=> :1,
		ownname	=> :2,
		pvalue 	=> :3
	);
END;
`
	SetTablePref = `
BEGIN
	DBMS_STATS.SET_TABLE_PREFS(pname=>'',
		ownname => :1,
		tabname => :2,
		pvalue 	=> :3
	);
END;
`
)

//Stats ..
type (
	//ResourceStats ....
	ResourceStats struct {
		Pname   string
		OwnName string
		TaBName string
		Pvalu   string
	}
	Stats struct {
		Pname   string
		OwnName string
		TaBName string
		Pvalu   string
	}
	statsService struct {
		client *Client
	}
)

func (r *statsService) ReadGlobalPre(tf ResourceStats) (*Stats, error) {
	log.Printf("[DEBUG] ReadGlobalPre pname: %s\n", tf.Pname)
	statsType := &Stats{}

	err := r.client.DBClient.QueryRow(queryGlobalPref, tf.Pname).Scan(&statsType.Pvalu)
	if err != nil {
		return nil, err
	}
	return statsType, nil
}

func (r *statsService) SetGlobalPre(tf ResourceStats) error {
	log.Printf("[DEBUG] SetGlobalPre pname: %s, pvalu: %s\n", tf.Pname, tf.Pvalu)

	_, err := r.client.DBClient.Exec(SetGlobalPref, tf.Pname, tf.Pvalu)
	if err != nil {
		return err
	}
	return nil
}

func (r *statsService) ReadSchemaPref(tf ResourceStats) (*Stats, error) {
	log.Printf("[DEBUG] ReadSchemaPref pname: %s owner: %s\n", tf.Pname, tf.OwnName)
	statsType := &Stats{}

	err := r.client.DBClient.QueryRow(queryTablePref, tf.Pname, tf.OwnName).Scan(&statsType.Pvalu)
	if err != nil {
		return nil, err
	}
	return statsType, nil
}
func (r *statsService) SetSchemaPre(tf ResourceStats) error {
	log.Printf("[DEBUG] SetSchemaPre pname: %sowner: %s pvalue: %s\n", tf.Pname, tf.OwnName, tf.Pvalu)

	_, err := r.client.DBClient.Exec(SetSchemaPref, tf.Pname, tf.OwnName, tf.Pvalu)
	if err != nil {
		return err
	}
	return nil
}

func (r *statsService) ReadTabPref(tf ResourceStats) (*Stats, error) {
	log.Printf("[DEBUG] ReadTabPref pname: %s owner: %s table: %s\n", tf.Pname, tf.OwnName, tf.TaBName)
	statsType := &Stats{}

	err := r.client.DBClient.QueryRow(queryTablePref, tf.Pname, tf.OwnName, tf.TaBName).Scan(&statsType.Pvalu)
	if err != nil {
		return nil, err
	}
	return statsType, nil
}
func (r *statsService) SetTabPre(tf ResourceStats) error {
	log.Printf("[DEBUG] SetTabPre pname: %s owner: %s table: %s\n", tf.Pname, tf.OwnName, tf.TaBName)

	_, err := r.client.DBClient.Exec(SetTablePref, tf.Pname, tf.OwnName, tf.TaBName, tf.Pvalu)
	if err != nil {
		return err
	}
	return nil
}
