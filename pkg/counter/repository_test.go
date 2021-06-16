package counter

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
	want []entity.Counter
	got  []entity.Counter
}

func new_testHelper(givCounter *entity.Counter, testRepo Repository) (h *testHelper) {
	var (
		want = *givCounter
		got  = *givCounter
	)
	return &testHelper{
		repo: testRepo,
		want: []entity.Counter{want},
		got:  []entity.Counter{got},
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

func Test_repositoryCounter(t *testing.T) {
	const _testDB = "test_counter.db"
	var (
		counter = &entity.Counter{Name: "test"}
		repo    = test_NewRepository(t, _testDB)
		th      = new_testHelper(counter, repo)
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

func Test_repositoryCounterGetAll(t *testing.T) {
	const _testDB = "test_counter_getall.db"
	var (
		counters = []entity.Counter{}
		nTests   = 5
		repo     = test_NewRepository(t, _testDB)
	)
	defer func() { assert.NoError(t, os.Remove(_testDB)) }()

	for i := 0; i < nTests; i++ {
		counters = append(counters, *entity.NewCounter(fmt.Sprintf("test %d", i)))
	}

	if err := repo.Create(&counters); !assert.NoError(t, err) {
		t.FailNow()
	}

	for _, counter := range counters {
		assert.NotNil(t, counter.ID)
		assert.NotNil(t, counter.CreatedAt)
	}

	all := []entity.Counter{}
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
