package main

import (
	"container/heap"
	"fmt"
	"log"
	"strconv"
)

type SupplyAndDemand struct {
	Shares int
	Value  int
}

type Supply struct{ SupplyAndDemand }
type Demand struct{ SupplyAndDemand }

func (s Supply) Less(other SupplyAndDemand) bool {
	return s.Value <= other.Value
}

func (d Demand) Less(other SupplyAndDemand) bool {
	return d.Value >= other.Value
}

type SupplyHeap []Supply
type DemandHeap []Demand

func (h SupplyHeap) Len() int           { return len(h) }
func (h SupplyHeap) Less(i, j int) bool { return h[i].Less(h[j].SupplyAndDemand) }
func (h SupplyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *SupplyHeap) Push(x interface{}) {
	*h = append(*h, x.(Supply))
}
func (h *SupplyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h DemandHeap) Len() int           { return len(h) }
func (h DemandHeap) Less(i, j int) bool { return h[i].Less(h[j].SupplyAndDemand) }
func (h DemandHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *DemandHeap) Push(x interface{}) {
	*h = append(*h, x.(Demand))
}
func (h *DemandHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Stock struct {
	SupplyHeap
	DemandHeap
	LastTrade int
}

func NewStock() *Stock {
	return &Stock{
		LastTrade: -1,
	}
}

func (s *Stock) Buy(shares, value int) {
	heap.Push(&s.DemandHeap, Demand{SupplyAndDemand{Shares: shares, Value: value}})
	s.simulateTrades()
}

func (s *Stock) Sell(shares, value int) {
	heap.Push(&s.SupplyHeap, Supply{SupplyAndDemand{Shares: shares, Value: value}})
	s.simulateTrades()
}

func (s *Stock) simulateTrades() {
	for len(s.SupplyHeap) > 0 && len(s.DemandHeap) > 0 && s.SupplyHeap[0].Less(s.DemandHeap[0].SupplyAndDemand) {
		s.LastTrade = s.SupplyHeap[0].Value
		// fmt.Println("print", s.SupplyHeap[0].Shares, s.DemandHeap[0].Shares)
		if s.SupplyHeap[0].Shares == s.DemandHeap[0].Shares {
			heap.Pop(&s.SupplyHeap)
			heap.Pop(&s.DemandHeap)
		} else if s.SupplyHeap[0].Shares < s.DemandHeap[0].Shares {
			supply := heap.Pop(&s.SupplyHeap).(Supply)
			shares := supply.Shares
			// values := supply.Value
			s.DemandHeap[0].Shares -= shares
			// fmt.Println("Supply ", shares, values)
		} else {
			demand := heap.Pop(&s.DemandHeap).(Demand)
			shares := demand.Shares
			// values := demand.Value

			s.SupplyHeap[0].Shares -= shares
			// fmt.Println("Demand ", shares, values)
		}
		// fmt.Println("Print => ", s.SupplyHeap[0].Shares)
	}
}

func (s *Stock) Print() string {
	// fmt.Println("DATA => ", s.SupplyHeap, s.DemandHeap, s.LastTrade)
	ask := "-"
	if len(s.SupplyHeap) > 0 {
		ask = strconv.Itoa(s.SupplyHeap[0].Value)
	}

	bid := "-"
	if len(s.DemandHeap) > 0 {
		bid = strconv.Itoa(s.DemandHeap[0].Value)
	}

	price := "-"
	if s.LastTrade > 0 {
		price = strconv.Itoa(s.LastTrade)
	}

	return fmt.Sprintf("%s %s %s", ask, bid, price)
}

func main() {
	var numCases int
	fmt.Scan(&numCases)

	for c := 0; c < numCases; c++ {
		var stock Stock
		var numCommands int
		fmt.Scan(&numCommands)
		for i := 0; i < numCommands; i++ {
			var cmd, sharesStr, garbage, garbage2, valueStr string
			fmt.Scanf("%s %s %s %s %s", &cmd, &sharesStr, &garbage, &garbage2, &valueStr)

			shares, err := strconv.Atoi(sharesStr)
			if err != nil {
				log.Fatal(err)
			}

			value, err := strconv.Atoi(valueStr)
			if err != nil {
				log.Fatal(err)
			}

			if cmd[0] == 'b' {
				stock.Buy(shares, value)
			} else {
				stock.Sell(shares, value)
			}

			fmt.Println(stock.Print())
		}
	}
}

// func main() {
// 	var numCases int
// 	fmt.Scan(&numCases)

// 	for c := 0; c < numCases; c++ {
// 		stock := NewStock()

// 		var numCommands int
// 		fmt.Scan(&numCommands)

// 		for i := 0; i < numCommands; i++ {
// 			var cmd, shares, _, _, valueStr string
// 			fmt.Scan(&cmd, &shares, _, _, &valueStr)
// 			value, _ := strconv.Atoi(valueStr)

// 			if cmd[0] == 'b' {
// 				stock.Buy(shares, value)
// 			} else {
// 				stock.Sell(shares, value)
// 			}
// 			fmt.Println(stock)
// 		}
// 	}
// }

// type Stock struct {
// 	price  int
// 	volume int
// }

// type AskPriorityQueue []*Stock

// func (pq AskPriorityQueue) Len() int { return len(pq) }
// func (pq AskPriorityQueue) Less(i, j int) bool {
// 	return pq[i].price < pq[j].price
// }
// func (pq AskPriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }
// func (pq *AskPriorityQueue) Push(x interface{}) {
// 	stock := x.(*Stock)
// 	*pq = append(*pq, stock)
// }
// func (pq *AskPriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(*pq)
// 	item := old[n-1]
// 	*pq = old[:n-1]
// 	return item
// }

// type BidPriorityQueue []*Stock

// func (pq BidPriorityQueue) Len() int { return len(pq) }
// func (pq BidPriorityQueue) Less(i, j int) bool {
// 	return pq[i].price > pq[j].price
// }
// func (pq BidPriorityQueue) Swap(i, j int) {
// 	pq[i], pq[j] = pq[j], pq[i]
// }
// func (pq *BidPriorityQueue) Push(x interface{}) {
// 	stock := x.(*Stock)
// 	*pq = append(*pq, stock)
// }
// func (pq *BidPriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(*pq)
// 	item := old[n-1]
// 	*pq = old[:n-1]
// 	return item
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)

