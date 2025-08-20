package rpclient

import (
	"testing"
)

func TestPagerStruct(t *testing.T) {
	// Test Pager struct creation and field types
	pager := Pager{
		Page:       1,
		PageSize:   10,
		TotalCount: 100,
		PageCount:  10,
		IsLastPage: false,
		Items:      []string{"item1", "item2"},
	}
	
	if pager.Page != 1 {
		t.Errorf("Expected Page to be 1, got %d", pager.Page)
	}
	if pager.PageSize != 10 {
		t.Errorf("Expected PageSize to be 10, got %d", pager.PageSize)
	}
	if pager.TotalCount != 100 {
		t.Errorf("Expected TotalCount to be 100, got %d", pager.TotalCount)
	}
	if pager.PageCount != 10 {
		t.Errorf("Expected PageCount to be 10, got %d", pager.PageCount)
	}
	if pager.IsLastPage != false {
		t.Errorf("Expected IsLastPage to be false, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.([]string)
	if !ok {
		t.Errorf("Expected Items to be []string, got %T", pager.Items)
	}
	if len(items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(items))
	}
	if items[0] != "item1" {
		t.Errorf("Expected first item to be 'item1', got '%s'", items[0])
	}
	if items[1] != "item2" {
		t.Errorf("Expected second item to be 'item2', got '%s'", items[1])
	}
}

func TestPagerWithZeroValues(t *testing.T) {
	// Test Pager with zero values
	pager := Pager{
		Page:       0,
		PageSize:   0,
		TotalCount: 0,
		PageCount:  0,
		IsLastPage: false,
		Items:      nil,
	}
	
	if pager.Page != 0 {
		t.Errorf("Expected Page to be 0, got %d", pager.Page)
	}
	if pager.PageSize != 0 {
		t.Errorf("Expected PageSize to be 0, got %d", pager.PageSize)
	}
	if pager.TotalCount != 0 {
		t.Errorf("Expected TotalCount to be 0, got %d", pager.TotalCount)
	}
	if pager.PageCount != 0 {
		t.Errorf("Expected PageCount to be 0, got %d", pager.PageCount)
	}
	if pager.IsLastPage != false {
		t.Errorf("Expected IsLastPage to be false, got %t", pager.IsLastPage)
	}
	if pager.Items != nil {
		t.Errorf("Expected Items to be nil, got %v", pager.Items)
	}
}

func TestPagerWithLastPage(t *testing.T) {
	// Test Pager with IsLastPage set to true
	pager := Pager{
		Page:       5,
		PageSize:   20,
		TotalCount: 100,
		PageCount:  5,
		IsLastPage: true,
		Items:      []int{1, 2, 3},
	}
	
	if pager.Page != 5 {
		t.Errorf("Expected Page to be 5, got %d", pager.Page)
	}
	if pager.PageSize != 20 {
		t.Errorf("Expected PageSize to be 20, got %d", pager.PageSize)
	}
	if pager.TotalCount != 100 {
		t.Errorf("Expected TotalCount to be 100, got %d", pager.TotalCount)
	}
	if pager.PageCount != 5 {
		t.Errorf("Expected PageCount to be 5, got %d", pager.PageCount)
	}
	if pager.IsLastPage != true {
		t.Errorf("Expected IsLastPage to be true, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.([]int)
	if !ok {
		t.Errorf("Expected Items to be []int, got %T", pager.Items)
	}
	if len(items) != 3 {
		t.Errorf("Expected 3 items, got %d", len(items))
	}
}

func TestPagerWithEmptyItems(t *testing.T) {
	// Test Pager with empty items slice
	pager := Pager{
		Page:       1,
		PageSize:   10,
		TotalCount: 0,
		PageCount:  0,
		IsLastPage: true,
		Items:      []string{},
	}
	
	if pager.Page != 1 {
		t.Errorf("Expected Page to be 1, got %d", pager.Page)
	}
	if pager.PageSize != 10 {
		t.Errorf("Expected PageSize to be 10, got %d", pager.PageSize)
	}
	if pager.TotalCount != 0 {
		t.Errorf("Expected TotalCount to be 0, got %d", pager.TotalCount)
	}
	if pager.PageCount != 0 {
		t.Errorf("Expected PageCount to be 0, got %d", pager.PageCount)
	}
	if pager.IsLastPage != true {
		t.Errorf("Expected IsLastPage to be true, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.([]string)
	if !ok {
		t.Errorf("Expected Items to be []string, got %T", pager.Items)
	}
	if len(items) != 0 {
		t.Errorf("Expected 0 items, got %d", len(items))
	}
}

func TestPagerWithMapItems(t *testing.T) {
	// Test Pager with map items
	pager := Pager{
		Page:       1,
		PageSize:   10,
		TotalCount: 1,
		PageCount:  1,
		IsLastPage: true,
		Items:      map[string]interface{}{"key": "value"},
	}
	
	if pager.Page != 1 {
		t.Errorf("Expected Page to be 1, got %d", pager.Page)
	}
	if pager.PageSize != 10 {
		t.Errorf("Expected PageSize to be 10, got %d", pager.PageSize)
	}
	if pager.TotalCount != 1 {
		t.Errorf("Expected TotalCount to be 1, got %d", pager.TotalCount)
	}
	if pager.PageCount != 1 {
		t.Errorf("Expected PageCount to be 1, got %d", pager.PageCount)
	}
	if pager.IsLastPage != true {
		t.Errorf("Expected IsLastPage to be true, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.(map[string]interface{})
	if !ok {
		t.Errorf("Expected Items to be map[string]interface{}, got %T", pager.Items)
	}
	if len(items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(items))
	}
	if items["key"] != "value" {
		t.Errorf("Expected item value to be 'value', got '%v'", items["key"])
	}
}

func TestPagerWithStructItems(t *testing.T) {
	// Test Pager with struct items
	type TestStruct struct {
		Name string
		Age  int
	}
	
	testItem := TestStruct{Name: "John", Age: 30}
	pager := Pager{
		Page:       1,
		PageSize:   10,
		TotalCount: 1,
		PageCount:  1,
		IsLastPage: true,
		Items:      testItem,
	}
	
	if pager.Page != 1 {
		t.Errorf("Expected Page to be 1, got %d", pager.Page)
	}
	if pager.PageSize != 10 {
		t.Errorf("Expected PageSize to be 10, got %d", pager.PageSize)
	}
	if pager.TotalCount != 1 {
		t.Errorf("Expected TotalCount to be 1, got %d", pager.TotalCount)
	}
	if pager.PageCount != 1 {
		t.Errorf("Expected PageCount to be 1, got %d", pager.PageCount)
	}
	if pager.IsLastPage != true {
		t.Errorf("Expected IsLastPage to be true, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.(TestStruct)
	if !ok {
		t.Errorf("Expected Items to be TestStruct, got %T", pager.Items)
	}
	if items.Name != "John" {
		t.Errorf("Expected Name to be 'John', got '%s'", items.Name)
	}
	if items.Age != 30 {
		t.Errorf("Expected Age to be 30, got %d", items.Age)
	}
}

func TestPagerWithNegativeValues(t *testing.T) {
	// Test Pager with negative values (edge case)
	pager := Pager{
		Page:       -1,
		PageSize:   -10,
		TotalCount: -100,
		PageCount:  -5,
		IsLastPage: false,
		Items:      nil,
	}
	
	if pager.Page != -1 {
		t.Errorf("Expected Page to be -1, got %d", pager.Page)
	}
	if pager.PageSize != -10 {
		t.Errorf("Expected PageSize to be -10, got %d", pager.PageSize)
	}
	if pager.TotalCount != -100 {
		t.Errorf("Expected TotalCount to be -100, got %d", pager.TotalCount)
	}
	if pager.PageCount != -5 {
		t.Errorf("Expected PageCount to be -5, got %d", pager.PageCount)
	}
	if pager.IsLastPage != false {
		t.Errorf("Expected IsLastPage to be false, got %t", pager.IsLastPage)
	}
	if pager.Items != nil {
		t.Errorf("Expected Items to be nil, got %v", pager.Items)
	}
}

func TestPagerWithLargeValues(t *testing.T) {
	// Test Pager with large values
	pager := Pager{
		Page:       999999,
		PageSize:   999999,
		TotalCount: 999999999,
		PageCount:  999999,
		IsLastPage: false,
		Items:      make([]int, 1000),
	}
	
	if pager.Page != 999999 {
		t.Errorf("Expected Page to be 999999, got %d", pager.Page)
	}
	if pager.PageSize != 999999 {
		t.Errorf("Expected PageSize to be 999999, got %d", pager.PageSize)
	}
	if pager.TotalCount != 999999999 {
		t.Errorf("Expected TotalCount to be 999999999, got %d", pager.TotalCount)
	}
	if pager.PageCount != 999999 {
		t.Errorf("Expected PageCount to be 999999, got %d", pager.PageCount)
	}
	if pager.IsLastPage != false {
		t.Errorf("Expected IsLastPage to be false, got %t", pager.IsLastPage)
	}
	
	items, ok := pager.Items.([]int)
	if !ok {
		t.Errorf("Expected Items to be []int, got %T", pager.Items)
	}
	if len(items) != 1000 {
		t.Errorf("Expected 1000 items, got %d", len(items))
	}
}