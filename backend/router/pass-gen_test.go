package router

import "testing"

func Test_convertStrToNumbers(t *testing.T) {
	type args struct {
		lowerC   string
		upperC   string
		num      string
		specialC string
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		want1   int8
		want2   int8
		want3   int8
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				lowerC:   "5",
				upperC:   "5",
				num:      "5",
				specialC: "5",
			},
			want:    5,
			want1:   5,
			want2:   5,
			want3:   5,
			wantErr: false,
		},
		{
			name: "with error",
			args: args{
				lowerC:   "5",
				upperC:   "s",
				num:      "5",
				specialC: "5",
			},
			want:    0,
			want1:   0,
			want2:   0,
			want3:   0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, got3, err := convertStrToNumbers(tt.args.lowerC, tt.args.upperC, tt.args.num, tt.args.specialC)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertStrToNumbers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertStrToNumbers() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("convertStrToNumbers() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("convertStrToNumbers() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("convertStrToNumbers() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}

func Test_validateCountChars(t *testing.T) {
	type args struct {
		lowerC   int8
		upperC   int8
		num      int8
		specialC int8
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				lowerC:   5,
				upperC:   5,
				num:      5,
				specialC: 5,
			},
			wantErr: false,
		},
		{
			name: "with error",
			args: args{
				lowerC:   5,
				upperC:   5,
				num:      10,
				specialC: 5,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateCountChars(tt.args.lowerC, tt.args.upperC, tt.args.num, tt.args.specialC); (err != nil) != tt.wantErr {
				t.Errorf("validateCountChars() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
