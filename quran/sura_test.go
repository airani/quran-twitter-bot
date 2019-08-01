package quran

import (
	"reflect"
	"testing"
)

func TestSura_Aya(t *testing.T) {
	q := New()
	s := q.Sura(1)
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields Sura
		args   args
		wantA  Aya
	}{
		{
			fields: s,
			args: args{
				n: 1,
			},
			wantA: s.Ayas[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sura{
				Index: tt.fields.Index,
				Name:  tt.fields.Name,
				Ayas:  tt.fields.Ayas,
			}
			if gotA := s.Aya(tt.args.n); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("Sura.Aya() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}
