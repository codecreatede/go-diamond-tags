package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-10-3

A diamond reads to protein aligner analyzer, which will take the read to protein alignments
and will estimate the coverage of the alignment with respect to the hsp aligned and then if you
want you can extract the regions using the gomapper dimaond

*/

import (
	"bufio"
	"fmt"
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

func sum(arr []float64) float64 {
	counter := float64(0)
	for i := range arr {
		counter += arr[i]
	}
	return counter
}

func pacbio() ([]string, []string, []float64) {
	readOpen, err := os.Open(referencefasta)
	if err != nil {
		log.Fatal(err)
	}

	readbuffer := bufio.NewScanner(readOpen)
	header := []string{}
	sequences := []string{}
	length := []float64{}

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
		length = append(length, float64(len(sequences[i])))
	}
	return header, sequences, length

}

func hspFunc(cmd *cobra.Command, args []string) {
	refID := []string{}
	alignID := []string{}
	refIdenStart := []float64{}
	refIdenEnd := []float64{}
	alignIdenStart := []float64{}
	alignIdenEnd := []float64{}
	fOpen, err := os.Open(alignmentfile)
	if err != nil {
		log.Fatal(err)
	}

	fRead := bufio.NewScanner(fOpen)

	for fRead.Scan() {
		line := fRead.Text()
		refID = append(refID, strings.Split(string(line), "\t")[0])
		alignID = append(alignID, strings.Split(string(line), "\t")[1])
		start1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[6], 32)
		end1, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[7], 32)
		start2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[8], 32)
		end2, _ := strconv.ParseFloat(strings.Split(string(line), "\t")[9], 32)
		refIdenStart = append(refIdenStart, start1)
		refIdenEnd = append(refIdenEnd, end1)
		alignIdenStart = append(alignIdenStart, start2)
		alignIdenEnd = append(alignIdenEnd, end2)
	}
	id, _, length := pacbio()

	type cov struct {
		id  string
		cov float64
	}

	coverageSeq := []cov{}
	for i := range id {
		for j := range refID {
			if id[i] == refID[j] {
				coverageSeq = append(coverageSeq, cov{
					id:  refID[j],
					cov: (refIdenEnd[j] - refIdenStart[j]) / length[i] * 100,
				})
			}
		}
	}

	for i := range coverageSeq {
		fmt.Println(coverageSeq[i].id, coverageSeq[i].cov)
	}

	file, err := os.Create("coveragestimation.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i := range coverageSeq {
		storeI := strconv.FormatFloat(coverageSeq[i].cov, 'f', -1, 64)
		_, err := file.WriteString(coverageSeq[i].id + "\t" + storeI + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
