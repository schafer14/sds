# Editor Input

When I build service I find the easiest way to interact with them is through
a command line tool. I don't like leaving the terminal to visit swagger docs
to figure out what inputs I need to provide. The problem with a traditional 
CLI is it is hard to provide more that three or four args without the cognitive
overhead required to use the system exploding.

This library is my idea about the easiest way to provide use data without 
leaving the terminal or thinking.

What happens is you provide a default object to the editor function. This will
then get rendered into a temporary file as JSON. The users will have their
text editor of choice opened to that file. When the editor is closed the file
will be read back, validated, and the object you provided will be updated with
the users input.

## Example

```go
package main

import (
	"context"
	"log"

	"github.com/schafer14/sds/lib/ei"
)

type SimpleData struct {
	Name  string
	Email string
	Phone []struct {
		Type     string
		AreaCode string
		Number   string
	}
}

func main() {

	ctx := context.Background()
	data := SimpleData{
		Name:  "",
		Email: "",
		Phone: []struct {
			Type     string
			AreaCode string
			Number   string
		}{
			{
				Type:     "home",
				AreaCode: "+61",
				Number:   "000 000 000",
			},
		},
	}

	err := ei.Fetch(ctx, &data)
	if err != nil {
		log.Panic(err)
	}

	log.Print(data)
}
```
