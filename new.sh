#!/bin/zsh

DAY=$(date +%e)
echo $DAY

if [[ ! -d "$DAY" ]]
then
  echo "not exists"
  mkdir "aoc/$DAY"
  touch "aoc/$DAY/$DAY.go"
  touch "aoc/$DAY/$DAY.inp"
  touch "aoc/$DAY/test.inp"
  cat << EOF > "aoc/$DAY/$DAY.go"
package _$DAY

func t01() {
}

func t02() {
}

func Run() {
	t01()
	t02()
}
EOF
  git add "aoc/$DAY/$DAY.go"
  git add "aoc/$DAY/$DAY.inp"
  git add "aoc/$DAY/test.inp"
fi