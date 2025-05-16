package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"math/rand"
	"time"
)

const (
	screenWidth = 800
	screenHeight = 600
	delay = 5 * time.Millisecond
)

var (
	array []int
	running bool
	mode = "menu"	
	comparisons int
	inputText = "100"          
	isTyping = false          
	N = 100
	mousePressedLastFrame bool
)

type Button struct {
	x, y, w, h int
	label      string
	action     func()
}

var buttons []Button
var returnButton Button


func (b Button) isHovered(mx, my int) bool {
	return mx >= b.x && mx <= b.x+b.w && my >= b.y && my <= b.y+b.h
}

func (b Button) draw(screen *ebiten.Image) {
	col := color.RGBA{100, 100, 100, 255}
	mx, my := ebiten.CursorPosition()

	if b.isHovered(mx, my) {
		col = color.RGBA{150, 150, 150, 255}
	}

	ebitenutil.DrawRect(screen, float64(b.x), float64(b.y), float64(b.w), float64(b.h), col)
	textX := b.x + b.w/2 - len(b.label)*3
	textY := b.y + b.h/2 - 5
	ebitenutil.DebugPrintAt(screen, b.label, textX, textY)
}
type Game struct{}

func (g *Game) Update() error {
	mx, my := ebiten.CursorPosition()
	mousePressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	if mousePressed && !mousePressedLastFrame {
		if mx >= 300 && mx <= 500 && my >= 110 && my <= 140 {
			isTyping = true
		} else {
			isTyping = false
		}
	}

	if isTyping {
		for _, r := range ebiten.AppendInputChars(nil) {
			if r >= '0' && r <= '9' && len(inputText) < 5 {
				inputText += string(r)
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(inputText) > 0 {
			inputText = inputText[:len(inputText)-1]
		}
	}

	if mousePressed && !mousePressedLastFrame {
		if mode == "menu" {
			for _, b := range buttons {
				if b.isHovered(mx, my) {
					b.action()
				}
			}
		} else if !running && returnButton.isHovered(mx, my) {
			mode = "menu"
		}
	}

	mousePressedLastFrame = mousePressed

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{32, 32, 32, 255})

	if mode == "menu" {
		for _, b := range buttons {
			b.draw(screen)
		}
		ebitenutil.DebugPrintAt(screen, "Tamanho do vetor:", 300, 80)
		border := color.RGBA{100, 100, 100, 255}
		if isTyping {
			border = color.RGBA{150, 150, 150, 255}
		}
		ebitenutil.DrawRect(screen, 300, 95, 200, 30, border)
		ebitenutil.DebugPrintAt(screen, inputText, 310, 100)
		ebitenutil.DebugPrintAt(screen, "Clique em um algoritmo", 300, 130)
		return
	}

	barWidth := float64(screenWidth) / float64(len(array))

	for idx, val := range array {
		x := float64(idx) * barWidth
		h := float64(val) / float64(len(array)) * screenHeight

		col := color.RGBA{0, 255, 0, 255}
		if running {
			col = color.RGBA{0, 150, 255, 255}
		}

		ebitenutil.DrawRect(screen, x, screenHeight-h, barWidth-1, h, col)
	}
	
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Algoritmo: %s", mode), 20, 20)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Comparacoes: %d", comparisons), 20, 40)
	if !running {
		returnButton.draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func startSort(selected string) {
	mode = selected
	fmt.Sscanf(inputText, "%d", &N)
	N = min(N, 500)
	N = max(N, 1)
	array = rand.Perm(N)
	running = true
	comparisons = 0;
	go func() {
		switch mode {
		case "bubble":
			bubbleSort(array)
		case "selection":
			selectionSort(array)
		case "insertion":
			insertionSort(array)
		case "merge":
			mergeSort(array,0,N-1)
		}
		running = false
	}()
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			comparisons += 1
			time.Sleep(delay)
		}
	}
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
			comparisons += 1
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		time.Sleep(delay * 10)
	}
}

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			comparisons += 1
			arr[j+1] = arr[j]
			j--
			time.Sleep(delay)
		}
		arr[j+1] = key
		time.Sleep(delay * 10)
	}
}

func mergeSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	m := (l + r) / 2
	mergeSort(arr, l, m)
	mergeSort(arr, m+1, r)
	merge(arr, l, m, r)
}

func merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m

	left := make([]int, n1)
	right := make([]int, n2)
	copy(left, arr[l:m+1])
	copy(right, arr[m+1:r+1])

	i, j, k := 0, 0, l
	for i < n1 || j < n2 {
		comparisons += 1
		if j == n2 || (i < n1 && left[i] <= right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
		time.Sleep(delay)
	}
}


func main() {
	rand.Seed(time.Now().UnixNano())

	buttons = []Button{
		{300, 150, 200, 40, "Bubble Sort", func() { startSort("bubble") }},
		{300, 200, 200, 40, "Selection Sort", func() { startSort("selection") }},
		{300, 250, 200, 40, "Insertion Sort", func() { startSort("insertion") }},
		{300, 300, 200, 40, "Merge Sort", func() { startSort("merge") }},
	}

	returnButton = Button{
		x: 300, y: 500, w: 200, h: 40,
		label: "Voltar ao menu",
		action: func() { mode = "menu" },
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sort Animation")

	ebiten.RunGame(&Game{})
}
