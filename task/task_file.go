package task

import "github.com/beego/beego/v2/task"

func CreateTask(taskName, schedule_time string, f task.TaskFunc) {
	task_Perform := task.NewTask(taskName, schedule_time, f)
	task.AddTask(taskName, task_Perform)
	task.StartTask()
	defer task.StopTask()
}
