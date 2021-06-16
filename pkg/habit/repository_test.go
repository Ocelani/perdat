package habit

import (
	"fmt"
	"os"
	"testing"

	"github.com/Ocelani/perdat/pkg/common/database"
	"github.com/Ocelani/perdat/pkg/entity"
	"github.com/stretchr/testify/assert"
)

type testHelper struct {
	repo Repository
	want []entity.Habit
	got  []entity.Habit
}

func new_testHelper(givHabit *entity.Habit, testRepo Repository) (h *testHelper) {
	var (
		want = *givHabit
		got  = *givHabit
	)
	return &testHelper{
		repo: testRepo,
		want: []entity.Habit{want},
		got:  []entity.Habit{got},
	}
}

func (h *testHelper) test(t *testing.T) {
	assert.NotNil(t, h.got)
	assert.Len(t, h.got, 1)
	assert.Equal(t, h.want[0].Name, h.got[0].Name)
	assert.NotEqual(t, fmt.Sprintf("%p", h.want), fmt.Sprintf("%p", h.got))
}

func test_NewRepository(t *testing.T, testDB string) Repository {
	var (
		want1 *Repository
		want2 *repo
		got   = NewRepository(database.SQLiteOpen(testDB, nil))
	)
	if !assert.NotNil(t, got) ||
		!assert.IsType(t, want2, got) ||
		!assert.Implements(t, want1, got) {
		t.FailNow()
	}
	return got
}

func Test_repositoryHabit(t *testing.T) {
	const _testDB = "test_habit.db"
	var (
		habit = &entity.Habit{Name: "test"}
		repo  = test_NewRepository(t, _testDB)
		th    = new_testHelper(habit, repo)
	)
	defer func() { assert.NoError(t, os.Remove(_testDB)) }()

	// Create
	err := th.repo.Create(&th.got)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	t.Run("Create", th.test)

	// Read
	err = th.repo.Read(&th.got)
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	t.Run("Read", th.test)

	// Update
	giv := "update name"
	th.got[0].Name = giv
	th.want[0].Name = giv
	err = th.repo.Update(&th.got[0])
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	t.Run("Update", th.test)

	// Delete
	err = th.repo.Delete(&th.got[0])
	if !assert.NoError(t, err) {
		t.FailNow()
	}
	t.Run("Delete", th.test)
}

func Test_repositoryHabitGetAll(t *testing.T) {
	const _testDB = "test_habit_getall.db"
	var (
		habits = []entity.Habit{}
		nTests = 5
		repo   = test_NewRepository(t, _testDB)
	)
	defer func() { assert.NoError(t, os.Remove(_testDB)) }()

	for i := 0; i < nTests; i++ {
		habits = append(habits, *entity.NewHabit(fmt.Sprintf("test %d", i)))
	}

	if err := repo.Create(&habits); !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, habit := range habits {
		assert.NotNil(t, habit.ID)
		assert.NotNil(t, habit.CreatedAt)
	}

	all := []entity.Habit{}
	if err := repo.Read(&all); !assert.NoError(t, err) {
		t.FailNow()
	}

	if !assert.Len(t, all, nTests) {
		t.FailNow()
	}

	for _, tAll := range all {
		t.Log("aa")
		assert.NotNil(t, tAll.ID)
		assert.NotNil(t, tAll.CreatedAt)
	}
}
