package infrastructure

import (
	"github.com/ranerane0101/domain/entity"
	"github.com/ranerane0101/domain/valueobject"
	"github.com/ranerane0101/repository"
	"gorm.io/gorm"
)

// TodoRepository はTodoモデルのリポジトリを表します。
type TodoRepository struct {
	DB *gorm.DB
}

// NewTodoRepository はTodoRepositoryのインスタンスを作成します。
func NewTodoRepository(db *gorm.DB) repository.ITodoRepository {
	return &TodoRepository{
		DB: db,
	}
}

// FindAllTodos は全てのToDoを取得します。
func (r *TodoRepository) FindAllTodos(ID valueobject.TodoID) ([]entity.Todo, error) {
	var todos []entity.Todo
	if err := r.DB.Where("user_id = ?", ID).Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// FindTodoByID は指定されたIDのToDoを取得します。
func (r *TodoRepository) FindTodoByID(ID valueobject.TodoID) (*entity.Todo, error) {
	var todo entity.Todo
	if err := r.DB.First(&todo, "id = ?", ID).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
