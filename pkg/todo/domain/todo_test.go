package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	type args struct {
		task string
	}
	tests := []struct {
		name string
		args args
		want Todo
	}{
		{
			name: "When a new task is added should create a todo",
			args: args{task: "dummy task"},
			want: Todo{task: "dummy Task", status: Created, creationDate: time.Now(), updateTime: time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTodo(tt.args.task); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTodo() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestTodo_GetTask(t *testing.T) {
//	type fields struct {
//		task         string
//		status       int
//		creationDate time.Time
//		updateTime   time.Time
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			todo := Todo{
//				task:         tt.fields.task,
//				status:       tt.fields.status,
//				creationDate: tt.fields.creationDate,
//				updateTime:   tt.fields.updateTime,
//			}
//			if got := todo.GetTask(); got != tt.want {
//				t.Errorf("GetTask() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestTodo_UpdateStatus(t *testing.T) {
//	type fields struct {
//		task         string
//		status       int
//		creationDate time.Time
//		updateTime   time.Time
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		want   Todo
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			todo := Todo{
//				task:         tt.fields.task,
//				status:       tt.fields.status,
//				creationDate: tt.fields.creationDate,
//				updateTime:   tt.fields.updateTime,
//			}
//			if got := todo.UpdateStatus(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("UpdateStatus() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
