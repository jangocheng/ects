package models

import (
	"github.com/betterde/ects/internal/utils"
)

type PipelineTaskPivot struct {
	Id          string     `json:"id" xorm:"not null pk comment('ID') CHAR(36)"`
	PipelineId  string     `json:"pipeline_id" validate:"required,uuid4" xorm:"not null comment('ID') index CHAR(36)"`
	TaskId      string     `json:"task_id" validate:"required,uuid4" xorm:"not null comment('ID') index CHAR(36)"`
	Step        int        `json:"step" validate:"numeric" xorm:"not null comment('步骤') SMALLINT(5)"`
	Timeout     int        `json:"timeout" validate:"numeric" xorm:"not null default 0 comment('超时时间') INT(10)"`
	Interval    int        `json:"interval" validate:"numeric" xorm:"not null default 0 comment('间隔时间') INT(10)"`
	Retries     int        `json:"retries" validate:"numeric" xorm:"not null default 0 comment('重试次数') TINYINT(3)"`
	Directory   string     `json:"directory" validate:"omitempty" xorm:"null comment('工作目录') VARCHAR(255)"`
	User        string     `json:"user" validate:"omitempty" xorm:"null comment('运行用户') VARCHAR(255)"`
	Environment string     `json:"environment" validate:"omitempty" xorm:"null comment('环境变量') VARCHAR(255)"`
	Dependence  string     `json:"dependence" validate:"required" xorm:"not null default 'strong' comment('依赖') VARCHAR(255)"`
	CreatedAt   utils.Time `json:"created_at" validate:"-" xorm:"not null created comment('创建于') DATETIME"`
	UpdatedAt   utils.Time `json:"updated_at" validate:"-" xorm:"not null updated comment('更新于') DATETIME"`
	Task        *Task      `json:"task" validate:"-" xorm:"extends"`
}

// 定义模型的数据表名称
func (pivot *PipelineTaskPivot) TableName() string {
	return "pipeline_task_pivot"
}

// 创建流水线和任务的关联关系
func (pivot *PipelineTaskPivot) Store() error {
	_, err := Engine.Insert(pivot)
	return err
}

func (pivot *PipelineTaskPivot) Update() error {
	_, err := Engine.Id(pivot.Id).Cols("step").Update(pivot)
	return err
}

// Delete a pipeline relation
func (pivot *PipelineTaskPivot) Destroy() error {
	_, err := Engine.Delete(pivot)
	return err
}
