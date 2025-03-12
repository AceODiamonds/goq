package go_qasm

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func Qasm_reader(file_path string) {
	content, err := os.ReadFile(file_path)
	if er != nil {
		log.Fatal("Error reading the file: ", err)
	}
	//now need to parse the conent line by line
	lines := strings.Split(string(content), "\n")
	qubitRegex := regexp.MustCompile(`\[(\d+)\]`)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		//Skip header and include statements
		if strings.HasPrefix(line, "OPENQASM") || strings.HasPrefix(line, "include") {
			continue
		}
		//Split into gates and arguments
		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			continue
		}

		gate := strings.TrimSpace(parts[0])
		args := strings.TrimSuffix(strings.TrimSpace(parts[1]), ";")
		var qubits []string
		for _, arg := range strings.Split(args, ",") {
			arg = strings.TrimSpace(arg)
			match := qubitRegex.FindStringSubmatch(arg)
			if len(match) > 1 {
				qubits = append(qubits, match[1])
			}
		}

		if len(qubits) > 0 {
			fmt.Printf("%s %s\n", gate, strings.Join(qubits, ","))
		}

	}
}
