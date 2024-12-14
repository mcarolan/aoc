package day9

import (
	"fmt"
	"sort"
)

type interval struct {
	start, end int
	value      int
}

const FREE_BLOCK_ID int = -1

func print(blocks []interval, freeBlocks []interval) {
	contains := func(intervals []interval, v int) (interval, bool) {
		for _, i := range intervals {
			if v >= i.start && v < i.end {
				return i, true
			}
		}
		var zero interval
		return zero, false
	}

	maxIndex := 0

	for _, i := range blocks {
		maxIndex = max(maxIndex, i.end)
	}
	for _, i := range freeBlocks {
		maxIndex = max(maxIndex, i.end)
	}

	for i := range maxIndex {
		b, hasBlock := contains(blocks, i)
		if hasBlock {
			fmt.Printf("%d", b.value)
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func insertSorted(intervals []interval, newInterval interval) []interval {
	index := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].start > newInterval.start
	})
	intervals = append(intervals[:index], append([]interval{newInterval}, intervals[index:]...)...)
	return intervals
}

func SolvePart1(input string) int {
	diskMap := make([]int, 0)

	for _, c := range input {
		diskMap = append(diskMap, int(c-'0'))
	}

	blocks := make([]interval, 0)
	freeBlocks := make([]interval, 0)

	nextBlockStart := 0
	for i, count := range diskMap {
		isFreeSpace := i%2 != 0

		if isFreeSpace {
			freeBlocks = append(freeBlocks, interval{
				start: nextBlockStart,
				end:   nextBlockStart + count,
				value: FREE_BLOCK_ID,
			})
		} else {
			blocks = append(blocks, interval{
				start: nextBlockStart,
				end:   nextBlockStart + count,
				value: i / 2,
			})
		}

		nextBlockStart += count
	}

	for {
		lastBlock := blocks[len(blocks)-1]
		firstFree := freeBlocks[0]
		if firstFree.start > lastBlock.end {
			break
		}

		blocks = blocks[:len(blocks)-1]
		freeBlocks = freeBlocks[1:]

		blockSize := (lastBlock.end - lastBlock.start)
		freeSize := (firstFree.end - firstFree.start)

		freeRemaining := max(0, freeSize-blockSize)
		blockRemaining := max(0, blockSize-freeSize)

		blocks = insertSorted(blocks, interval{
			start: firstFree.start,
			end:   firstFree.start + (blockSize - blockRemaining),
			value: lastBlock.value,
		})

		if freeRemaining > 0 {
			freeBlocks = insertSorted(freeBlocks, interval{
				start: firstFree.start + blockSize,
				end:   firstFree.end,
				value: FREE_BLOCK_ID,
			})
		}

		if blockRemaining > 0 {
			blocks = insertSorted(blocks, interval{
				start: lastBlock.start,
				end:   lastBlock.start + blockRemaining,
				value: lastBlock.value,
			})
		}
	}

	result := 0

	for _, b := range blocks {
		rangeSum := (b.end - b.start) * (b.start + b.end - 1) / 2
		result += b.value * rangeSum
	}

	return result
}

func SolvePart2(input string) int {
	diskMap := make([]int, 0)

	for _, c := range input {
		diskMap = append(diskMap, int(c-'0'))
	}

	blocks := make([]interval, 0)
	freeBlocks := make([]interval, 0)

	nextBlockStart := 0
	for i, count := range diskMap {
		isFreeSpace := i%2 != 0

		if isFreeSpace {
			freeBlocks = append(freeBlocks, interval{
				start: nextBlockStart,
				end:   nextBlockStart + count,
				value: FREE_BLOCK_ID,
			})
		} else {
			blocks = append(blocks, interval{
				start: nextBlockStart,
				end:   nextBlockStart + count,
				value: i / 2,
			})
		}

		nextBlockStart += count
	}

	blockMap := make(map[int]int) // Map of value to index in blocks
	for index, block := range blocks {
		blockMap[block.value] = index
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		var block interval
		var blockIndex int

		for j, b := range blocks {
			if b.value == i {
				block = b
				blockIndex = j
				break
			}
		}
		blockSize := (block.end - block.start)

		for freeBlockIndex, freeBlock := range freeBlocks {
			if freeBlock.start >= block.start {
				break
			}
			freeSize := (freeBlock.end - freeBlock.start)

			if freeSize >= blockSize {
				freeRemaining := max(0, freeSize-blockSize)
				if freeRemaining > 0 {
					freeBlocks = insertSorted(freeBlocks, interval{
						start: freeBlock.start + blockSize,
						end:   freeBlock.end,
						value: FREE_BLOCK_ID,
					})
				}
				freeBlocks = append(freeBlocks[:freeBlockIndex], freeBlocks[freeBlockIndex+1:]...)
				blocks = append(blocks[:blockIndex], blocks[blockIndex+1:]...)
				blocks = insertSorted(blocks, interval{
					start: freeBlock.start,
					end:   freeBlock.start + blockSize,
					value: block.value,
				})
				break
			}
		}
	}

	result := 0

	for _, b := range blocks {
		rangeSum := (b.end - b.start) * (b.start + b.end - 1) / 2
		result += b.value * rangeSum
	}

	return result
}
