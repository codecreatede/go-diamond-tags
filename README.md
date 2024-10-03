# go-diamond-tags

- a subfunction for the diamond-aligner and extracting the hsp
- this is also integratd into the go-mapper-diamond
- you can use this as a separate if you have already aligned reads to the protein.
- when you have a metagenome and you want to generate the annotation tags using the bacterial genome proteins.
- when you have MAGS and you want to generate the specific mags to protein alignment.  
- same goes for other species also. 

```
gauravsablok@gaurav-sablok ~/Desktop/codecreatede/golang/go-diamond-hsp-extract ±main⚡ » go run main.go -h
Analyzes the hsp from the diamond read to protein alignment

Usage:
  alignment [flags]

Flags:
  -a, --alignmentfile string   alignment (default "alignment file to be analyzed")
  -h, --help                   help for alignment
  -p, --pacbioreads string     pacbio file (default "pacbio reads file")
```
Gaurav Sablok
