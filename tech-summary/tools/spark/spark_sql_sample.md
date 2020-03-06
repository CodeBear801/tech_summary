# How Spark SQL works by a simple query

## Case:
In the university, generate a summary of weighted mean scores for each department


```sql
SELECT dept, avg(math_score * 1.2) + avg(eng_score * 0.8) FROM students
GROUP BY dept;
```

