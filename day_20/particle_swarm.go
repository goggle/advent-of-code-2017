package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
	z int
}

type Particle struct {
	index int
	p     Point
	v     Point
	a     Point
}

func (p *Point) Add(q Point) {
	p.x += q.x
	p.y += q.y
	p.z += q.z
}

func (p *Point) Multiply(alpha int) {
	p.x *= alpha
	p.y *= alpha
	p.z *= alpha
}

func (p *Point) ManhattenDistance() int {
	x := p.x
	y := p.y
	z := p.z
	if x < 0 {
		x *= -1
	}
	if y < 0 {
		y *= -1
	}
	if z < 0 {
		z *= -1
	}
	return x + y + z
}

func (p *Point) EuclideanDistance2() int {
	x := p.x
	y := p.y
	z := p.z
	return x*x + y*y + z*z
}

func (p *Point) Difference(q *Point) Point {
	x := p.x - q.x
	y := p.y - q.y
	z := p.z - q.z
	return Point{x, y, z}
}

func (p *Point) Equal(q *Point) bool {
	if p.x == q.x && p.y == q.y && p.z == q.z {
		return true
	}
	return false
}

func (par *Particle) Update() {
	par.v.x += par.a.x
	par.v.y += par.a.y
	par.v.z += par.a.z
	par.p.x += par.v.x
	par.p.y += par.v.y
	par.p.z += par.v.z
}

type ParticleSorter []Particle

func (ps ParticleSorter) Len() int {
	return len(ps)
}

func (ps ParticleSorter) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (ps ParticleSorter) Less(i, j int) bool {
	ai, aj := ps[i].a.ManhattenDistance(), ps[j].a.ManhattenDistance()
	if ai < aj {
		return true
	} else if ai > aj {
		return false
	}
	vi, vj := ps[i].v.ManhattenDistance(), ps[j].v.ManhattenDistance()
	if vi < vj {
		return true
	} else if vi > vj {
		return false
	}
	pi, pj := ps[i].p.ManhattenDistance(), ps[j].p.ManhattenDistance()
	if pi < pj {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	particles := []Particle{}
	index := 0
	for scanner.Scan() {
		// This is very fragile...
		rowElements := strings.Split(scanner.Text(), " ")
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input: ", err)
		}
		re := regexp.MustCompile(`[[:alpha:]]=<(-?\d+),(-?\d+),(-?\d+)>`)
		particle := Particle{}
		particle.index = index
		for i := 0; i < 3; i++ {
			match := re.FindStringSubmatch(rowElements[i])
			var x, y, z int
			var err error
			x, err = strconv.Atoi(match[1])
			if err != nil {
				fmt.Println(err)
			}
			y, err = strconv.Atoi(match[2])
			if err != nil {
				fmt.Println(err)
			}
			z, err = strconv.Atoi(match[3])
			if err != nil {
				fmt.Println(err)
			}
			switch i {
			case 0:
				particle.p.x = x
				particle.p.y = y
				particle.p.z = z
			case 1:
				particle.v.x = x
				particle.v.y = y
				particle.v.z = z
			case 2:
				particle.a.x = x
				particle.a.y = y
				particle.a.z = z
			}
		}
		index++
		particles = append(particles, particle)
	}

	nearestIndex := getNearestIndexLongRun(particles)
	fmt.Println("Solution to part 1:", nearestIndex)

	timeLimit := calculateTimeLimit(particles)
	fmt.Println("Solution to part 2:", simulate(particles, timeLimit))

}

func getNearestIndexLongRun(particles []Particle) int {
	sort.Sort(ParticleSorter(particles))
	return particles[0].index
}

func calculatePosition(particle Particle, t int) Point {
	position := Point(particle.p)
	v0 := Point(particle.v)
	a0 := Point(particle.a)
	v0.Multiply(t)
	a0.Multiply(t * (t + 1) / 2)
	position.Add(v0)
	position.Add(a0)

	return position
}

func calculateTimeLimit(particles []Particle) int {
	timeLimit := 0
	for i, par1 := range particles {
		for _, par2 := range particles[i+1:] {
			diffP := (par1.p).Difference(&par2.p)
			distPrev := diffP.EuclideanDistance2()
			decreasing := true
			for n := 1; decreasing; n++ {
				pos1 := calculatePosition(par1, n)
				pos2 := calculatePosition(par2, n)
				diff := pos1.Difference(&pos2)
				dist := diff.EuclideanDistance2()
				if dist > distPrev {
					decreasing = false
					if n > timeLimit {
						timeLimit = n
					}
				} else {
					distPrev = dist
				}
			}
		}
	}
	return timeLimit
}

func removeCollisions(particles []Particle) []Particle {
	newParticles := []Particle{}
	removeIndices := []int{}
	for i, par1 := range particles {
		iAdded := false
		for j, par2 := range particles {
			if i == j {
				continue
			}
			if par1.p.Equal(&par2.p) {
				if !iAdded && !contains(removeIndices, i) {
					removeIndices = append(removeIndices, i)
					iAdded = true
				}
				if !contains(removeIndices, j) {
					removeIndices = append(removeIndices, j)
				}
			}
		}
	}
	for i, par := range particles {
		if !contains(removeIndices, i) {
			newParticles = append(newParticles, par)
		}
	}
	return newParticles
}

func contains(indices []int, val int) bool {
	for _, ind := range indices {
		if ind == val {
			return true
		}
	}
	return false
}

func update(particles []Particle) {
	for i := range particles {
		particles[i].Update()
	}
}

func simulate(particles []Particle, timeLimit int) int {
	newParticles := particles
	for i := 0; i < timeLimit; i++ {
		newParticles = removeCollisions(newParticles)
		update(newParticles)
	}
	return len(newParticles)
}
