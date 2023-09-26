/*
数据模型定义
sam
2023-03-02
*/

package dean

// DataLine 一行的数据
type DataLine map[string]interface{}

// DataSet 一个二维表
type DataSet []DataLine

type DataMap map[string]DataSet

// ITask 数据处理接口定义
type ITask interface {

	// LoadData 载入数据
	LoadData(dm DataMap)

	// Process 处理数据
	Process() error

	// Result 返回结果数据
	Result() DataMap
}
