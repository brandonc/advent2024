package day09

import (
	"io"

	"github.com/brandonc/advent2024/internal/ui"
	"github.com/brandonc/advent2024/solutions/solution"
)

type day09 struct {
	root *block
	tail *block
	disk []byte
}

func Factory() solution.Solver {
	return day09{}
}

var (
	BlockTypeFile = byte('F')
	BlockTypeFree = byte('.')
)

type block struct {
	BlockType byte
	ID        int
	Size      int
	Next      *block
	Prev      *block
}

func (d *day09) defragFile() {
	tail := d.tail

	for {
		head := d.root.Next

		// Find next file block using tail pointer
		for {
			if tail == nil || tail.BlockType == BlockTypeFile {
				break
			}
			tail = tail.Prev
		}

		if tail == nil {
			break
		}

		// Work forwards to find a suitable free block
		for {
			if head == tail {
				break
			}

			if head.BlockType == BlockTypeFree && head.Size >= tail.Size {
				break
			}

			head = head.Next
		}

		if head == tail || head == nil {
			tail = tail.Prev
			continue
		}

		// Remove tail block and insert it before the free block
		head.Size -= tail.Size

		insertAfter := head.Prev

	outer:
		for {
			// Insert after the rightmost file block
			if insertAfter.BlockType == BlockTypeFile {
				for {
					if insertAfter.Next != nil && insertAfter.Next.BlockType == BlockTypeFile {
						insertAfter = insertAfter.Next
					} else {
						break outer
					}
				}
			}
			insertAfter = insertAfter.Prev
		}

		newBlock := block{ID: tail.ID, Size: tail.Size, BlockType: BlockTypeFile, Prev: insertAfter, Next: insertAfter.Next}
		insertAfter.Next = &newBlock

		tail.BlockType = BlockTypeFree
		tail.ID = 0

		tail = tail.Prev
	}
}

func (d *day09) defragBlock() {
	free := 0
	z := len(d.disk)

	for {
		z--
		if d.disk[z] == '.' {
			continue
		}

		for {
			if d.disk[free] == '.' {
				break
			}
			free++
		}

		if free >= z {
			break
		}

		// free is at a free space, z is at a block
		ui.Debugf("Moving %d to %d", z, free)
		d.disk[free], d.disk[z] = d.disk[z], d.disk[free]
		free = 0
	}
}

func (d *day09) decode(encoded []byte) {
	d.root = &block{}
	d.disk = make([]byte, 0, 20_000)

	current := d.root
	var previous *block
	position := 0

	for i, b := range encoded {
		newBlock := &block{}
		if i%2 == 0 {
			newBlock.BlockType = BlockTypeFile
			newBlock.ID = position

			for i := 0; i < int(b-'0'); i++ {
				d.disk = append(d.disk, byte(position))
			}
		} else {
			newBlock.BlockType = BlockTypeFree
			position += 1

			for i := 0; i < int(b-'0'); i++ {
				d.disk = append(d.disk, '.')
			}
		}

		newBlock.Size = int(b - '0')
		newBlock.Prev = previous
		d.tail = newBlock

		current.Next = newBlock

		current = newBlock
		previous = current
	}
}

func (d *day09) checksumDisk() int {
	checksum := 0
	for i := 0; i < len(d.disk); i++ {
		if d.disk[i] != '.' {
			checksum += i * int(d.disk[i])
		}
	}
	return checksum
}

func (d *day09) checksumBlocks() int {
	current := d.root.Next

	checksum := 0
	position := 0
	for current != nil {
		if current.BlockType == BlockTypeFile {
			for pos := position; pos < position+current.Size; pos++ {
				ui.Debugf("Adding %d at position %d", pos*current.ID, pos)
				checksum += pos * current.ID
			}
		}
		position += current.Size
		current = current.Next
	}
	return checksum
}

func (d day09) Part1(reader io.Reader) int {
	encoded, err := io.ReadAll(reader)
	if err != nil {
		ui.Die(err)
	}

	d.decode(encoded)
	d.defragBlock()

	return d.checksumDisk()
}

func (d day09) Part2(reader io.Reader) int {
	encoded, err := io.ReadAll(reader)
	if err != nil {
		ui.Die(err)
	}

	d.decode(encoded)
	d.defragFile()

	return d.checksumBlocks()
}
