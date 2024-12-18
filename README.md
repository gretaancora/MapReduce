1) Launch first reducer -> $ go run main.go -port 50053
2) Launch second reducer -> $ go run main.go -port 50054
3) Launch first mapper -> $ go run main.go -port 50052
4) Launch second mapper -> $ go run main.go -port 50051
5) Launch master -> $ go run main.go
