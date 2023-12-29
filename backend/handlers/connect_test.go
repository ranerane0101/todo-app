package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/dto"
	mock_usecase "github.com/ranerane0101/mock/usecase"
	"github.com/stretchr/testify/assert"
)

func TestGetTodosHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("正常系テスト", func(t *testing.T) {
		// モックのusecase.TodoUsecaseInterfaceを作成
		mockUsecase := mock_usecase.NewMockTodoUsecaseInterface(ctrl)

		// テスト対象のハンドラーを作成
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			todoID := valueobject.NewTodoID("1")

			mockUsecase.EXPECT().
				ListTodos(todoID).
				Return([]dto.TodoDTO{}, nil)

			GetTodos(w, r, mockUsecase)
		})

		// リクエストを作成 (実際には使われませんが、ハンドラーの関数を呼び出すために必要)
		req, err := http.NewRequest("GET", "/api/todos", nil)
		assert.NoError(t, err)

		// レスポンスのWriterを作成
		rr := httptest.NewRecorder()

		// ハンドラーを呼び出す
		handler.ServeHTTP(rr, req)

		// レスポンスのステータスコードを確認
		assert.Equal(t, http.StatusOK, rr.Code)

		// レスポンスのContent-Typeを確認
		assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

		// レスポンスボディが空でないことを確認
		var response []dto.TodoDTO
		err = json.Unmarshal(rr.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.NotNil(t, response)
	})
}
