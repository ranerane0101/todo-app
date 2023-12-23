package interactor

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ranerane0101/domain/entity"
	mock_repository "github.com/ranerane0101/domain/mock/repository"
	"github.com/ranerane0101/domain/valueobject"

	"github.com/stretchr/testify/assert"
)

func TestTodoInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系テストコード", func(t *testing.T) {
		mockRepo := mock_repository.NewMockITodoRepository(ctrl)
		mockInteractor := NewTodoInteractor(mockRepo)

		ID := valueobject.NewTodoID("1")
		//10の長さのスライスを生成
		expectedList := make([]entity.Todo, 10)
		for i := range expectedList {
			id := valueobject.NewTodoID(fmt.Sprintf("%d", i+1))
			//整数を文字列に変換
			text := fmt.Sprintf("Task %d", i+1)
			done := i == 1 // 2番目のタスクのDoneをtrueに設定する例

			expectedList[i] = entity.Todo{
				ID:   id,
				Text: text,
				Done: done,
			}
		}

		// モックのリポジトリが呼び出されたときの振る舞いを設定
		mockRepo.EXPECT().
			FindAllTodos(ID).
			Return(expectedList, nil)

		// テスト対象のメソッドを呼び出す
		todos, err := mockInteractor.ListTodos(ID)

		// 結果を検証する
		assert.NoError(t, err)
		assert.NotNil(t, todos)
		assert.Equal(t, expectedList, todos)

	})

	t.Run("異常系テストコード", func(t *testing.T) {
		mockRepo := mock_repository.NewMockITodoRepository(ctrl)
		mockInteractor := NewTodoInteractor(mockRepo)

		invalidTodoID := valueobject.NewTodoID("123")
		expectedError := errors.New("error while fetching todos")

		// モックのリポジトリが呼び出されたときの振る舞いを設定
		mockRepo.EXPECT().
			FindAllTodos(invalidTodoID).
			Return(nil, expectedError)

		// テスト対象のメソッドを呼び出す
		todos, err := mockInteractor.ListTodos(invalidTodoID)

		// 結果を検証する
		assert.Error(t, err)
		assert.Nil(t, todos)
		assert.Equal(t, expectedError, err)
	})
}
