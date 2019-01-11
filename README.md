SelectFields
=====================

## Usage
```go
type Item struct {
	ID             int           `db:"id"`
	Title          string        `db:"title"`
	Content        string        `db:"content"`
	Date           string        `db:"-"`
}

func main() {
	var item = new(Item)
	fmt.Println("1.", dbhelper.SelectFields(item, ""))
	fmt.Println("2.", dbhelper.SelectFields(item, "i"))

	//Output
	//1. id,title,content
	//2. i.id.i.title,i.content
}
```