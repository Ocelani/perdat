package fact

import (
	"fmt"

	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
)

func Handler(db *database.DB, operation string, args []string) (facts *[]entity.Fact, err error) {
	svc := NewService(NewRepository(db))

	switch operation {
	case "new":
		facts = entity.NewFacts(args)
		err = svc.Create(facts)

	case "list":
		facts = entity.NewFacts(args)
		err = svc.Read(facts)

	case "edit":
		facts, err = svc.Update(MakeMapUpdate(args))

	case "remove":
		facts = entity.NewFacts(args)
		if err = svc.Read(facts); err != nil {
			return
		}
		err = svc.Delete(facts)

	default:
		err = fmt.Errorf("invalid handler operation")
	}

	return facts, err
}
