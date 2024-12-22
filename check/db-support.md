To make the expected result come directly from the JSON document, the structure can be adjusted to ensure that the expected output for the SQL query is included directly under each test case. This way, the backend can retrieve the expected result directly from the JSON when validating the query submitted by the user.

### Updated JSON Structure

```json
{
  "id": "e4a2fb80-3d6f-4c58-8aab-2e5c95c10d3f",
  "title": "Find Top 5 Actors in Movies",
  "description": "Write a query to find the top 5 actors who have starred in the most movies.",
  "question_text": "Write an SQL query that selects the top 5 actors who have starred in the most movies from the Sakila database.",
  "db": "sakila",
  "table": "actor",
  "language": "sql",
  "difficulty": "medium",
  "tags": ["actors", "movies", "top 5", "join", "group by"],
  "query_template": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
  "testcases": [
    {
      "input": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
      "expected_result": {
        "columns": ["actor_id", "first_name", "last_name", "film_count"],
        "rows": [
          [1, "PENELOPE", "GUINESS", 10],
          [2, "NICK", "WAHLBERG", 8],
          [3, "ED", "CHASE", 8],
          [4, "JENNIFER", "DAVIS", 7],
          [5, "MATTHEW", "JOHANSSON", 6]
        ]
      }
    }
  ]
}
```

### Explanation of the Updated Structure

- **id**: A unique identifier for the question.
- **title**: The title of the question.
- **description**: A description of the question.
- **question_text**: The prompt that the user will see. It describes the task and how to approach it.
- **db**: The name of the database (`sakila` in this case).
- **table**: The name of the table that the query relates to (`actor` in this case).
- **language**: The type of language for the query (`sql`).
- **difficulty**: The difficulty level of the question.
- **tags**: Tags associated with the question to help categorize it.
- **query_template**: A reference SQL query that could be shown as a guideline.
- **testcases**: An array of test cases:
  - **input**: The SQL query submitted by the user.
  - **expected_result**: The expected result of the query in terms of columns and rows.

### Expected Result

- **columns**: The names of the columns that should be returned in the result set (e.g., `actor_id`, `first_name`, `last_name`, `film_count`).
- **rows**: The expected data in the result set for those columns. This is used for validation to check whether the user's query returns the expected data.

### Validation Process

1. **Submit Query**: The user submits an SQL query through the UI.
2. **Execute Query**: The backend (in Go) executes the user-submitted query against the **Sakila** database (MySQL/PostgreSQL).
3. **Compare Result**:
   - The backend fetches the `expected_result` from the JSON document (specifically under the `testcases` section).
   - It compares the actual result returned from the database with the expected result.
   - If the results match, the query is correct. If not, the backend flags it as incorrect.

### Example Go Backend Function for Validation

Here’s an updated example of how the backend can execute the query and validate the result:

```go
import (
	"database/sql"
	"fmt"
	"log"
	"github.com/go-sql-driver/mysql" // MySQL driver or other database drivers
)

// Struct to represent the expected result in JSON
type ExpectedResult struct {
	Columns []string           `json:"columns"`
	Rows    [][]interface{}    `json:"rows"`
}

// Function to validate the query against the Sakila database
func validateQuery(db *sql.DB, query string, expectedResult ExpectedResult) bool {
	// Execute the user's query
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error executing query:", err)
		return false
	}
	defer rows.Close()

	var result []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		log.Println("Error getting columns:", err)
		return false
	}

	// Scan each row into a map with column names as keys
	for rows.Next() {
		row := make(map[string]interface{})
		columnsPointers := make([]interface{}, len(columns))
		for i := range columnsPointers {
			columnsPointers[i] = &columnsPointers[i]
		}
		if err := rows.Scan(columnsPointers...); err != nil {
			log.Println("Error scanning row:", err)
			return false
		}

		for i, colName := range columns {
			row[colName] = columnsPointers[i]
		}
		result = append(result, row)
	}

	// Compare the result with the expected output
	return compareResults(result, expectedResult.Rows)
}

// Function to compare the actual and expected results
func compareResults(actual []map[string]interface{}, expected [][]interface{}) bool {
	// Check if actual and expected results match in terms of number of rows
	if len(actual) != len(expected) {
		return false
	}

	// Compare each row and column in actual and expected result
	for i := range actual {
		for j, column := range expected[i] {
			if actual[i][j] != column {
				return false
			}
		}
	}

	return true
}
```

