package index

import (
	"fmt"
	"indexer/preprocessor"
	"indexer/scrapper"
	"os"
	"strings"
	"sync"
	// "time"
)

var (
	mu            sync.Mutex // can be considered to use RWLock if index will be dynamically scaled(profit in fast read)
	maxGoroutines = 40
	semaphore     = make(chan struct{}, maxGoroutines)
)

func hash(key string) int {
	hashVal := 0
	for _, char := range key {
		hashVal += int(char)
	}

	return hashVal
}

type IndexItem struct {
	Key, Value string
}

type Index struct {
	Buckets         []IndexItem
	OverflowBuckets []LinkedList
	Length          uint32
}

func (idx *Index) AddItem(item IndexItem, routineName string) {
	hash := hash(item.Key)
	bucketIdx := hash % int(idx.Length)
	mu.Lock()
	fmt.Printf("Routine %s acquiring the Lock \n", routineName)
	// time.Sleep(5 * time.Second)
	defer mu.Unlock()
	bucketValue := idx.Buckets[bucketIdx]

	if bucketValue == (IndexItem{}) { // then it means that the value in the bucket is empty
		idx.Buckets[bucketIdx] = item
	} else {
		if item.Key == bucketValue.Key {
			idx.Buckets[bucketIdx].Value = idx.Buckets[bucketIdx].Value + "; " + item.Value
		} else {
			idx.OverflowBuckets[bucketIdx].Insert(item)
		}
	}
	fmt.Printf("Routine %s releasing the Lock \n", routineName)
}

func (idx *Index) GetItem(key string) string {
	hash := hash(key)
	bucketIdx := hash % int(idx.Length)
	mu.Lock()
	defer mu.Unlock()
	bucketValue := idx.Buckets[bucketIdx]

	if bucketValue == (IndexItem{}) { // then it means that the value in the bucket is empty
		return ""
	} else {
		if key == bucketValue.Key {
			return bucketValue.Value
		} else {
			return idx.OverflowBuckets[bucketIdx].Search(key)
		}
	}
}

func (idx *Index) Display() {
	// fmt.Println(idx.Buckets)
	count := 0
	for _, item := range idx.Buckets {
		if item != (IndexItem{}) {
			fmt.Println(item)
			count += 1
		}
	}
	fmt.Printf("Total amount non empty buckets = %d \n", count)
	fmt.Println(strings.Repeat("*", 50))
	for _, list := range idx.OverflowBuckets {
		if list != (LinkedList{}) {
			fmt.Println(list.Display())
		}
	}
}

func BuildIndex(folders []string, size uint32) Index {
	var wg sync.WaitGroup
	temp := make([]IndexItem, size)
	temp_over := make([]LinkedList, size)
	idx := Index{Buckets: temp, Length: size, OverflowBuckets: temp_over}
	for _, folder := range folders {
		files := scrapper.GetAllFilesToRead(folder)
		for i, file := range files {
			if i >= 2000 {
				break
			}
			wg.Add(1)
			semaphore <- struct{}{}
			go processFile(file, &wg, &idx, fmt.Sprintf("goroutine-%d", i))
		}
	}
	wg.Wait()

	return idx
}

func processFile(filePath string, wg *sync.WaitGroup, idx *Index, routineName string) {
	defer wg.Done()
	data, _ := os.ReadFile(filePath)
	text := strings.Fields(string(data))
	fmt.Printf("Routine %s started execution \n", routineName)
	for _, term := range text {
		correctTerm, error := preprocessor.CleanTerm(strings.ToLower(term))
		if error != nil {
			panic(error)
		}
		if !preprocessor.CheckForStopWord(correctTerm) {
			if correctTerm != "" {
				idx.AddItem(IndexItem{correctTerm, filePath}, routineName)
			}
		}
	}
	<-semaphore
}
