# README

## 這是什麼

一個簡單的在命令行 ToDo List 管理工具（iktodo 完工之前的簡易替代品）

## 安裝

### go get

```bash
go install github.com/alctny/todo
```

### 編譯安裝

```bash
git clone --depth 1 https://github.com/alctny/todo.git && cd todo && bash install.sh
```

## 使用

### 查看支持的所有命令

```sh
todo -h
```

### 添加任务

只是添加一个任务

```sh
todo add "大吉大利，今晚吃鸡"
```

附加一些额外的信息

```sh
todo add "娱乐:game,Battlegrounds:winner winner chicken dinner"
```

完整的格式是 `class:tag1,tag2...:task`。使用冒号作为分隔，第一部分表示 class，表示任务的分类，后续版本中可能会作为任务列表的列表名。第二部分表示 tag，一个任务可以有多个 tag，tag 之间使用(英文)逗号分隔。第三部分表示任务名称

> P.S
>
> 1. 如果任务名中有空格最好使用引号将其包裹，避免任务名称被截断
> 2. 如果只想填写 class 或 tag 其中一个只需要使用 `class::task` 或 `:tag:task` 即可
> 3. 如果任务名称有英文冒号，但你并不希望添加 class 和 tag，可使用 `::task`
>

### 查看

查看任务列表，可以通过 `--class` 和 `--tag` 进行筛选，如果不添加任何参数则是默认查看全部

```sh
todo list --class 娱乐 --tag Battlegrounds
```

如果你希望查看特定状态的任务则可以使用 `--undone` 或 `--done`

### 标示完成/未完成

新创建的任务默认处于待办状态，如果希望将其设置为完成状态则可以使用

```sh
todo done 01
```

`done` 会翻转任务的状态，如果任务已完成，则将其修改为待办，如果任务未完成则将其修改为已完成

### 排序

`todo done` 和 `todo add` 不会触发排序，如果希望对列表排序可以使用

```sh
todo sort
```

这条命令可以让未完成任务上浮到顶部，已完成任务下沉到底部

### 清理

如果需要清理列表中的所有已完成任务则使用 `clear` 命令

```sh
todo clear
```

### 重命名

需要对任务重命名则使用 `mv`

```sh
todo mv 00 召唤师峡谷
```

### 等价行为

1. `todo` = `todo list`
2. `todo game:lol:pvp` = `todo add game:lol:pvp`

> P.S. 虽然 todo 等价 todo list，但是 todo --class 并不等价 todo list --class。todo 命令本身没有任何 Flag（无值参数） 或 Option（有值参数）
