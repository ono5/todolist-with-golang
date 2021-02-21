package repository

// repository/todo_map_test.go

import (
	"testing"
	"todo/domain"

	"github.com/stretchr/testify/require"
)

func TestSyncMapRepository(t *testing.T) {
	// レポジトリをインスタンス化
	// このように操作したいDBのインターフェースを
	// 返せばメソッドを呼び出せる
	dbRepo := NewSyncMapTodoRepository()
	// 例えば、PostgreSQLがを導入したい場合は、
	// 以下のように呼びだすとテストコードの変更を最小限にできる
	// dbRepo := NewPostgreSQLTodoRepository()
	id := 1
	testData := domain.Todo{
		ID:        id,
		Text:      "test",
		Completed: false,
	}
	t.Run("Allget Todo Test", func(t *testing.T) {
		// Allget Todo
		r1, _ := dbRepo.AllGet()
		require.Empty(t, r1)
	})

	t.Run("Store Todo Test", func(t *testing.T) {
		dbRepo.Store(testData)
		r2, _ := dbRepo.AllGet()
		require.Equal(t, r2[0], testData)
	})

	t.Run("Status Update Test", func(t *testing.T) {
		dbRepo.StatusUpdate(id)
		r3, _ := dbRepo.AllGet()
		require.Equal(t, r3[0].Completed, true)
	})

	t.Run("Delete Todo Test", func(t *testing.T) {
		dbRepo.Delete(id)
		r4, _ := dbRepo.AllGet()
		require.Empty(t, r4)
	})
}
