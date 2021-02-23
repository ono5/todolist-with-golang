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

func TestSyncMapRepositorySearch(t *testing.T) {
	dbRepo := NewSyncMapTodoRepository()
	testData1 := domain.Todo{
		ID:        1,
		Text:      "First Post",
		Completed: false,
	}
	testData2 := domain.Todo{
		ID:        2,
		Text:      "こんにちは。読者のみなさん",
		Completed: false,
	}
	dbRepo.Store(testData1)
	dbRepo.Store(testData2)

	t.Run("Search Todo Test", func(t *testing.T) {
		// Test1 - Firstで検索
		todos, _ := dbRepo.Search("First")
		// 空でないこと
		require.NotEmpty(t, todos)
		// 取得結果が1件
		require.Equal(t, len(todos), 1)
		// 取得したタスクがID1
		require.Equal(t, todos[0].ID, 1)

		// Test2 - こんにちはで検索
		todos, _ = dbRepo.Search("こんにちは。")
		// 空でないこと
		require.NotEmpty(t, todos)
		// 取得結果が1件
		require.Equal(t, len(todos), 1)
		// 取得したタスクがID2
		require.Equal(t, todos[0].ID, 2)

		// Test3 - NORESULTで検索
		todos, _ = dbRepo.Search("NORESULT")
		// 空であること
		require.Empty(t, todos)
	})
}
