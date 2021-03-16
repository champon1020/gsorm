# Docs
mgormのドキュメントです．

[English](https://github.com/champon1020/mgorm/tree/main/docs/README.md)

## Overview
- [Instllation](https://github.com/champon1020/mgorm/tree/main/docs/README_jp.md#installation)
- [Quick Start](https://github.com/champon1020/mgorm/tree/docs/docs/README_jp.md#quick-start)
- [Introduction](https://github.com/champon1020/mgorm/tree/main/docs/introduction_jp.md)
- [Select](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md)
  - [From](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#from)
  - [Join / LeftJoin / RightJoin](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#join)
  - [Where](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#where)
  - [And / Or](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#and--or)
  - [Group By](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#groupby)
  - [Having](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#having)
  - [Union / UnionAll](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#union)
  - [Order By](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#orderby)
  - [Limit](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#limit)
  - [Offset](https://github.com/champon1020/mgorm/tree/main/docs/select_jp.md#offset)
- [Function Query](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md)
  - [Count](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#count)
  - [Sum](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#sum)
  - [Avg](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#avg)
  - [Max](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#max)
  - [Min](https://github.com/champon1020/mgorm/tree/main/docs/fnquery_jp.md#min)
- [Insert](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md)
  - [Values](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#values)
  - [Select](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#select)
  - [Model](https://github.com/champon1020/mgorm/tree/main/docs/insert_jp.md#model)
- [Update](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md)
  - [Set](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#set)
  - [Where](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#where)
  - [And / Or](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#and-or)
  - [Model](https://github.com/champon1020/mgorm/tree/main/docs/update_jp.md#model)
- [Delete](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md)
  - [From](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#from)
  - [Where](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#where)
  - [And / Or](https://github.com/champon1020/mgorm/tree/main/docs/delete_jp.md#and-or)
- [CreateDB]()
- [CreateTable]()
  - [Column]()
  - [Constraint]()
  - [Model]()
- [CreateIndex]()
- [AlterTable]()
  - [Rename]()
  - [AddColumn]()
  - [DropColumn]()
  - [RenameColumn]()
  - [AddCons]()
- [DropDB]()
- [DropTable]()
- [DropIndex]()
- [Model]()
- [Transaction]()
- [Mock]()
  - [Expect]()
  - [Return]()
  - [Complete]()

## インストール
```
go get -u github.com/champon1020/mgorm
```

## クイックスタート
```go
package main

import (
	"log"
	"time"

	"github.com/champon1020/mgorm"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID        int
	Name      string    `mgorm:"first_name"`
	BirthDate time.Time `layout:"time.RFC3339"`
}

func main() {
	db, err := mgorm.New("mysql", "root:toor@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatalln(err)
	}

	// CREATE TABLE person (
	//   id         INT,
	//   first_name VARCHAR(128),
	//   birth_date DATETIME,
	// )
	err = mgorm.CreateTable(db, "person").Model(&Person{}).Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	// INSERT INTO person VALUES (10001, 'Taro', '2006-01-02T15:04:05Z00:00')
	err = mgorm.Insert(db, "person").
		Model(&Person{
			ID:        10001,
			Name:      "Taro",
			BirthDate: time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)},
		).Exec()
	if err != nil {
		log.Fatalln(err)
	}

	// SELECT id, first_name FROM person
	p := &Person{}
	err = mgorm.Select(db, "id", "first_name").From("person").Query(p)
	if err != nil {
		log.Fatalln(err)
	}

	// UPDATE person SET first_name='Jiro'
	err = mgorm.Update(db, "person", "first_name").Set("Jiro").Exec()
	if err != nil {
		log.Fatalln(err)
	}

	// DELETE FROM person
	err = mgorm.Delete(db).From("person").Exec()
	if err != nil {
		log.Fatalln(err)
	}

	// DROP TABLE person
	err = mgorm.DropTable(db, "person").Migrate()
	if err != nil {
		log.Fatalln(err)
	}
}
```
