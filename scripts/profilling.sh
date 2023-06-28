#!/bin/bash
go test -benchmem -run=^$ ./app -bench ^BenchmarkSimulation$ -cpuprofile cpu.out
go tool pprof -http localhost:8080 cpu.out