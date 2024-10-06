package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-3

A diamond reads to protein aligner analyzer, which will take the read to protein alignments
and will extract the corresponding hsps from the sequences. This can be used to extract all the
aligned hsps or you can use to generate the coverage profile to see how much portion of each hsp
is covered and implemented.


*/

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	alignmentfile  string
	referencefasta string
)

var rootCmd = &cobra.Command{
	Use:  "analyze alignments",
	Long: "Analyzer for the diamond aligner and pacbio reads for hints",
}

var alignmentCmd = &cobra.Command{
	Use:  "alignment",
	Long: "Analyzes the hsp from the diamond read to protein alignment",
	Run:  hspFunc,
}

func init() {
	alignmentCmd.Flags().
		StringVarP(&alignmentfile, "alignmentfile", "a", "alignment file to be analyzed", "alignment")
	alignmentCmd.Flags().
		StringVarP(&referencefasta, "referencefasta", "p", "pacbio reads file", "pacbio file")

	rootCmd.AddCommand(alignmentCmd)
}

func sum(arr []int) int {
	count := 0
	for i := range arr {
		count += arr[i]
	}
	return count
}

func pacbio() []int {
	readOpen, err := os.Open(referencefasta)
	if err != nil {
		log.Fatal(err)
	}

	readbuffer := bufio.NewScanner(readOpen)
	header := []string{}
	sequences := []string{}
	length := []int{}

	for readbuffer.Scan() {
		line := readbuffer.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" ||
			string(line[0]) == "C" {
			sequences = append(sequences, line)
		}
		if string(line[0]) == ">" {
			header = append(header, line)
		}
	}
	for i := range sequences {
		length = append(length, len(sequences[i]))
	}
	return length
}

func hspFunc(cmd *cobra.Command, args []string) {
	refID := []string{}
	alignID := []string{}
	refIdenStart := []int{}
	refIdenEnd := []int{}
	alignIdenStart := []int{}
	alignIdenEnd := []int{}
	uniqueAlign := make(map[string]string)
	fOpen, err := os.Open(alignmentfile)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)

	for fRead.Scan() {
		line := fRead.Text()
		refID = append(refID, strings.Split(string(line), "\t")[0])
		alignID = append(alignID, strings.Split(string(line), "\t")[1])
		start1, _ := strconv.Atoi(strings.Split(string(line), "\t")[6])
		end1, _ := strconv.Atoi(strings.Split(string(line), "\t")[7])
		start2, _ := strconv.Atoi(strings.Split(string(line), "\t")[8])
		end2, _ := strconv.Atoi(strings.Split(string(line), "\t")[9])
		refIdenStart = append(refIdenStart, start1)
		refIdenEnd = append(refIdenEnd, end1)
		alignIdenStart = append(alignIdenStart, start2)
		alignIdenEnd = append(alignIdenEnd, end2)
		uniqueAlign[strings.Split(string(line), "\t")[0]] = strings.Split(string(line), "\t")[10]
	}

	//last part left rest all bugs fixed.
	/*
		calID := []string{}
		calDiff := []int{}

		for i := range refID {
			calID = append(calID, refID[i])
			calDiff = append(calDiff, refIdenEnd[i]-refIdenStart[i])
		}

		calIDcov := []string{}
		calCov := []int{}
		intermediateCov := []int	// moved the pacbio function to outside and calling the varibale as shadowing.
		length := pacbio()
		var intermediateCovSum int
		for i := range calID {
			for j := range length {
				for k := range uniqueAlign {
					if uniqueAlign[k] == calID[i] {
						intermediateCov = append(intermediateCov, calCov[i])
						intermediateCovSum = sum(intermediateCov)
						calIDcov = append(calIDcov, calID[i])
						calCov = append(calCov, intermediateCovSum/length[j])
					}
				}
			}
		}
		   file, err := os.Create("coverage estimation")

		   	if err != nil {
		   		log.Fatal(err)
		   	}

		   defer file.Close()
		   /*

		   	for i := range calIDcov {
		   		_, err := file.WriteString(calIDcov[i))
		   		if err != nil {
		   			log.Fatal(err)
		   		}
		   	}
	*/
}
