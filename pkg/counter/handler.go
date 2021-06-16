package counter

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

func Handler(db *database.DB, operation string, args []string) (counter *[]entity.Counter, err error) {
	svc := NewService(NewRepository(db))

	switch operation {
	case "new":
		counter = entity.NewCounters(args)
		err = svc.Create(counter)

	case "list":
		counter = entity.NewCounters(args)
		err = svc.Read(counter)

	case "edit":
		counter, err = svc.Update(MakeMapUpdate(args))

	case "remove":
		counter = entity.NewCounters(args)
		if err = svc.Read(counter); err != nil {
			return
		}
		err = svc.Delete(counter)

	default:
		err = fmt.Errorf("invalid handler operation")
	}

	return counter, err
}