// 	// read total cases
// 	var cases int
// 	if scanner.Scan() {
// 		fmt.Sscan(scanner.Text(), &cases)
// 	}

// 	for i := 0; i < cases; i++ {
// 		var n, volume, price int
// 		var command, garbage string
// 		askQueue := make(AskPriorityQueue, 0)
// 		bidQueue := make(BidPriorityQueue, 0)
// 		mapAsk := make(map[int]int)
// 		mapBid := make(map[int]int)

// 		if scanner.Scan() {
// 			fmt.Sscan(scanner.Text(), &n)
// 		}
// 		for j := 0; j < n; j++ {
// 			if scanner.Scan() {
// 				//fields := strings.Fields(scanner.Text())
// 				fmt.Sscanf(scanner.Text(), "%s %d %s %s %d", &command, &volume, &garbage, &garbage, &price)
// 			}

// 			if command == "buy" {
// 				mapBid[price] += volume
// 				heap.Push(&bidQueue, &Stock{volume: mapBid[price], price: price})
// 			}

// 			if command == "sell" {
// 				mapAsk[price] += volume
// 				heap.Push(&askQueue, &Stock{volume: mapAsk[price], price: price})
// 			}
// 		}

// 		for askQueue.Len() > 0 {
// 			item := heap.Pop(&askQueue).(*Stock)
// 			fmt.Printf("Ask Volume: %d, Priority: %d\n", item.volume, item.price)
// 		}
// 		for bidQueue.Len() > 0 {
// 			item := heap.Pop(&bidQueue).(*Stock)
// 			fmt.Printf("Bid Volume: %d, Priority: %d\n", item.volume, item.price)
// 		}
// 	}
// }
