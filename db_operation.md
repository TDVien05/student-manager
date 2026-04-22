# Go database/sql: Exec vs Query vs QueryRow

This document explains when to use `Exec`, `Query`, and `QueryRow` in Go's `database/sql` package.

---

# 🧠 Overview

When working with SQL in Go, you always choose one of three main methods depending on the result type:

| Method     | Use case                                  | Returns             |
| ---------- | ----------------------------------------- | ------------------- |
| `Exec`     | INSERT, UPDATE, DELETE (no rows returned) | `sql.Result, error` |
| `Query`    | SELECT multiple rows                      | `*sql.Rows, error`  |
| `QueryRow` | SELECT single row                         | `*sql.Row`          |

---

# 🚀 1. Exec

## 👉 When to use

Use `Exec` when the SQL does NOT return data.

### Common operations:

* INSERT
* UPDATE
* DELETE

## Example

```go
result, err := db.Exec("DELETE FROM students WHERE id = ?", id)
```

## Why use it?

Because these queries only modify data, they don't return rows.

## Getting affected rows

```go
rowsAffected, _ := result.RowsAffected()
```

---

# 📊 2. Query

## 👉 When to use

Use `Query` when SQL returns MULTIPLE rows.

### Common operations:

* SELECT all students
* Search with LIKE
* Filtered lists

## Example

```go
rows, err := db.Query("SELECT id, name FROM students WHERE name LIKE ?", "%vien%")
```

## Processing results

```go
for rows.Next() {
    var s Student
    rows.Scan(&s.ID, &s.Name)
}
```

## Why use it?

Because the result is a table (0..n rows).

---

# 👤 3. QueryRow

## 👉 When to use

Use `QueryRow` when SQL returns ONLY ONE row.

### Common operations:

* Get by ID
* Get by unique email

## Example

```go
row := db.QueryRow("SELECT id, name FROM students WHERE id = ?", id)

var s Student
err := row.Scan(&s.ID, &s.Name)
```

## Why use it?

Because only one record is expected.

---

# ⚠️ Important Rules

## ❌ Don't use QueryRow for multiple results

Example:

```sql
SELECT * FROM students WHERE name LIKE '%a%'
```

This may return many rows → use `Query` instead.

---

# 🧩 Quick Memory Trick

* `Exec` → "Do action, no result"
* `Query` → "Give me many rows"
* `QueryRow` → "Give me one row"

---

# 🏗️ Real Project Mapping

| Feature           | Method   |
| ----------------- | -------- |
| Create student    | Exec     |
| Update student    | Exec     |
| Delete student    | Exec     |
| Get student by ID | QueryRow |
| Search students   | Query    |
| List students     | Query    |

---

# 💡 Summary

Choosing the correct method is critical:

* Use `Exec` for write operations
* Use `Query` for lists
* Use `QueryRow` for single records

---

# 🎯 End

Once you understand this, working with Go SQL becomes very predictable and clean.