### Conclusion

In this structure, the **expected result** comes directly from the JSON document in the `testcases` section. The backend then uses this expected result to validate the SQL query submitted by the user. The backend compares the actual query results with the expected ones and provides feedback accordingly.


To show the user how the table looks, you can include the schema of the table (e.g., column names and types) in the JSON document. This will allow you to present a clear view of the table structure to the user, making it easier for them to craft their SQL query.

### Updated JSON Structure with Table Schema

```json
{
  "id": "e4a2fb80-3d6f-4c58-8aab-2e5c95c10d3f",
  "title": "Find Top 5 Actors in Movies",
  "description": "Write a query to find the top 5 actors who have starred in the most movies.",
  "question_text": "Write an SQL query that selects the top 5 actors who have starred in the most movies from the Sakila database.",
  "db": "sakila",
  "table": "actor",
  "language": "sql",
  "difficulty": "medium",
  "tags": ["actors", "movies", "top 5", "join", "group by"],
  "table_schema": {
    "columns": [
      {
        "name": "actor_id",
        "type": "INT",
        "description": "Unique identifier for each actor."
      },
      {
        "name": "first_name",
        "type": "VARCHAR(45)",
        "description": "Actor's first name."
      },
      {
        "name": "last_name",
        "type": "VARCHAR(45)",
        "description": "Actor's last name."
      },
      {
        "name": "last_update",
        "type": "TIMESTAMP",
        "description": "Timestamp of the last update to the actor's record."
      }
    ],
    "sample_data": [
      [1, "PENELOPE", "GUINESS", "2021-01-01 10:00:00"],
      [2, "NICK", "WAHLBERG", "2021-01-01 10:00:00"],
      [3, "ED", "CHASE", "2021-01-01 10:00:00"],
      [4, "JENNIFER", "DAVIS", "2021-01-01 10:00:00"],
      [5, "MATTHEW", "JOHANSSON", "2021-01-01 10:00:00"]
    ]
  },
  "query_template": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
  "testcases": [
    {
      "input": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
      "expected_result": {
        "columns": ["actor_id", "first_name", "last_name", "film_count"],
        "rows": [
          [1, "PENELOPE", "GUINESS", 10],
          [2, "NICK", "WAHLBERG", 8],
          [3, "ED", "CHASE", 8],
          [4, "JENNIFER", "DAVIS", 7],
          [5, "MATTHEW", "JOHANSSON", 6]
        ]
      }
    }
  ]
}
```

### Explanation of the Changes

1. **table_schema**: 
   - **columns**: A list of column objects, where each object has:
     - `name`: The column name.
     - `type`: The type of the column (e.g., `INT`, `VARCHAR(45)`, `TIMESTAMP`).
     - `description`: A brief description of what the column represents.
   - **sample_data**: A small set of sample rows from the table to help the user visualize the data structure.

2. **query_template**: An example SQL query template that the user could refer to while formulating their own query.

3. **testcases**: Each test case includes:
   - **input**: The SQL query that the user submits.
   - **expected_result**: The expected result of the query, including the column names and rows.

### How This Helps the User

- **Table Schema**: Users can refer to the schema to understand what columns are available in the table, their data types, and what they represent. This helps them craft the correct SQL query.
  
- **Sample Data**: Providing some sample rows from the table allows users to better visualize the data in context and make more informed decisions when writing the query.

### Example UI Flow

