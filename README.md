# Sort Animation (Go + Ebiten)

Este projeto é uma visualização animada de algoritmos de ordenação implementados em Go utilizando a biblioteca [Ebiten](https://ebiten.org/).

## Algoritmos disponíveis

- Bubble Sort
- Selection Sort
- Insertion Sort
- Merge Sort

## Como rodar

Talvez seja necessario instalar algumas dependencias:
```bash
sudo apt install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config
```
Logo após seguir o processo abaixo:
```bash
git clone https://github.com/gustav0leal/SortGolang
cd SortGolang
go get github.com/hajimehoshi/ebiten/v2
go run sort_animation2.go
```
