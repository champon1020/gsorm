# Docs
gsormのドキュメントです．

[English](https://github.com/champon1020/gsorm/tree/main/docs/README.md)

## Overview
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
  - [Type](https://github.com/champon1020/gsorm/tree/main/docs/model_ja.md#type)
  - [Tag](https://github.com/champon1020/gsorm/tree/main/docs/model_ja.md#tag)
- [Connection](https://github.com/champon1020/gsorm/tree/main/docs/connection_ja.md)
  - [DB](https://github.com/champon1020/gsorm/tree/main/docs/connection_ja.md#db)
  - [Tx](https://github.com/champon1020/gsorm/tree/main/docs/connection_ja.md#tx)
- [Mock](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md)
  - [MockDB](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mockdb)
    - [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mockdbexpect)
    - [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mockdbexpectreturn)
    - [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mockdbcomplete)
    - [ExpectBegin](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mockdbexpectbegin)
  - [MockTx](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktx)
    - [Expect](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktxexpect)
    - [ExpectWithReturn](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktxexpectreturn)
    - [Complete](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktxcomplete)
    - [ExpectCommit](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktxcommit)
    - [ExpectRollback](https://github.com/champon1020/gsorm/tree/main/docs/mock_ja.md#mocktxrollback)
