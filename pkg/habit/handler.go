package habit

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

func Handler(db *database.DB, operation string, args []string) (habits *[]entity.Habit, err error) {
	svc := NewService(NewRepository(db))

	switch operation {
	case "new":
		habits = entity.NewHabits(args)
		err = svc.Create(habits)

	case "list":
		habits = entity.NewHabits(args)
		err = svc.Read(habits)

	case "edit":
		habits, err = svc.Update(MakeMapUpdate(args))

	case "remove":
		habits = entity.NewHabits(args)
		if err = svc.Read(habits); err != nil {
			return
		}
		err = svc.Delete(habits)

	default:
		err = fmt.Errorf("invalid handler operation")
	}

	return habits, err
}
