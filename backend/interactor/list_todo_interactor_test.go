package interactor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
	mock_repository "github.com/ranerane0101/mock"

	"github.com/stretchr/testify/assert"
)

func TestGetTodoList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockITodoRepository(ctrl)
	mockInteractor := NewTodoInteractor(mockRepo)

	userID := "123"

	// 期待されるリスト
	expectedList := []entity.Todo{
		{ID: valueobject.NewTodoID("1"), Text: "Task 1", Done: false},
		{ID: valueobject.NewTodoID("2"), Text: "Task 2", Done: true},
	}

	// モックのリポジトリが呼び出されたときの振る舞いを設定
	mockRepo.EXPECT().
		FindAllTodos(userID).
		Return(expectedList, nil)

	// テスト対象のメソッドを呼び出す
	todos, err := mockInteractor.GetTodoList(userID)

	// 結果を検証する
	assert.NoError(t, err)
	assert.NotNil(t, todos)
	assert.Equal(t, expectedList, todos)

	fmt.Println("expectedList:", expectedList)
	fmt.Println("todos:", todos)
}

func TestGetTodoList_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockITodoRepository(ctrl)
	mockInteractor := NewTodoInteractor(mockRepo)

	userID := "不正なuserID"
	expectedError := errors.New("error while fetching todos")

	// モックのリポジトリが呼び出されたときの振る舞いを設定
	mockRepo.EXPECT().
		FindAllTodos(userID).
		Return(nil, expectedError)

	// テスト対象のメソッドを呼び出す
	todos, err := mockInteractor.GetTodoList(userID)

	// 結果を検証する
	assert.Error(t, err)
	assert.Nil(t, todos)
	assert.Equal(t, expectedError, err)
}
