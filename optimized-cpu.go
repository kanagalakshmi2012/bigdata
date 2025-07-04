package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Solution struct {
	CPU      float64
	IOPS     float64
	Network  float64
	Fitness  float64
}

func randomSolution() Solution {
	return Solution{
		CPU:     rand.Float64() * 100,
		IOPS:    rand.Float64() * 100,
		Network: rand.Float64() * 100,
	}
}

func fitness(s Solution, wCPU, wIOPS, wNet float64) float64 {
	cpuScore := 100 - math.Abs(70-s.CPU)
	iopsScore := 100 - math.Abs(80-s.IOPS)
	netScore := 100 - math.Abs(75-s.Network)
	return wCPU*cpuScore + wIOPS*iopsScore + wNet*netScore
}

func crossover(p1, p2 Solution) Solution {
	return Solution{
		CPU:     (p1.CPU + p2.CPU) / 2,
		IOPS:    (p1.IOPS + p2.IOPS) / 2,
		Network: (p1.Network + p2.Network) / 2,
	}
}

func mutate(s Solution) Solution {
	if rand.Float64() < 0.1 {
		s.CPU += (rand.Float64()*20 - 10)
		if s.CPU < 0 {
			s.CPU = 0
		}
		if s.CPU > 100 {
			s.CPU = 100
		}
	}
	if rand.Float64() < 0.1 {
		s.IOPS += (rand.Float64()*20 - 10)
		if s.IOPS < 0 {
			s.IOPS = 0
		}
		if s.IOPS > 100 {
			s.IOPS = 100
		}
	}
	if rand.Float64() < 0.1 {
		s.Network += (rand.Float64()*20 - 10)
		if s.Network < 0 {
			s.Network = 0
		}
		if s.Network > 100 {
			s.Network = 100
		}
	}
	return s
}

func selectParent(pop []Solution) Solution {
	total := 0.0
	for _, s := range pop {
		total += s.Fitness
	}
	r := rand.Float64() * total
	sum := 0.0
	for _, s := range pop {
		sum += s.Fitness
		if sum >= r {
			return s
		}
	}
	return pop[len(pop)-1]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	popSize := 50
	generations := 100
	wCPU, wIOPS, wNet := 0.4, 0.3, 0.3
	population := make([]Solution, popSize)
	for i := 0; i < popSize; i++ {
		population[i] = randomSolution()
		population[i].Fitness = fitness(population[i], wCPU, wIOPS, wNet)
	}
	for gen := 0; gen < generations; gen++ {
		newPop := make([]Solution, 0, popSize)
		for len(newPop) < popSize {
			p1 := selectParent(population)
			p2 := selectParent(population)
			child := crossover(p1, p2)
			child = mutate(child)
			child.Fitness = fitness(child, wCPU, wIOPS, wNet)
			newPop = append(newPop, child)
		}
		population = newPop
		best := population[0]
		for _, s := range population {
			if s.Fitness > best.Fitness {
				best = s
			}
		}
		if gen%10 == 0 {
			fmt.Printf("Gen %d - Best Fitness: %.2f CPU: %.2f IOPS: %.2f Net: %.2f\n", gen, best.Fitness, best.CPU, best.IOPS, best.Network)
		}
	}
	best := population[0]
	for _, s := range population {
		if s.Fitness > best.Fitness {
			best = s
		}
	}
	fmt.Printf("Optimized solution - Fitness: %.2f CPU: %.2f IOPS: %.2f Net: %.2f\n", best.Fitness, best.CPU, best.IOPS, best.Network)
}
