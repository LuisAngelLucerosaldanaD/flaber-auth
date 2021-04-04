package models

import "time"

type Token struct {
	ID                   int64      `json:"id" db:"id" valid:""`
	User                 string     `json:"user" db:"user_id" valid:"required"`
	Document             int64      `json:"document" db:"document" valid:"required"`
	Doctype              string     `json:"doctype" db:"doctype" valid:"required"`
	AutoName             string     `json:"auto_name" db:"auto_name" valid:"required"`
	Process              string     `json:"process" db:"process" valid:"required"`
	Queue                string     `json:"queue" db:"queue" valid:"required"`
	NewQueues            []string   `json:"new_queues" db:"new_queues" valid:"required"`
	Execution            string     `json:"execution" db:"execution" valid:"required"`
	QueueIncomplete      *string    `json:"queue_incomplete" db:"queue_incomplete" valid:"-"`
	ExecutionIncomplete  *string    `json:"execution_incomplete" db:"execution_incomplete" valid:"-"`
	RuleIncomplete       *string    `json:"rule_incomplete" db:"rule_incomplete" valid:"-"`
	TimeEstimatedProcess *time.Time `json:"time_estimated_process" db:"time_estimated_process" valid:"-"`
	TimeEstimatedQueue   *time.Time `json:"time_estimated_queue" db:"time_estimated_queue" valid:"-"`
	Priority             *int       `json:"priority" db:"priority" valid:"required"`
	InputMessage         bool       `json:"input_message,omitempty" db:"input_message" valid:"-"`
	Response             string     `json:"response,omitempty" db:"response" valid:"-"`
	IsCompleted          bool       `json:"is_complete,omitempty" db:"is_complete" valid:"-"`
	IsRemove             bool       `json:"is_remove,omitempty" valid:"-"`
	IsTransit            bool       `json:"is_transit,omitempty" valid:"-"`
	UsersBalance         string     `json:"users_balance,omitempty" db:"assigned_user" valid:"-"`
	UsersDate            time.Time  `json:"assigned_date,omitempty" db:"assigned_date" valid:"-"`
	NewPriority          int        `json:"new_priority,omitempty" valid:"-"`
	ExecutionMain        string     `json:"execution_main,omitempty" valid:"-"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
}

