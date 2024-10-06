# go-diamond-tags

- a subfunction for the diamond-aligner and extracting the hsp
- this is also integratd into the go-mapper-diamond
- you can use this as a separate if you have already aligned reads to the protein.
- when you have a metagenome and you want to generate the annotation tags using the bacterial genome proteins.
- when you have MAGS and you want to generate the specific mags to protein alignment.  
- same goes for other species also. 

```
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-tags ±main⚡ » \
go run main.go -h
Analyzer for the diamond aligner and pacbio reads for hints

Usage:
  analyze [command]

Available Commands:
  alignment
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for analyze

Use "analyze [command] --help" for more information about a command.
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-tags ±main⚡ » \
go run main.go alignment -h
Analyzes the hsp from the diamond read to protein alignment

Usage:
  analyze alignment [flags]

Flags:
  -a, --alignmentfile string   alignment (default "alignment file to be analyzed")
  -h, --help                   help for alignment
  -p, --pacbioreads string     pacbio file (default "pacbio reads file")

```
Gaurav Sablok
