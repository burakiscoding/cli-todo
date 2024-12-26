# CLI To-Do

Terminal app developed in Go. Manage your todos.

## How to use

Create new todo:

```bash
./cli-todo create --title="Drink water"
```

<hr>
Read todos:

```bash
./cli-todo read
```

<hr>
Read todo by id:

```bash
./cli-todo read --id=1
```

<hr>
Update todo:
(<strong>--checked</strong> parameter is true as default)

```bash
./cli-todo update --id=1 --title="Play some Fortnite OG"
```

<hr>
Delete todo:

```bash
./cli-todo delete --id=1
```

## Preview

<img alt="Screenshot of the terminal app that shows CRUD operations and results" src="./previews/preview.png">
