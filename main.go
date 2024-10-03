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
	alignmentfile string
	pacbioreads   string
)

var rootCmd = &cobra.Command{
	Use:  "alignment",
	Long: "Analyzes the hsp from the diamond read to protein alignment",
	Run:  hspFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&alignmentfile, "alignmentfile", "a", "alignment file to be analyzed", "alignment")
	rootCmd.Flags().StringVarP(&pacbioreads, "pacbioreads", "p", "pacbio reads file", "pacbio file")
}

func hspFunc(cmd *cobra.Command, args []string) {
	refID := []string{}
	alignID := []string{}
	refIdenStart := []int{}
	refIdenEnd := []int{}
	alignIdenStart := []int{}
	alignIdenEnd := []int{}
	// calDiff := []int32{}
	// map as a set for making the unique ids for the fasta for the calculation.
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
	}
	for fRead.Scan() {
		line := fRead.Text()
		start1, _ := strconv.Atoi(strings.Split(string(line), "\t")[6])
		end1, _ := strconv.Atoi(strings.Split(string(line), "\t")[7])
		start2, _ := strconv.Atoi(strings.Split(string(line), "\t")[8])
		end2, _ := strconv.Atoi(strings.Split(string(line), "\t")[9])
		refIdenStart = append(refIdenStart, start1)
		refIdenEnd = append(refIdenEnd, end1)
		alignIdenStart = append(alignIdenStart, start2)
		alignIdenEnd = append(alignIdenEnd, end2)
	}
	// storing the unique AlignIDs as a map and then using the same for the calculation.
	// Not interested in the e-value, so the only role of the map is to make the unique set of
	// the reference alignment.
	for fRead.Scan() {
		line := fRead.Text()
		uniqueAlign[strings.Split(string(line), "\t")[0]] = strings.Split(string(line), "\t")[10]
	}

	readOpen, err := os.Open(pacbioreads)
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
		if string(line[0]) == "@" {
			header = append(header, line)
		}
	}
	for i := range sequences {
		length = append(length, len(sequences[i]))
	}
}
