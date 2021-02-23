const form = document.getElementById('form')
const input = document.getElementById('input')
const todosUL = document.getElementById('todos')
// 要素を取得
const searchForm = document.getElementById('search-form')
const search = document.getElementById('search')

const baseURL = 'http://localhost/'

let todos
getAllTodo()

async function getAllTodo() {
    const response = await fetch(baseURL + 'todos')
    const todos = await response.json()
    if(todos) {
        todos.forEach(todo => addTodo(todo))
    }
}

async function store(todo){
    await fetch(baseURL + 'todo/store',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(todo),
    })
}

async function statusUpdate(id){
    await fetch(baseURL + 'todo/statusupdate',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: id}),
    })
}

async function deleteTodo(id){
    await fetch(baseURL + 'todo/delete',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({id: id}),
    })
}

// Todoを検索
async function searchTodo(searchKey){
    const response = await fetch(baseURL + 'todo/search',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({text: searchKey}),
    })
    const todos = await response.json()
    // 画面を開いた時にリストを生成する
    if(todos) {
        todos.forEach(todo => addTodo(todo))
    }
}

// Todo入力時に発火する
form.addEventListener('submit', (e) => {
    // デフォルトの動きをキャンセル
    e.preventDefault()

    // Todoを作成
    addTodo()
})

// Search時に発火する
searchForm.addEventListener('submit', (e) => {
    // デフォルトの動きをキャンセル
    e.preventDefault()

    // Todoリストをクリアする
    while(todosUL.firstChild) {
        // ulの子要素(li)を全て削除
        todosUL.removeChild(todosUL.firstChild)
    }

    const searchKey = search.value
    if (search) {
        searchTodo(searchKey)
    } else {
        // 入力値がない場合は全件取得する(検索リセット)
        getAllTodo()
    }
})

function addTodo(todo) {
    let todoText = input.value

    let id = Math.floor( Math.random() * (10000 + 1 - 1) ) + 1

    const todoData = {
        id: id,
        text: todoText
    }

    if(todo) {
        id = todo.id
        todoText = todo.text
    }

    if (!todo && todoText) {
        store(todoData)
    }

    if(todoText) {

        const todoEl = document.createElement('li')
        if(todo && todo.completed) {
            todoEl.classList.add('completed')
        }

        todoEl.innerText = todoText

        todoEl.addEventListener('click', () => {
            todoEl.classList.toggle('completed')
            statusUpdate(id)
        })

        todoEl.addEventListener('contextmenu', (e) => {
            e.preventDefault()
            todoEl.remove()
            deleteTodo(id)
        })
        todosUL.appendChild(todoEl)
        input.value = ''
    }
}
