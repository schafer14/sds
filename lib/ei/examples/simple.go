package examples

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

func Simple() int {

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
		log.Print(err)
		return 1
	}

	log.Print(data)
	return 0

}
