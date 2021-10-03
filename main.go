package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const symbols = "23456789TJQKA"

var symbolRanking = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Result int

const (
	Tie Result = iota
	Hand1
	Hand2
)

var results = map[Result]string{
	Tie:   "Tie",
	Hand1: "Hand 1",
	Hand2: "Hand 2",
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("wrong number of args, expect 2, got %d", len(os.Args)-1)
		return
	}

	result, err := compareHand(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	println(result)
}

func compareHand(firstHand, secondHand string) (string, error) {
	firstHandRank, err := calculateHandRank(firstHand)
	if err != nil {
		return "", fmt.Errorf("cannot parse the first hand, err: %+v", err)
	}

	secondHandRank, err := calculateHandRank(secondHand)
	if err != nil {
		return "", fmt.Errorf("cannot parse the second hand, err: %+v", err)
	}

	return compareHandRank(firstHandRank, secondHandRank), nil
}

type handRank []struct {
	value string
	count int
}

func calculateHandRank(value string) (rank handRank, err error) {
	if len(value) != 5 {
		err = fmt.Errorf("wrong length of hand, expect 5, got %d", len(value))
		return
	}

	r := map[string]int{}

	for _, v := range strings.Split(value, "") {
		if _, exist := symbolRanking[v]; !exist {
			err = fmt.Errorf("wrong symbol, expect %v, got %v", symbols, v)
			return
		}

		if _, ok := r[v]; ok {
			r[v] += 1
		} else {
			r[v] = 1
		}
	}

	for k, v := range r {
		rank = append(rank, struct {
			value string
			count int
		}{k, v})
	}

	sort.Slice(rank, func(i, j int) bool {
		if rank[i].count == rank[j].count {
			return symbolRanking[rank[i].value] > symbolRanking[rank[j].value]
		}

		return rank[i].count > rank[j].count
	})

	return
}

func compareHandRank(firstRank, secondRank handRank) string {
	for i, f := range firstRank {
		s := secondRank[i]

		if f.count > s.count {
			return results[Hand1]
		}
		if f.count < s.count {
			return results[Hand2]
		}
	}

	for i, f := range firstRank {
		s := secondRank[i]

		if symbolRanking[f.value] > symbolRanking[s.value] {
			return results[Hand1]
		}
		if symbolRanking[f.value] < symbolRanking[s.value] {
			return results[Hand2]
		}
	}

	return results[Tie]
}