1. **Table Display**: The UI can display the table schema and sample data in a format that resembles a database table (columns and rows). For example:

   | actor_id | first_name | last_name  | last_update          |
   |----------|------------|------------|----------------------|
   | 1        | PENELOPE   | GUINESS    | 2021-01-01 10:00:00  |
   | 2        | NICK       | WAHLBERG   | 2021-01-01 10:00:00  |
   | 3        | ED         | CHASE      | 2021-01-01 10:00:00  |
   | 4        | JENNIFER   | DAVIS      | 2021-01-01 10:00:00  |
   | 5        | MATTHEW    | JOHANSSON  | 2021-01-01 10:00:00  |

2. **Query Input**: Once the user sees the table, they can write their SQL query in an editor (e.g., a Monaco editor) based on this information.

3. **Submit & Validate**: The user submits the query, and the backend executes it on the database, comparing the result with the expected output from the `expected_result` field in the JSON.

### Benefits

- This approach provides a clear view of the data, which is crucial for the user to formulate their query.
- The backend can use the table schema and sample data to ensure the query is valid, while also providing the user with an understandable question context.


When sorting or ordering is involved in the query, the validation process can be a bit more nuanced. To handle this, we need to ensure that the expected results are sorted in the same way as the query and that the comparison considers the order of rows as well. Here's how we can approach this:

### Key Steps for Validation with Sorting or Ordering:

1. **Sorting in Query**: If the user's query includes an `ORDER BY` clause, the expected result should ideally reflect the sorting logic, so the query result matches the expected ordering when compared.

2. **Expected Result Should Be Sorted**: The expected result in the JSON should be provided in the correct sorted order, just as it would appear if the query is run with the `ORDER BY` clause.

3. **Handling Sorting Dynamically**: If the user has not included an `ORDER BY` clause, but the expected result relies on a specific order, we may need to compare the data ignoring the order of rows, which means checking the contents rather than row positions.

4. **Comparing Sorted Results**: During the validation, after the user's query is executed:
   - **If the query includes an `ORDER BY` clause**: Sort the results from both the database query and the expected result by the same column(s), and then compare them.
   - **If the query does not include an `ORDER BY` clause**: Ensure that the validation compares the results without considering the order of the rows. In this case, the results can be compared by ensuring that each row in the result exists in the expected set, regardless of the order.

### Example Scenario

Consider a case where a query is expected to return a list of actors sorted by the number of movies they have acted in. If the query includes an `ORDER BY` clause, the validation should account for this.

### Example JSON for Validation with Sorting

```json
{
  "id": "e4a2fb80-3d6f-4c58-8aab-2e5c95c10d3f",
  "title": "Find Top 5 Actors in Movies",
  "description": "Write a query to find the top 5 actors who have starred in the most movies.",
  "question_text": "Write an SQL query that selects the top 5 actors who have starred in the most movies from the Sakila database.",
  "db": "sakila",
  "table": "actor",
  "language": "sql",
  "difficulty": "medium",
  "tags": ["actors", "movies", "top 5", "join", "group by"],
  "table_schema": {
    "columns": [
      {
        "name": "actor_id",
        "type": "INT",
        "description": "Unique identifier for each actor."
      },
      {
        "name": "first_name",
        "type": "VARCHAR(45)",
        "description": "Actor's first name."
      },
      {
        "name": "last_name",
        "type": "VARCHAR(45)",
        "description": "Actor's last name."
      },
      {
        "name": "last_update",
        "type": "TIMESTAMP",
        "description": "Timestamp of the last update to the actor's record."
      }
    ],
    "sample_data": [
      [1, "PENELOPE", "GUINESS", "2021-01-01 10:00:00"],
      [2, "NICK", "WAHLBERG", "2021-01-01 10:00:00"],
      [3, "ED", "CHASE", "2021-01-01 10:00:00"],
      [4, "JENNIFER", "DAVIS", "2021-01-01 10:00:00"],
      [5, "MATTHEW", "JOHANSSON", "2021-01-01 10:00:00"]
    ]
  },
  "query_template": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
  "testcases": [
    {
      "input": "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;",
      "expected_result": {
        "columns": ["actor_id", "first_name", "last_name", "film_count"],
        "rows": [
          [1, "PENELOPE", "GUINESS", 15],
          [2, "NICK", "WAHLBERG", 12],
          [3, "ED", "CHASE", 10],
          [4, "JENNIFER", "DAVIS", 9],
          [5, "MATTHEW", "JOHANSSON", 8]
        ]
      }
    }
  ]
}
```

