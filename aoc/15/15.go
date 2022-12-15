package _15

import (
	"AoC2022/helper"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"image"
	"regexp"
)

func distance(x image.Point, y image.Point) int {
	return helper.Abs(x.X - y.X) + helper.Abs(x.Y - y.Y)
}

func t01() {
	lines := helper.Split(helper.ReadLines("aoc/15/15.inp"), "\n")

	pattern,_ := regexp.Compile(`.*?x=(-?\d*), y=(-?\d*).*?x=(-?\d*), y=(-?\d*)$`)
	sensors := make([]image.Point, 0)
	beacons := make([]image.Point, 0)
	helper.Apply(lines, func(line string) {
		match := pattern.FindAllStringSubmatch(line, -1)
		sensors = append(sensors, image.Point{X: helper.ConvInt(match[0][1]), Y: helper.ConvInt(match[0][2])})
		beacons = append(beacons, image.Point{X: helper.ConvInt(match[0][3]), Y: helper.ConvInt(match[0][4])})
	})
	sensor_distance := helper.Collect(sensors, func(sensor image.Point) int {
		return helper.MinArray(helper.Collect(beacons, func(beacon image.Point) int {
			return distance(sensor, beacon)
		}))
	})

	yrow := 2000000
	ypoints := mapset.NewSet[image.Point]()
	for i := range sensors {
		sensor := sensors[i]
		dist := sensor_distance[i]
		fmt.Println(dist)
		for k := 0; k <= dist; k++ {
			j := helper.Abs(yrow- sensor.Y)
			if distance(sensor, image.Point{X: sensor.X + k , Y: sensor.Y + j}) <= dist {
				if sensor.Y + j == yrow {
					ypoints.Add(image.Point{X: sensor.X + k , Y: sensor.Y + j})
					ypoints.Add(image.Point{X: sensor.X - k , Y: sensor.Y + j})
				}
				if sensor.Y - j == yrow {
					ypoints.Add(image.Point{X: sensor.X + k , Y: sensor.Y - j})
					ypoints.Add(image.Point{X: sensor.X - k , Y: sensor.Y - j})
				}
			}
		}
	}
	helper.Apply(beacons, func(beacon image.Point) { ypoints.Remove(beacon) })
	fmt.Println(ypoints.Cardinality())
}

func t02() {
	lines := helper.Split(helper.ReadLines("aoc/15/15.inp"), "\n")

	pattern,_ := regexp.Compile(`.*?x=(-?\d*), y=(-?\d*).*?x=(-?\d*), y=(-?\d*)$`)
	sensors := make([]image.Point, 0)
	beacons := make([]image.Point, 0)
	helper.Apply(lines, func(line string) {
		match := pattern.FindAllStringSubmatch(line, -1)
		sensors = append(sensors, image.Point{X: helper.ConvInt(match[0][1]), Y: helper.ConvInt(match[0][2])})
		beacons = append(beacons, image.Point{X: helper.ConvInt(match[0][3]), Y: helper.ConvInt(match[0][4])})
	})
	sensor_distance := helper.Collect(sensors, func(sensor image.Point) int {
		return helper.MinArray(helper.Collect(beacons, func(beacon image.Point) int {
			return distance(sensor, beacon)
		}))
	})

	max := 4000000
	pos := image.Point{0,0}
	for {
		if pos.X >= max{
			pos.Y += 1
			pos.X = 0
		}
		jumped := false
		for i,sensor := range sensors {
			if distance(pos, sensor) <= sensor_distance[i] {
				y_delta := helper.Abs(sensor.Y - pos.Y)
				pos.X = sensor.X + (sensor_distance[i]-y_delta+1)
				jumped = true
				break
			}
		}
		if !jumped {
			fmt.Printf("found point: %v, %d", pos, pos.X*max + pos.Y)
			return
		}

	}


}

func Run() {
	//t01()
	t02()
}
