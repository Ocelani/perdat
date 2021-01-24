package test

import (
	"testing"
	"time"

	"github.com/Ocelani/perdat/pkg"
)

func TestFact_Create(t *testing.T) {
	type fields struct {
		ID        uint
		Name      string
		DateTime  time.Time
		CratedAt  time.Time
		UpdatedAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"Test 1",
			fields{
				Name:     "Name",
				DateTime: time.Now(),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &pkg.Fact{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				DateTime:  tt.fields.DateTime,
				CratedAt:  tt.fields.CratedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := f.Create(); (err != nil) != tt.wantErr {
				t.Errorf("Fact.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
