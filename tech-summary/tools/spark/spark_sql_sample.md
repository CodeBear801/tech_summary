# How Spark SQL works by a simple query

## Case:
In the university, generate a summary of weighted mean scores for each department

```sql
SELECT dept, avg(math_score * 1.2) + avg(eng_score * 0.8) FROM students
GROUP BY dept;
```

## Query analysis
The purpose of parser(Lexer + Parser) is transforming a streaming of string to a group of tokens, then based on grammar generate a abstract syntax tree(AST).  

More info:
- [antlr4](https://github.com/antlr/antlr4)
- [spark/sql/catalyst/parser/SqlBase.g4](https://github.com/apache/spark/blob/master/sql/catalyst/src/main/antlr4/org/apache/spark/sql/catalyst/parser/SqlBase.g4) generates sqlbaselexer and sqlbaseparser
-  [spark/sql/catalyst/parser/AstBuilder.scala](https://github.com/apache/spark/blob/master/sql/catalyst/src/main/scala/org/apache/spark/sql/catalyst/parser/AstBuilder.scala) convert parsetree to logical plan

```js
TableScan(students)
-> Project(dept, avg(math_score * 1.2) + avg(eng_score * 0.8))
-> TableSink
```
The result will be generated as a unresolved logical plan.