### Validation Logic

1. **Query Contains Sorting**: 
   - If the query contains an `ORDER BY` clause (like in the `query_template` above), the validation needs to compare the actual result from the database with the expected result after both have been sorted in the same order.

2. **Query Does Not Contain Sorting**: 
   - If the query does **not** contain an `ORDER BY` clause, the comparison must ensure the values match, but the order of the rows does not matter.
   - This can be achieved by converting both the result and the expected data into sets (ignoring order) and comparing them.

### Example of Validation with Sorting in Go (Backend)

```go
package main

import (
	"database/sql"
	"fmt"
	"sort"
	"log"
)

type TestResult struct {
	Columns []string
	Rows    [][]interface{}
}

func compareResults(expected, actual TestResult) bool {
	// Compare the number of columns
	if len(expected.Columns) != len(actual.Columns) {
		return false
	}

	// Sort both expected and actual results based on the first column (for sorting)
	sort.Slice(expected.Rows, func(i, j int) bool {
		return expected.Rows[i][0].(int) < expected.Rows[j][0].(int)
	})
	sort.Slice(actual.Rows, func(i, j int) bool {
		return actual.Rows[i][0].(int) < actual.Rows[j][0].(int)
	})

	// Compare rows and columns
	for i, row := range expected.Rows {
		for j, val := range row {
			if val != actual.Rows[i][j] {
				return false
			}
		}
	}
	return true
}

func runQuery(db *sql.DB, query string) (TestResult, error) {
	// Execute the query and fetch results
	rows, err := db.Query(query)
	if err != nil {
		return TestResult{}, err
	}
	defer rows.Close()

	// Assuming we have a fixed structure, we can iterate over rows
	var columns []string
	var resultRows [][]interface{}
	for rows.Next() {
		var actor_id int
		var first_name, last_name string
		var film_count int
		if err := rows.Scan(&actor_id, &first_name, &last_name, &film_count); err != nil {
			return TestResult{}, err
		}
		resultRows = append(resultRows, []interface{}{actor_id, first_name, last_name, film_count})
	}

	// Return results
	return TestResult{
		Columns: columns,
		Rows:    resultRows,
	}, nil
}

func main() {
	// Assuming db is already connected
	var db *sql.DB

	// Sample query
	query := "SELECT actor_id, first_name, last_name, COUNT(film_id) AS film_count FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id GROUP BY actor_id ORDER BY film_count DESC LIMIT 5;"

	// Run the query
	actualResult, err := runQuery(db, query)
	if err != nil {
		log.Fatal(err)
	}

	// Sample expected result (can be loaded from JSON)
	expectedResult := TestResult{
		Columns: []string{"actor_id", "first_name", "last_name", "film_count"},
		Rows: [][]interface{}{
			{1, "PENELOPE", "GUINESS", 15},
			{2, "NICK", "WAHLBERG", 12},
			{3, "ED", "CHASE", 10},
			{4, "JENNIFER", "DAVIS", 9},
			{5, "MATTHEW", "JOHANSSON", 8},
		},
	}

	// Validate results
	if compareResults(expectedResult, actualResult) {
		fmt.Println("Query is correct!")
	} else {
		fmt.Println("Query result does not match expected output.")
	}
}
```

### Explanation of the Code:

1. **Sorting**: The `compareResults` function sorts the expected and actual rows before comparing them. Sorting is based on the first column (assuming that column is unique).
   
2. **Validation**: After sorting, it compares each row and column value between the expected and actual result.

3. **No Sorting**: If sorting is not involved in the query, the comparison is made without sorting the rows, but rather ensuring that each result set contains the same rows, irrespective of their order.

### Conclusion:

