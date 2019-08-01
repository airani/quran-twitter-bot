package quran

import "testing"

func TestAya_IsSajdaObligatory(t *testing.T) {
	q := New()
	s := q.Sura(32)
	tests := []struct {
		name string
		aya  Aya
		want bool
	}{
		{
			name: "Aya is obligatory",
			aya:  s.Aya(15),
			want: true,
		},
		{
			name: "Aya is not obligatory",
			aya:  s.Aya(14),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Aya{
				Index:     tt.aya.Index,
				Text:      tt.aya.Text,
				Bismillah: tt.aya.Bismillah,
				Sajda:     tt.aya.Sajda,
			}
			if got := a.IsSajdaObligatory(); got != tt.want {
				t.Errorf("Aya.IsSajdaObligatory() = %v, want %v", got, tt.want)
			}
		})
	}
}
