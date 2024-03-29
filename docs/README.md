# Docs
This is the gsorm document.

[Japanese](https://github.com/champon1020/gsorm/tree/main/docs/README_ja.md)

## Overview
- [Introduction](https://github.com/champon1020/gsorm/tree/main/docs/introduction.md)
- [Select](https://github.com/champon1020/gsorm/tree/main/docs/select.md)
  - [From](https://github.com/champon1020/gsorm/tree/main/docs/select.md#from)
  - [Join](https://github.com/champon1020/gsorm/tree/main/docs/select.md#join)
  - [LeftJoin](https://github.com/champon1020/gsorm/tree/main/docs/select.md#leftjoin)
  - [RightJoin](https://github.com/champon1020/gsorm/tree/main/docs/select.md#rightjoin)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/select.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/select.md#and)
  - [Or](https://github.com/champon1020/gsorm/tree/main/docs/select.md#or)
  - [Group By](https://github.com/champon1020/gsorm/tree/main/docs/select.md#groupby)
  - [Having](https://github.com/champon1020/gsorm/tree/main/docs/select.md#having)
  - [Union](https://github.com/champon1020/gsorm/tree/main/docs/select.md#union)
  - [Union](https://github.com/champon1020/gsorm/tree/main/docs/select.md#unionall)
  - [Order By](https://github.com/champon1020/gsorm/tree/main/docs/select.md#orderby)
  - [Limit](https://github.com/champon1020/gsorm/tree/main/docs/select.md#limit)
  - [Offset](https://github.com/champon1020/gsorm/tree/main/docs/select.md#offset)
- [Function Query](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md)
  - [Count](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md#count)
  - [Sum](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md#sum)
  - [Avg](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md#avg)
  - [Max](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md#max)
  - [Min](https://github.com/champon1020/gsorm/tree/main/docs/fnquery.md#min)
- [Insert](https://github.com/champon1020/gsorm/tree/main/docs/insert.md)
  - [Values](https://github.com/champon1020/gsorm/tree/main/docs/insert.md#values)
  - [Select](https://github.com/champon1020/gsorm/tree/main/docs/insert.md#select)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/insert.md#model)
- [Update](https://github.com/champon1020/gsorm/tree/main/docs/update.md)
  - [Set](https://github.com/champon1020/gsorm/tree/main/docs/update.md#set)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/update.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/update.md#and)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/update.md#or)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/update.md#model)
- [Delete](https://github.com/champon1020/gsorm/tree/main/docs/delete.md)
  - [From](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#from)
  - [Where](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#where)
  - [And](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#and)
  - [Or](https://github.com/champon1020/gsorm/tree/main/docs/delete.md#or)
- [CreateDB](https://github.com/champon1020/gsorm/tree/main/docs/createdb.md)
- [CreateIndex](https://github.com/champon1020/gsorm/tree/main/docs/createindex.md)
- [CreateTable](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md)
  - [Column](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#column)
    - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#notnull)
    - [Default](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#default)
  - [Cons](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#cons)
    - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#unique)
    - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#primary)
    - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#foreign)
      - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#ref)
  - [Model](https://github.com/champon1020/gsorm/tree/main/docs/createtable.md#model)
- [AlterTable](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md)
  - [Rename](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#rename)
  - [AddColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcolumn)
    - [NotNull](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#notnull)
    - [Default](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#default)
  - [DropColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#dropcolumn)
  - [RenameColumn](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#renamecolumn)
  - [AddCons](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#addcons)
    - [Unique](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#unique)
    - [Primary](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#primary)
    - [Foreign](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#foreign)
      - [Ref](https://github.com/champon1020/gsorm/tree/main/docs/altertable.md#ref)
- [DropDB](https://github.com/champon1020/gsorm/tree/main/docs/dropdb.md)
- [DropTable](https://github.com/champon1020/gsorm/tree/main/docs/droptable.md)
- [Raw](https://github.com/champon1020/gsorm/tree/main/docs/raw.md)
  - [RawClause](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawclause)
  - [RawStmt](https://github.com/champon1020/gsorm/tree/main/docs/raw.md#rawstmt)
- [Model](https://github.com/champon1020/gsorm/tree/main/docs/model.md)
  - [Type](https://github.com/champon1020/gsorm/tree/main/docs/model.md#type)
  - [Tag](https://github.com/champon1020/gsorm/tree/main/docs/model.md#tag)
- [Connection](https://github.com/champon1020/gsorm/tree/main/docs/connection.md)
  - [DB](https://github.com/champon1020/gsorm/tree/main/docs/connection.md#db)
  - [Tx](https://github.com/champon1020/gsorm/tree/main/docs/connection.md#tx)
- [Mock](https://github.com/champon1020/gsorm/tree/main/docs/mock.md)
  - [MockDB](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdb)
    - [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpect)
    - [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpectreturn)
    - [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbcomplete)
    - [ExpectBegin](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mockdbexpectbegin)
  - [MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktx)
    - [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpect)
    - [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxexpectreturn)
    - [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxcomplete)
    - [ExpectCommit](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxcommit)
    - [ExpectRollback](https://github.com/champon1020/gsorm/tree/main/docs/mock.md#mocktxrollback)