For validation with sorting, the main tasks are:
- Ensure that the **expected result is sorted** according to the query’s `ORDER BY` clause.
- If no `ORDER BY` clause exists, ensure that the results match regardless of order.

By following this approach, you can handle queries with sorting and ensure that the user’s query is validated correctly against the expected results.


When a single query involves two or more tables (typically through `JOIN` operations), validation of the results becomes slightly more complex. However, the core validation principles remain the same: we need to ensure that the query returns the expected data, even when multiple tables are involved. Here's how you can handle validation in such scenarios:

### Key Points to Consider

1. **JOIN Operations**: The query might include various types of joins (e.g., `INNER JOIN`, `LEFT JOIN`, etc.) between two or more tables. The expected result must account for how these joins affect the data.

2. **Table Data Representation**: The expected results must reflect the structure of the join. For example, if you're joining two tables, the expected result should reflect all the columns from both tables in the correct order.

3. **Column Matching**: The query result will include columns from both tables, so the validation needs to check that each column's data matches the expected output across all tables involved.

4. **Sorting**: If sorting is involved (e.g., an `ORDER BY` clause), the expected result should also be sorted in the same way as the query.

### Approach for Validation

1. **Expected Result in JSON**: The expected result for a multi-table query should include all columns that come from the joined tables. Each row in the expected result should represent the data as it would appear after the join operation.

2. **Comparing Query Results**: The query's result should be compared to the expected result in terms of both content and order (if applicable). For multi-table joins, we need to make sure that:
   - The join is correctly represented in the expected result.
   - The data is correctly matched between the tables.
   - Sorting (if any) is considered.

3. **Handling Multiple Tables**: You should also ensure that the correct relationships between the tables are maintained in both the query and the expected result.

### Example Scenario

Suppose we have two tables: `actor` and `film_actor`, and we want to write a query that joins these tables to get the names of actors who have starred in a specific film. Here's an example of such a query:

```sql
SELECT a.actor_id, a.first_name, a.last_name, fa.film_id
FROM actor a
JOIN film_actor fa ON a.actor_id = fa.actor_id
WHERE fa.film_id = 1;
```

The query joins the `actor` and `film_actor` tables to retrieve actors who have starred in film `1`. We need to provide both the table schema and the expected results for this query.

### Example JSON for Validation with Multiple Tables

```json
{
  "id": "5f3dbb21-1234-4d95-8cdb-bd7b88fd0132",
  "title": "Find Actors for Film ID 1",
  "description": "Write an SQL query to find actors who starred in the film with ID 1.",
  "question_text": "Write an SQL query that selects the actor details who starred in film with ID 1 from the Sakila database.",
  "db": "sakila",
  "tables": [
    {
      "name": "actor",
      "columns": [
        {
          "name": "actor_id",
          "type": "INT",
          "description": "Unique identifier for each actor."
        },
        {
          "name": "first_name",
          "type": "VARCHAR(45)",
          "description": "Actor's first name."
        },
        {
          "name": "last_name",
          "type": "VARCHAR(45)",
          "description": "Actor's last name."
        }
      ]
    },
    {
      "name": "film_actor",
      "columns": [
        {
          "name": "actor_id",
          "type": "INT",
          "description": "Foreign key referencing actor_id in the actor table."
        },
        {
          "name": "film_id",
          "type": "INT",
          "description": "Foreign key referencing film_id in the film table."
        }
      ]
    }
  ],
  "query_template": "SELECT a.actor_id, a.first_name, a.last_name, fa.film_id FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id WHERE fa.film_id = 1;",
  "testcases": [
    {
      "input": "SELECT a.actor_id, a.first_name, a.last_name, fa.film_id FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id WHERE fa.film_id = 1;",
      "expected_result": {
        "columns": ["actor_id", "first_name", "last_name", "film_id"],
        "rows": [
          [1, "PENELOPE", "GUINESS", 1],
          [2, "NICK", "WAHLBERG", 1],
          [3, "ED", "CHASE", 1]
        ]
      }
    }
  ]
}
```

### Explanation of JSON:

