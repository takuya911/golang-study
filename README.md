

## Makefile
新しいSQLファイルを作成する
``` zsh
make newsql NAME=< 作成するファイルの名前 CreateChatなど >
```

マイグレーションをローカルからやる
``` zsh
make migrate MIGRATE=up
```


volume-data削除
``` zsh
docker volume rm $(docker volume ls -qf dangling=true)
```