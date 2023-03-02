package dean

import (
	"github.com/zituocn/logx"
)

// Job 作业
type Job struct {

	// Tasks 多个等待处理的任务
	Tasks []ITask

	// ds 二维表数据
	ds DataSet
}

// NewJob 返回Job实例
//	载入初始数据
//	须将待处理的数据，转换成DataSet
func NewJob(ds DataSet) *Job {
	return &Job{
		ds: ds,
	}
}

// AddTask 添加待执行任务的实现
func (j *Job) AddTask(task ITask) *Job {
	j.Tasks = append(j.Tasks, task)
	return j
}

// GetData 返回job上的DataSet
func (j *Job) GetData() DataSet {
	return j.ds
}

// SetData 设置job上的DataSet
func (j *Job) SetData(ds DataSet) *Job {
	j.ds = ds
	return j
}

// Do 开始执行作业
func (j *Job) Do() {
	if len(j.Tasks) == 0 {
		logx.Errorf("没有待执行的任务")
		return
	}
	if j.ds == nil || len(j.ds) == 0 {
		logx.Errorf("没有待处理的数据")
		return
	}
	for _, task := range j.Tasks {
		name := GetStructName(task)

		logx.Infof("[%s] -> beginning", name)

		// 载入数据到任务实例
		task.LoadData(j.ds)

		// 处理数据
		err := task.Process()
		if err != nil {
			logx.Error(err)
		}

		// 返回处理后的数据到作业ds
		j.ds = task.Result()

		logx.Infof("[%s] -> finished", name)

	}
}
