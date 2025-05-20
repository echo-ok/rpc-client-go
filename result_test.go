package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// 测试用例：TC01 - outputPtr 不是指针
func TestConvertDataTo_NonPointerOutput(t *testing.T) {
	r := Result{
		Data: map[string]string{"key": "value"},
	}

	var output map[string]string
	err := r.ConvertDataTo(output) // 传入非指针
	assert.Error(t, err)
	assert.EqualError(t, err, "dstPtr 必须是一个指针")
}

// 测试用例：TC02 - Data 无法被序列化为 JSON（例如包含循环引用）
func TestConvertDataTo_MarshalError(t *testing.T) {
	type S struct {
		Self *S
	}
	data := &S{Self: &S{}}
	// 强制制造循环引用导致 Marshal 失败
	(*data).Self = data

	r := Result{
		Data: data,
	}

	var output S
	err := r.ConvertDataTo(&output)
	assert.Error(t, err)
}

// 测试用例：TC03 - 类型不匹配导致 Unmarshal 失败
func TestConvertDataTo_UnmarshalTypeError(t *testing.T) {
	r := Result{
		Data: map[string]string{"key": "value"},
	}

	var output string
	err := r.ConvertDataTo(&output)
	assert.Error(t, err)
}

// 测试用例：TC04 - 成功转换
func TestConvertDataTo_Success(t *testing.T) {
	inputData := map[string]string{"name": "test"}
	expected := inputData

	r := Result{
		Data: inputData,
	}

	var output map[string]string
	err := r.ConvertDataTo(&output)
	require.NoError(t, err)
	assert.Equal(t, expected, output)
}

// 测试用例：TC05 - 结构体成功转换
func TestConvertDataTo_StructSuccess(t *testing.T) {
	inputData := map[string]any{
		"name": "Peter",
		"age":  18,
		"sex":  "male",
		"desc": "I'm a boy",
	}

	r := Result{
		Data: inputData,
	}

	var output struct {
		Name        string
		Age         int
		Sex         string
		Description string `json:"desc"`
	}
	err := r.ConvertDataTo(&output)
	require.NoError(t, err)
	assert.Equal(t, "Peter", output.Name)
	assert.Equal(t, 18, output.Age)
	assert.Equal(t, "male", output.Sex)
	assert.Equal(t, "I'm a boy", output.Description)
}
