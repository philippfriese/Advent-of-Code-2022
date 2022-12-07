package _5

import (
	"AoC2022/helper"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type File struct {
	name   string
	ending string
	size   int
}

type Directory struct {
	name   string
	parent int
	files  map[string]int
	dirs   map[string]int
	size   int
}

type Content interface {
	File | Directory
}

func assign_size(dir int, dirs *[]Directory, files *[]File) int {
	dir_size := 0
	for _, f := range (*dirs)[dir].files {
		dir_size += (*files)[f].size
	}
	for _, d := range (*dirs)[dir].dirs {
		dir_size += assign_size(d, dirs, files)
	}
	(*dirs)[dir].size = dir_size
	return dir_size
}

func print_dir(dir int, dirs *[]Directory, files *[]File) {
	fmt.Printf("%s: %d\n", (*dirs)[dir].name, (*dirs)[dir].size)
	for _, d := range (*dirs)[dir].dirs {
		print_dir(d, dirs, files)
	}
}

func sum_dirs(dir int, dirs *[]Directory, files *[]File, acc *int, threshold int) {
	for _, d := range (*dirs)[dir].dirs {
		if (*dirs)[d].size < threshold {
			*acc += (*dirs)[d].size
		}
		sum_dirs(d, dirs, files, acc, threshold)
	}
}

func parse(fn string) (Directory, []Directory, []File) {
	lines := helper.Split(helper.ReadLines(fn), "\n$")

	dirs := make([]Directory, 0)
	files := make([]File, 0)
	root := Directory{name: "/", parent: -1, files: map[string]int{}, dirs: map[string]int{}, size: 0}
	dirs = append(dirs, root)
	current_dir := 0
	for _, s := range lines[1:] {
		ls := strings.Split(s, "\n")
		cmd := ls[0][1:]
		cmd_tokens := strings.Split(cmd, " ")
		switch cmd_tokens[0] {
		case "cd":
			{

				if cmd_tokens[1] == ".." {
					current_dir = dirs[current_dir].parent
				} else {
					current_dir = dirs[current_dir].dirs[cmd_tokens[1]]
				}
			}
		case "ls":
			{
				for _, token := range ls[1:] {
					subtokens := strings.Split(token, " ")
					if subtokens[0] == "dir" {
						d := Directory{name: subtokens[1], parent: current_dir,
							files: map[string]int{}, dirs: map[string]int{}, size: 0}
						dirs = append(dirs, d)
						dirs[current_dir].dirs[d.name] = len(dirs) - 1
					} else {
						size, _ := strconv.Atoi(subtokens[0])
						endings := strings.Split(subtokens[1], ".")
						ending := endings[len(endings)-1]
						f := File{name: subtokens[1], ending: ending, size: size}
						files = append(files, f)
						dirs[current_dir].files[f.name] = len(files) - 1
					}
				}
			}
		}
	}

	return root, dirs, files
}

func t01() {
	_, dirs, files := parse("aoc/07/07.inp")
	assign_size(0, &dirs, &files)
	acc := 0
	sum_dirs(0, &dirs, &files, &acc, 100000)
	fmt.Println(acc)
}

func t02() {
	_, dirs, files := parse("aoc/07/07.inp")
	assign_size(0, &dirs, &files)
	total_size := 70000000
	required_size := 30000000
	current_size := total_size - dirs[0].size
	missing_size := required_size - current_size

	sorted := dirs
	slices.SortFunc(sorted, func(l Directory, r Directory) bool { return l.size < r.size })
	sorted = helper.Filter(sorted, func(d Directory) bool { return d.size >= missing_size })
	fmt.Println(sorted[0].size)
}
func Run() {
	// funfunfunfunfunfunfunfunfunfunfun
	t01()
	t02()
}
