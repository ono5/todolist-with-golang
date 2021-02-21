// 要素の取得
const form = document.getElementById('form')
const input = document.getElementById('input')
const todosUL = document.getElementById('todos')

// ローカルストレージからデータを取得
const todos = JSON.parse(localStorage.getItem('todos'))

// 画面を開いた時にリストを生成する
if(todos) {
    todos.forEach(todo => addTodo(todo))
}

// Todo入力時に発火する
form.addEventListener('submit', (e) => {
    // デフォルトの動きをキャンセル
    e.preventDefault()

    // Todoを作成
    addTodo()
})

// Todo作成
function addTodo(todo) {
    // 入力文字を取得
    let todoText = input.value

    if(todo) {
        todoText = todo.text
    }

    if(todoText) {
        // liリストを作成
        const todoEl = document.createElement('li')
        // タスク完了かチェック
        // todoはオブジェクトを格納でき、completedを持っている
        if(todo && todo.completed) {
            // completedクラスをつける
            todoEl.classList.add('completed')
        }

        todoEl.innerText = todoText

        // Todoをクリックした時に発火
        todoEl.addEventListener('click', () => {
            // completedクラスがついてたら削除、そうでない場合は付与
            todoEl.classList.toggle('completed')
            // ローカルストレージを更新
            updateLS()
        })

        // 右クリックした時に発火するイベント
        todoEl.addEventListener('contextmenu', (e) => {
            // デフォルトの動きをキャンセル
            e.preventDefault()

            // Todo削除
            todoEl.remove()
            // ローカルストレージを更新
            updateLS()
        })

        // Todoを親要素の子要素として追加
        todosUL.appendChild(todoEl)

        // 入力欄をクリア
        input.value = ''

        // ローカルストレージを更新
        updateLS()
    }
}

// ローカルストレージ更新
function updateLS() {
    // li要素を取得
    todosEl = document.querySelectorAll('li')

    // 保存データ
    const todos = []

    todosEl.forEach(todoEl => {
        // オブジェクトを配列にpush
        todos.push({
            text: todoEl.innerText,
            completed: todoEl.classList.contains('completed')
        })
    })

    // ローカルストレージにtodosキーで保存保存
    localStorage.setItem('todos', JSON.stringify(todos))
}
