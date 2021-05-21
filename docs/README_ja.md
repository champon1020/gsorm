# Docs
gsormのドキュメントです．

[English](https://github.com/champon1020/gsorm/tree/main/docs/README.md)

## Overview
- [Instllation](https://github.com/champon1020/gsorm/tree/main/docs/README_ja.md#installation)
- [Quick Start](https://github.com/champon1020/gsorm/tree/docs/docs/README_ja.md#quick-start)
- [Introduction](https://github.com/champon1020/gsorm/tree/main/docs/introduction_ja.md)
- [Select](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md)
  - [From](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#from)
  - [Join](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#join)
  - [LeftJoin](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#leftjoin)
  - [RightJoin](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#rightjoin)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#and)
  - [Or](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#or)
  - [Group By](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#groupby)
  - [Having](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#having)
  - [Union](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#union)
  - [Union](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#unionall)
  - [Order By](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#orderby)
  - [Limit](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#limit)
  - [Offset](https://github.com/champon1020/gsorm/tree/main/docs/select_ja.md#offset)
- [Function Query](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md)
  - [Count](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#count)
  - [Sum](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#sum)
  - [Avg](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#avg)
  - [Max](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#max)
  - [Min](https://github.com/champon1020/gsorm/tree/main/docs/fnquery_ja.md#min)
- [Insert](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md)
  - [Values](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#values)
  - [Select](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#select)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/insert_ja.md#model)
- [Update](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md)
  - [Set](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#set)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#and)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#or)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/update_ja.md#model)
- [Delete](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md)
  - [From](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#from)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#and)
  - [Or](https://github.com/champon1020/gsorm/tree/main/docs/delete_ja.md#or)
- [CreateDB](https://github.com/champon1020/gsorm/tree/main/docs/createdb_ja.md)
- [CreateIndex](https://github.com/champon1020/gsorm/tree/main/docs/createindex_ja.md)
- [CreateTable](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md)
  - [Column](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#column)
    - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#notnull)
    - [Default](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#default)
  - [Cons](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#cons)
    - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#unique)
    - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#primary)
    - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#foreign)
      - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#ref)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/createtable_ja.md#model)
- [AlterTable](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md)
  - [Rename](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#rename)
  - [AddColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcolumn)
    - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#notnull)
    - [Default](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#default)
  - [DropColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#dropcolumn)
  - [RenameColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#renamecolumn)
  - [AddCons](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#addcons)
    - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#unique)
    - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#primary)
    - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#foreign)
      - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/altertable_ja.md#ref)
- [DropDB](https://github.com/champon1020/gsorm/tree/main/docs/dropdb_ja.md)
- [DropTable](https://github.com/champon1020/gsorm/tree/main/docs/droptable_ja.md)
- [Raw](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md)
  - [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawclause)
  - [RawStmt](https://github.com/champon1020/gsorm/tree/main/docs/raw_ja.md#rawstmt)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/model_ja.md)
- [Transaction](https://github.com/champon1020/gsorm/tree/main/docs/transaction_ja.md)
- [Mock](https://github.com/champon1020/gsorm/tree/main/docs/droptable_ja.md)
  - [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#expect)
  - [Return](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#return)
  - [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#complete)

## インストール
```
go get -u github.com/champon1020/gsorm
```

## クイックスタート
```go
package main

import (
	"log"
	"time"

	"github.com/champon1020/gsorm"
	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	ID        int
	Name      string    `gsorm:"first_name"`
	BirthDate time.Time `layout:"time.RFC3339"`
}

func main() {
	db, err := gsorm.New("mysql", "root:toor@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatalln(err)
	}

	// CREATE TABLE person (
	//   id         INT,
	//   first_name VARCHAR(128),
	//   birth_date DATETIME,
	// )
	err = gsorm.CreateTable(db, "person").Model(&Person{}).Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	// INSERT INTO person VALUES (10001, 'Taro', '2006-01-02T15:04:05Z00:00')
	err = gsorm.Insert(db, "person").
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
	err = gsorm.Select(db, "id", "first_name").From("person").Query(p)
	if err != nil {
		log.Fatalln(err)
	}

	// UPDATE person SET first_name='Jiro'
	err = gsorm.Update(db, "person", "first_name").Set("Jiro").Exec()
	if err != nil {
		log.Fatalln(err)
	}

	// DELETE FROM person
	err = gsorm.Delete(db).From("person").Exec()
	if err != nil {
		log.Fatalln(err)
	}

	// DROP TABLE person
	err = gsorm.DropTable(db, "person").Migrate()
	if err != nil {
		log.Fatalln(err)
	}
}
```
