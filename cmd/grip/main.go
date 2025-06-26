package main

import (
	"debug/elf"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <path_to_elf_file>")
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	f, err := elf.NewFile(file)
	if err != nil {
		log.Fatalf("Error reading ELF file: %v", err)
	}

	fmt.Printf("Entry point: 0x%x\n", f.FileHeader.Entry)
	fmt.Println("Sections:")
	for _, section := range f.Sections {
		fmt.Printf("  Name: %s, Type: %s, Address: 0x%x, Size: %d bytes\n",
			section.Name, section.Type, section.Addr, section.Size)
	}

	fmt.Println("Program Headers:")
	for _, prog := range f.Progs {
		fmt.Printf("  Type: %s, Offset: 0x%x, Virtual Address: 0x%x, Physical Address: 0x%x, File Size: %d bytes, Memory Size: %d bytes\n",
			prog.Type, prog.Off, prog.Vaddr, prog.Paddr, prog.Filesz, prog.Memsz)
	}
}
