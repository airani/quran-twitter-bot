package quran

import (
	"reflect"
	"testing"
)

func TestQuran_Sura(t *testing.T) {
	q := New()
	type fields struct {
		Suras []Sura
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  Sura
	}{
		{
			name: "Get first Sura",
			fields: fields{
				Suras: q.Suras,
			},
			args: args{
				n: 1,
			},
			wantS: q.Suras[0],
		},
		{
			name: "Get Sura number 0",
			fields: fields{
				Suras: q.Suras,
			},
			args: args{
				n: 0,
			},
			wantS: Sura{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Quran{
				Suras: tt.fields.Suras,
			}
			if gotS := q.Sura(tt.args.n); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("Quran.Sura() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
