package code068

import (
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test1",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name: "test2",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name: "test3",
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "compute.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  compute.  Art  is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			name: "test4",
			args: args{
				words:    []string{"Give", "me", "my", "Romeo;", "and,", "when", "he", "shall", "die,", "Take", "him", "and", "cut", "him", "out", "in", "little", "stars,", "And", "he", "will", "make", "the", "face", "of", "heaven", "so", "fine", "That", "all", "the", "world", "will", "be", "in", "love", "with", "night", "And", "pay", "no", "worship", "to", "the", "garish", "sun."},
				maxWidth: 25,
			},
			want: []string{
				"Give  me  my  Romeo; and,",
				"when  he  shall die, Take",
				"him  and  cut  him out in",
				"little stars, And he will",
				"make  the  face of heaven",
				"so   fine  That  all  the",
				"world  will  be  in  love",
				"with  night  And  pay  no",
				"worship   to  the  garish",
				"sun.                     ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fullJustify() = %v, want %v", got, tt.want)
			}
		})
	}
}
