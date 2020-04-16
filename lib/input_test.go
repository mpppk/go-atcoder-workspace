package lib

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func Test_toLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]string
	}{
		{
			name: "can convert scanner to [][]string",
			input: `
					1
					2 3

					test hoge fuga
			`,
			want: [][]string{
				{"1"},
				{"2", "3"},
				{},
				{"test", "hoge", "fuga"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(TrimSpaceAndNewLineCodeAndTab(tt.input)))
			if got := toLines(scanner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toIntLine(t *testing.T) {
	type args struct {
		line []string
	}
	tests := []struct {
		name        string
		args        args
		wantIntLine []int64
		wantErr     bool
	}{
		{
			name: "can convert from string slice to int64 slice",
			args: args{
				line: []string{"1", "2", "3"},
			},
			wantIntLine: []int64{1, 2, 3},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIntLine, err := StringSliceToInt64Slice(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("toSpecificBitIntLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIntLine, tt.wantIntLine) {
				t.Errorf("toSpecificBitIntLine() = %v, want %v", gotIntLine, tt.wantIntLine)
			}
		})
	}
}

func TestInput_GetLine(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "can get specified index line",
			fields: fields{
				lines: [][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
					{"7", "8", "9"},
				},
			},
			args:    args{index: 0},
			want:    []string{"1", "2", "3"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetLine(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetColLine(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		colIndex int
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantNewLine []string
		wantErr     bool
	}{
		{
			name: "GetColLine",
			fields: fields{
				lines: [][]string{
					{"1", "2", "3"},
					{"4", "5", "6"},
				},
			},
			args: args{
				colIndex: 0,
			},
			wantNewLine: []string{"1", "4"},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLine, err := i.GetColLine(tt.args.colIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetColLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLine, tt.wantNewLine) {
				t.Errorf("Input.GetColLine() = %v, want %v", gotNewLine, tt.wantNewLine)
			}
		})
	}
}

func TestInput_MustGetFirstValue(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "MustGetFirstValue",
			fields: fields{
				lines: [][]string{
					{"foo", "bar", "piyo"},
					{"foo2", "bar2", "piyo2"},
				},
			},
			args: args{
				rowIndex: 0,
			},
			want: "foo",
		},
		{
			name: "1x1 column",
			fields: fields{
				lines: [][]string{
					{"foo"},
				},
			},
			args: args{
				rowIndex: 0,
			},
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			if got, err := i.GetFirstValue(tt.args.rowIndex); got != tt.want {
				if err != nil {
					t.Errorf("unexpected errror occurred: %v", err)
				}
				t.Errorf("Input.MustGetFirstValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLinesFromReader(t *testing.T) {
	type args struct {
		reader *bufio.Reader
	}
	tests := []struct {
		name      string
		args      args
		wantLines [][]string
		wantErr   bool
	}{
		{
			name: "toLinesFromReader",
			args: args{
				reader: bufio.NewReader(strings.NewReader("1 2 3\n4 5 6")),
			},
			wantLines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLines, err := toLinesFromReader(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("toLinesFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("toLinesFromReader() = %v, want %v", gotLines, tt.wantLines)
			}
		})
	}
}

func TestInput_GetLines(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		startRowIndex int
		endRowIndex   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "GetLines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			},
			args: args{
				startRowIndex: 0,
				endRowIndex:   1,
			},
			want:    [][]string{{"1", "2", "3"}},
			wantErr: false,
		},
		{
			name: "GetLines2",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			},
			args: args{
				startRowIndex: 1,
				endRowIndex:   3,
			},
			want:    [][]string{{"4", "5", "6"}, {"7", "8", "9"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.GetLines(tt.args.startRowIndex, tt.args.endRowIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.GetLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInput_GetStringLinesFrom(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		fromIndex int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantNewLines [][]string
		wantErr      bool
	}{
		{
			name: "GetLines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			},
			args: args{
				fromIndex: 0,
			},
			wantNewLines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			wantErr:      false,
		},
		{
			name: "GetLines",
			fields: fields{
				lines: [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}},
			},
			args: args{
				fromIndex: 1,
			},
			wantNewLines: [][]string{{"4", "5", "6"}, {"7", "8", "9"}},
			wantErr:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			gotNewLines, err := i.GetStringLinesFrom(tt.args.fromIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.GetStringLinesFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLines, tt.wantNewLines) {
				t.Errorf("Input.GetStringLinesFrom() = %v, want %v", gotNewLines, tt.wantNewLines)
			}
		})
	}
}

func TestInput_ReadAsStringGridFrom(t *testing.T) {
	type fields struct {
		lines [][]string
	}
	type args struct {
		fromIndex int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "ReadAsStringGridFrom",
			fields: fields{
				lines: [][]string{{"123"}, {"456"}},
			},
			args: args{
				fromIndex: 0,
			},
			want:    [][]string{{"1", "2", "3"}, {"4", "5", "6"}},
			wantErr: false,
		},
		{
			name: "ReadAsStringGridFrom",
			fields: fields{
				lines: [][]string{{"123"}, {"456"}},
			},
			args: args{
				fromIndex: 1,
			},
			want:    [][]string{{"4", "5", "6"}},
			wantErr: false,
		},
		{
			name: "ReadAsStringGridFrom",
			fields: fields{
				lines: [][]string{{"12", "3"}, {"456"}},
			},
			args: args{
				fromIndex: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "ReadAsStringGridFrom",
			fields: fields{
				lines: [][]string{{"123"}, {"45", "6"}},
			},
			args: args{
				fromIndex: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Input{
				lines: tt.fields.lines,
			}
			got, err := i.ReadAsStringGridFrom(tt.args.fromIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Input.ReadAsStringGridFrom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Input.ReadAsStringGridFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
