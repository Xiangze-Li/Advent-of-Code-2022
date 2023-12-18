package util

import (
	"fmt"
	"testing"
)

func TestToBalancedQuinary(t *testing.T) {
	tests := []struct {
		args int64
		want string
	}{
		{args: 0, want: "0"},
		{args: 1747, want: "1=-0-2"},
		{args: -1747, want: "-2101="},
		{args: 906, want: "12111"},
		{args: -906, want: "-=---"},
		{args: 198, want: "2=0="},
		{args: -198, want: "=202"},
		{args: 11, want: "21"},
		{args: -11, want: "=-"},
		{args: 201, want: "2=01"},
		{args: -201, want: "=20-"},
		{args: 31, want: "111"},
		{args: -31, want: "---"},
		{args: 1257, want: "20012"},
		{args: -1257, want: "=00-="},
		{args: 32, want: "112"},
		{args: -32, want: "--="},
		{args: 353, want: "1=-1="},
		{args: -353, want: "-21-2"},
		{args: 107, want: "1-12"},
		{args: -107, want: "-1-="},
		{args: 7, want: "12"},
		{args: -7, want: "-="},
		{args: 3, want: "1="},
		{args: -3, want: "-2"},
		{args: 37, want: "122"},
		{args: -37, want: "-=="}}
	for idx, tt := range tests {
		t.Run(fmt.Sprint("case ", idx), func(t *testing.T) {
			if got := ToBalancedQuinary(tt.args); got != tt.want {
				t.Errorf("ToBalancedQuinary() = %q, want %q(%d)", got, tt.want, tt.args)
			}
			negLine := ToBalancedQuinary(-tt.args)
			t.Log(negLine)
		})
	}
}

func TestFromBalancedQuinary(t *testing.T) {
	tests := []struct {
		args    string
		want    int64
		wantErr bool
	}{
		{args: "0", want: 0},
		{args: "1=-0-2", want: 1747},
		{args: "-2101=", want: -1747},
		{args: "12111", want: 906},
		{args: "-=---", want: -906},
		{args: "2=0=", want: 198},
		{args: "=202", want: -198},
		{args: "21", want: 11},
		{args: "=-", want: -11},
		{args: "2=01", want: 201},
		{args: "=20-", want: -201},
		{args: "111", want: 31},
		{args: "---", want: -31},
		{args: "20012", want: 1257},
		{args: "=00-=", want: -1257},
		{args: "112", want: 32},
		{args: "--=", want: -32},
		{args: "1=-1=", want: 353},
		{args: "-21-2", want: -353},
		{args: "1-12", want: 107},
		{args: "-1-=", want: -107},
		{args: "12", want: 7},
		{args: "-=", want: -7},
		{args: "1=", want: 3},
		{args: "-2", want: -3},
		{args: "122", want: 37},
		{args: "-==", want: -37},
		{args: "-==+", wantErr: true},
		{args: "abcdefg", wantErr: true},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprint("case ", idx), func(t *testing.T) {
			got, err := FromBalancedQuinary(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromBalancedQuinary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FromBalancedQuinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNegateBalancedQuinary(t *testing.T) {
	tests := []struct {
		args  string
		want  string
		value int64
	}{
		{args: "-22=-=1=2-2120-01--11201102=", want: "1==212-2=1=-=010-11--=0--0=2", value: -4009878327599108617},
		{args: "1=01200=10211-2-=-0122001=1", want: "-20-=002-0=--1=1210-==00-2-", value: 910724980886834391},
		{args: "-2-02=0121=12101==210-00=2-0", want: "1=10=20-=-2-=-0-22=-01002=10", value: -4749162072622672080},
		{args: "1-2=00=0=--2-11110-==01==210", want: "-1=20020211=1----01220-22=-0", value: 6436305519088907930},
		{args: "10=2-10202001-1==-20-112-=0", want: "-02=1-0=0=00-1-221=01--=120", value: 1392881016971472715},
		{args: "=2=0=220-=2==2-=-2-0=100", want: "2=202==012=22=121=102-00", value: -20056193196143975},
		{args: "1---2022=20=---22-2=000-120=", want: "-111=0==2=02111==1=20001-=02", value: 5627792015960155798},
		{args: "1020--=1---0--=2=211-=11=--1", want: "-0=0112-1110112=2=--12--211-", value: 8031439933273722221},
		{args: "-01=20010-21=2=2=01=-=-=212", want: "10-2=00-01=-2=2=20-21212=-=", value: -1449566324872835193},
		{args: "110=02==2-0210022021121-1012", want: "--020=22=10=-00==0=--=-1-0-=", value: 8825146057719252632},
	}
	for idx, tt := range tests {
		t.Run(fmt.Sprint("case ", idx), func(t *testing.T) {
			if got := NegateBalancedQuinary(tt.args); got != tt.want {
				t.Errorf("NegateBalancedQuinary() = %v, want %v", got, tt.want)
			}
			if got := NegateBalancedQuinary(tt.want); got != tt.args {
				t.Errorf("NegateBalancedQuinary() = %v, want %v", got, tt.args)
			}
		})
	}
}
