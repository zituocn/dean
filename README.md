# Dean

一个数据处理、数据清洗的任务流框架



### 安装

```shell
go get github.com/zituocn/dean
```

### 使用demo

```go
package main

import (
	"fmt"
	"github.com/zituocn/dean"
)

func main() {

	// 原始数据
	ds := loadData()

	//创建作业
	job := dean.NewJob(ds)

	//添加任务实现
	job.AddTask(new(SomeTask)).
		AddTask(new(OtherTask))

	// 执行作业
	job.Do()

	// 返回最终数据
	result := job.GetData()

	for _, item := range result {
		fmt.Println(item)
	}
}

/*
task 1
*/

type SomeTask struct {
	ds dean.DataSet
}

func (m *SomeTask) LoadData(ds dean.DataSet) {
	fmt.Println("load data")
	m.ds = ds
}

func (m *SomeTask) Process() error {
	fmt.Println("Process")
	m.ds[0]["name"] = "四川大学"
	return nil
}

func (m *SomeTask) Result() dean.DataSet {
	fmt.Println("result")
	return m.ds
}

/*
task 2
*/

type OtherTask struct {
	ds dean.DataSet
}

func (m *OtherTask) LoadData(ds dean.DataSet) {
	fmt.Println("other load data")
	m.ds = ds
}

func (m *OtherTask) Process() error {
	fmt.Println("other Process")
	m.ds[5]["major_name"] = "计算机科学与技术(工科实验班)"
	return nil
}

func (m *OtherTask) Result() dean.DataSet {
	fmt.Println("other result")
	return m.ds
}

func loadData() (ds dean.DataSet) {
	for i := 0; i < 100; i++ {
		dataLine := dean.DataLine{
			"name":          "清华大学",
			"prov_id":       "51",
			"major_name":    "工业设计",
			"number":        "43",
			"min_score":     "664",
			"min_score_def": "166",
			"avg_score":     "665",
			"max_score":     "669",
		}
		ds = append(ds, dataLine)
	}

	return ds
}

```