1. **Tables and Columns**: The `tables` array specifies two tables: `actor` and `film_actor`. Each table includes a `columns` array that describes the columns and their data types.
   
2. **Query**: The `query_template` provides the SQL query that needs to be executed. It performs a join between `actor` and `film_actor` tables to get the list of actors who starred in film `1`.

3. **Expected Result**: The `expected_result` section includes the columns and rows that are expected to be returned when the query is executed. The `rows` array should match the result of the query after it has been executed on the database.

### Validation Logic for Multi-Table Queries:

1. **Join Validation**: 
   - Ensure that the data from both tables is included in the result.
   - The `actor_id` from the `actor` table should match the `actor_id` from the `film_actor` table.

2. **Column Matching**: 
   - The columns in the result should match the columns in the expected result in the correct order (unless the query does not involve sorting, in which case the rows should be compared regardless of order).

3. **Sorting**: 
   - If the query contains an `ORDER BY` clause, ensure that both the actual and expected results are sorted in the same way.
   - If the query does not contain `ORDER BY`, check if the data is valid regardless of the row order.

4. **Edge Case Handling**:
   - Ensure the query handles edge cases, such as no results (empty set) or unexpected data types.

### Example Go Code for Validation:

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
)

type TestResult struct {
	Columns []string
	Rows    [][]interface{}
}

func compareResults(expected, actual TestResult) bool {
	// Compare the number of columns
	if len(expected.Columns) != len(actual.Columns) {
		return false
	}

	// Sort both expected and actual results (if query involves sorting)
	sort.Slice(expected.Rows, func(i, j int) bool {
		return expected.Rows[i][0].(int) < expected.Rows[j][0].(int)
	})
	sort.Slice(actual.Rows, func(i, j int) bool {
		return actual.Rows[i][0].(int) < actual.Rows[j][0].(int)
	})

	// Compare rows and columns
	for i, row := range expected.Rows {
		for j, val := range row {
			if val != actual.Rows[i][j] {
				return false
			}
		}
	}
	return true
}

func runQuery(db *sql.DB, query string) (TestResult, error) {
	// Execute the query and fetch results
	rows, err := db.Query(query)
	if err != nil {
		return TestResult{}, err
	}
	defer rows.Close()

	// Assuming we have a fixed structure, we can iterate over rows
	var columns []string
	var resultRows [][]interface{}
	for rows.Next() {
		var actor_id int
		var first_name, last_name string
		var film_id int
		if err := rows.Scan(&actor_id, &first_name, &last_name, &film_id); err != nil {
			return TestResult{}, err
		}
		resultRows = append(resultRows, []interface{}{actor_id, first_name, last_name, film_id})
	}

	// Return results
	return TestResult{
		Columns: columns,
		Rows:    resultRows,
	}, nil
}

