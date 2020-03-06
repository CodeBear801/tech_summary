# How Spark SQL works by a simple query

## Case:
In the university, generate a summary of weighted mean scores for each department

```sql
SELECT dept, avg(math_score * 1.2) + avg(eng_score * 0.8) FROM students
GROUP BY dept;
```

## Query analysis
The purpose of parser(Lexer + Parser) is transforming a streaming of string to a group of tokens, then based on grammar generate a abstract syntax tree(AST).  


```js
TableScan(students)
-> Project(dept, avg(math_score * 1.2) + avg(eng_score * 0.8))
-> TableSink
```
The result will be generated as a unresolved logical plan.

More info:
- [antlr4](https://github.com/antlr/antlr4)
- [spark/sql/catalyst/parser/SqlBase.g4](https://github.com/apache/spark/blob/master/sql/catalyst/src/main/antlr4/org/apache/spark/sql/catalyst/parser/SqlBase.g4) generates sqlbaselexer and sqlbaseparser
-  [spark/sql/catalyst/parser/AstBuilder.scala](https://github.com/apache/spark/blob/master/sql/catalyst/src/main/scala/org/apache/spark/sql/catalyst/parser/AstBuilder.scala) convert parsetree to logical plan


## Analyzer
After previous step, we already have a basic structure, but Spark don't understand the keyword in tables, and have no idea of `sum, select, join, where`.  Now need schema catalog to represent those tokens.  

Unresolved logical plan missing following meta data information: table schema and function info.  Table schema contains table's column(name, type), physical address, format, how to retrieve
Function info is functions' signature and location of related class.

```js
TableScan(students=>dept:String, eng_score:double, math_score:double)
->Project(dept, math_score * 1.2:expr1, eng_score * 0.8:expr2)
->Aggregate(avg(expr1):expr3, avg(expr2):expr4, GROUP:dept)
->Project(dept, expr3+expr4:avg_result)
->TableSink(dept, avg_result->Client)
```
The result is an un-optimized logical plan

More info
- [spark/sql/catalyst/analysis/Analyzer.scala](https://github.com/apache/spark/blob/8d5ef2f766166cce3cc7a15a98ec016050ede4d8/sql/catalyst/src/main/scala/org/apache/spark/sql/catalyst/analysis/Analyzer.scala#L201)


