package main

import (
	"fmt"
	"math"
)

// floydWarshall вычисляет кратчайшие расстояния и восстанавливает пути для всех пар вершин.
func floydWarshall(graph map[int]map[int]int) {
	// Шаг 0: Добавляем "невидимые" вершины
	for _, neighbours := range graph {
		for v := range neighbours {
			if _, ok := graph[v]; !ok {
				graph[v] = map[int]int{}
			}
		}
	}

	// Шаг 1: Теперь, когда все вершины гарантированно есть в graph — собираем их в nodes
	nodes := make([]int, 0, len(graph))
	for u := range graph {
		nodes = append(nodes, u)
	}

	// Шаг 2: Инициализируем матрицу расстояний и массив для восстановления путей
	dist := make(map[int]map[int]int)
	next := make(map[int]map[int]int)

	for _, i := range nodes {
		dist[i] = make(map[int]int)
		next[i] = make(map[int]int)
		for _, j := range nodes {
			if i == j {
				dist[i][j] = 0
				next[i][j] = -1 // -1 значит, что это одна и та же вершина
			} else if w, ok := graph[i][j]; ok {
				dist[i][j] = w
				next[i][j] = j
			} else {
				dist[i][j] = math.MaxInt / 2 // используем MaxInt/2, чтобы избежать переполнения
				next[i][j] = -1
			}
		}
	}

	// Шаг 3: Алгоритм Флойда–Уоршелла (релаксация через промежуточные вершины)
	for _, k := range nodes {
		for _, i := range nodes {
			for _, j := range nodes {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
					next[i][j] = next[i][k]
				}
			}
		}
	}

	// Шаг 4: Проверка на отрицательные циклы
	for _, node := range nodes {
		if dist[node][node] < 0 {
			fmt.Println("В графе обнаружен отрицательный цикл.")
			return
		}
	}

	// Шаг 5: Вывод кратчайших расстояний и путей между всеми парами
	for _, i := range nodes {
		for _, j := range nodes {
			if dist[i][j] == math.MaxInt/2 {
				fmt.Printf("Нет пути из %d в %d\n", i, j)
			} else {
				// Восстановление пути
				path := []int{}
				for u := i; u != j; u = next[u][j] {
					path = append(path, u)
				}
				path = append(path, j)
				fmt.Printf("Путь из %d в %d: %v, расстояние: %d\n", i, j, path, dist[i][j])
			}
		}
	}
}

func main() {
	// Пример графа с положительными и отрицательными ребрами
	graph := map[int]map[int]int{
		0: {1: 3, 2: 8, 4: -4},
		1: {3: 1, 4: 7},
		2: {1: 4},
		3: {0: 2, 2: -5},
		4: {3: 6},
	}

	floydWarshall(graph)
}