func main() {
	// Assuming db is already connected
	var db *sql.DB

	// Sample query
	query := "SELECT a.actor_id, a.first_name, a.last_name, fa.film_id FROM actor a JOIN film_actor fa ON a.actor_id = fa.actor_id WHERE fa.film_id = 1;"

	// Run the query
	actualResult, err := runQuery(db, query)
	if err != nil {
		log.Fatal(err)
	}

	// Sample expected result (can be loaded from JSON)
	expectedResult := TestResult{
		Columns: []string{"actor_id", "first_name", "last_name", "film_id"},
		Rows: [][]interface{}{
			{1, "PENELOPE", "GUINESS", 1},
			{2, "NICK", "WAHLBERG", 1},
			{3, "ED", "CHASE", 1},
		},
	}

	// Validate results
	if compareResults(expectedResult, actualResult) {
		fmt.Println("Query is correct!")
	} else {
		fmt.Println("Query result does not match expected output.")
	}
}
```

### Conclusion:

For multi-table queries, ensure that the expected results reflect the join logic correctly and include columns from all relevant tables. During validation:
- Handle the relationships between the tables.
- Ensure the query result is compared with the expected output, considering sorting if applicable.
- Use similar logic for comparing the results, whether the query involves multiple tables or just one.

{
  "id": "a1234bcd-567e-890f-ghi1-234567jkl890",
  "title": "Update User Age",
  "description": "Update the age of all users above 30 to 25 in the `users` table.",
  "database": "sakila",
  "difficulty": "medium",
  "tags": ["SQL", "UPDATE", "MySQL", "PostgreSQL"],
  "tables": [
    {
      "table_name": "users",
      "create_table_query": "CREATE TABLE users (id INT, name VARCHAR(50), age INT);",
      "columns": ["id", "name", "age"],
      "example_data": [
        { "id": 1, "name": "Alice", "age": 35 },
        { "id": 2, "name": "Bob", "age": 28 },
        { "id": 3, "name": "Charlie", "age": 40 },
        { "id": 4, "name": "David", "age": 22 }
      ],
      "restore_query": "DELETE FROM users; INSERT INTO users (id, name, age) VALUES (1, 'Alice', 35), (2, 'Bob', 28), (3, 'Charlie', 40), (4, 'David', 22);"
    },
    {
      "table_name": "orders",
      "create_table_query": "CREATE TABLE orders (order_id INT, user_id INT, amount DECIMAL(10,2), status VARCHAR(20));",
      "columns": ["order_id", "user_id", "amount", "status"],
      "example_data": [
        { "order_id": 101, "user_id": 1, "amount": 100, "status": "shipped" },
        { "order_id": 102, "user_id": 2, "amount": 200, "status": "pending" }
      ],
      "restore_query": "DELETE FROM orders; INSERT INTO orders (order_id, user_id, amount, status) VALUES (101, 1, 100, 'shipped'), (102, 2, 200, 'pending');"
    }
  ],
  "question": {
    "query_type": "UPDATE",
    "query_modifies_data": true,
    "question_text": "Update the age of all users above 30 to 25 in the `users` table.",
    "expected_result": [
      { "id": 1, "name": "Alice", "age": 25 },
      { "id": 2, "name": "Bob", "age": 28 },
      { "id": 3, "name": "Charlie", "age": 25 },
      { "id": 4, "name": "David", "age": 22 }
    ],
    "hints": [
      "You can use the SQL `UPDATE` statement.",
      "Check for users whose age is greater than 30."
    ],
    "expected_query": "UPDATE users SET age = 25 WHERE age > 30;"
  },
  "validation": {
    "strict_ordering": false,
    "ignore_case": true,
    "ignore_whitespace": true
  }
}


CREATE TABLE film (
    film_id VARCHAR(255) NOT NULL, -- Matches the insert for '1'
    title VARCHAR(255) NOT NULL, -- Title is a string and cannot be NULL
    description TEXT, -- Description is a long text
    release_year VARCHAR(4), -- Insert has a 4-digit string, so we use VARCHAR(4)
    language_id INTEGER NOT NULL, -- Insert uses 1, so this is an INTEGER
    original_language_id INTEGER, -- Can be NULL as per the insert
    rental_duration SMALLINT DEFAULT 3 NOT NULL, -- Default to 3 as per the original statement
    rental_rate NUMERIC(4,2) DEFAULT 4.99 NOT NULL, -- Default to 4.99
    length SMALLINT, -- Length can be NULL in the insert
    replacement_cost NUMERIC(5,2) DEFAULT 19.99 NOT NULL, -- Default to 19.99
    rating VARCHAR(5) DEFAULT 'G', -- Original type was mpaa_rating, but VARCHAR works
    special_features TEXT[], -- Array of text for 'Deleted Scenes,Behind the Scenes'
    last_update TIMESTAMP DEFAULT now() NOT NULL, -- Timestamp default is 'now()'
    fulltext TSVECTOR  -- Full-text search vector
); this is film table structure and you have actors table structure CREATE TABLE film_actor (
    actor_id integer NOT NULL,
    film_id integer NOT NULL,
    last_update timestamp without time zone DEFAULT now() NOT NULL
); and this is junction table film_actor now create questions based on aggregate function for given tables all should be select type