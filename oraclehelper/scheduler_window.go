package oraclehelper

import (
	"fmt"
	"log"
)

const (
	querySchedulerWindow = `
SELECT
	sw.owner,
    sw.window_name,
    sw.resource_plan,
    sw.schedule_type,
    sw.repeat_interval,
    sw.duration,
    sw.enabled,
    sw.comments
FROM dba_scheduler_windows sw
WHERE sw.owner = UPPER(:1)
AND sw.window_name = UPPER(:2)
`
	execDropWindow = `
BEGIN
    DBMS_SCHEDULER.DROP_WINDOW (
		window_name => :1
	);
END;
`
)

type (
	//ResourceSchedulerWindow ....
	ResourceSchedulerWindow struct {
		Owner          string
		WindowName     string
		ResourcePlan   string
		ScheduleType   string
		RepeatInterval string
		Duration       string
		Enabled        string
		StartDate      string
		Comments       string
		WindowPriority string
	}
	schedulerWindowService struct {
		client *Client
	}
)

func (s *schedulerWindowService) ReadSchedulerWindow(tf ResourceSchedulerWindow) (*ResourceSchedulerWindow, error) {
	log.Printf("[DEBUG] ReadSchedulerWindow windowname: %s\n", tf.WindowName)
	resourceSchedulerWindow := &ResourceSchedulerWindow{}

	err := s.client.DBClient.QueryRow(
		querySchedulerWindow,
		tf.Owner,
		tf.WindowName,
	).Scan(&resourceSchedulerWindow.Owner,
		&resourceSchedulerWindow.WindowName,
		&resourceSchedulerWindow.ResourcePlan,
		&resourceSchedulerWindow.ScheduleType,
		&resourceSchedulerWindow.RepeatInterval,
		&resourceSchedulerWindow.Duration,
		&resourceSchedulerWindow.Enabled,
		&resourceSchedulerWindow.Comments,
	)
	if err != nil {
		return nil, err
	}

	return resourceSchedulerWindow, nil
}

func (s *schedulerWindowService) CreateSchedulerWindow(tf ResourceSchedulerWindow) error {
	sqlCommand := fmt.Sprintf("BEGIN")
	sqlCommand += fmt.Sprintf(" DBMS_SCHEDULER.CREATE_WINDOW(")

	if tf.WindowName != "" {
		sqlCommand += fmt.Sprintf("window_name => '%s',", tf.WindowName)
	}
	if tf.ResourcePlan != "" {
		sqlCommand += fmt.Sprintf("resource_plan => '%s',", tf.ResourcePlan)
	}
	if tf.StartDate != "" {
		sqlCommand += fmt.Sprintf("start_date => %s,", tf.StartDate)
	}
	if tf.Duration != "" {
		sqlCommand += fmt.Sprintf("duration => %s,", tf.Duration)
	}
	if tf.RepeatInterval != "" {
		sqlCommand += fmt.Sprintf("repeat_interval => '%s',", tf.RepeatInterval)
	}
	if tf.WindowPriority != "" {
		sqlCommand += fmt.Sprintf("window_priority => '%s',", tf.WindowPriority)
	}
	if tf.Comments != "" {
		sqlCommand += fmt.Sprintf("comments => '%s'", tf.Comments)
	}
	sqlCommand += fmt.Sprintf(");")
	sqlCommand += fmt.Sprintf("END;")

	log.Printf("[DEBUG] CreateSchedulerWindow sqlcommand: %s", sqlCommand)
	_, err := s.client.DBClient.Exec(sqlCommand)
	if err != nil {
		log.Printf("[ERROR] Create schedule Window failed with error: %v\n", err)
		return err
	}
	return nil
}

func (s *schedulerWindowService) DropSchedulerWindow(tf ResourceSchedulerWindow) error {
	_, err := s.client.DBClient.Exec(execDropWindow, tf.WindowName)
	if err != nil {
		log.Printf("[ERROR] drop schedule Window failed with error: %v\n", err)
		return err
	}
	return nil
}
