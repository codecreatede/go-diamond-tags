# go-diamond-tags

- a subfunction for the diamond-aligner and estimating the hsp coverage
- this is also integratd into the go-mapper-diamond
- you can use this as a separate if you have already aligned reads to the protein.
- when you have a metagenome and you want to generate the annotation tags using the bacterial genome proteins.
- when you have MAGS and you want to generate the specific mags to protein alignment.  
- same goes for other species also. 
- your pacbio reads should be a linear fasta, which is the way the pacbio reads usually comes.
- Incase of the other fasta sequences such as genomes or others, remember to run the awk utility to linearize it. This is faster than implementing a loop iteration. 

```
awk '/^>/ {printf("\n%s\n",$0);next; } { printf("%s",$0);}  \
                         END {printf("\n");}' inputfasta > output.fasta
```

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

gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-tags ±main⚡ » \
go run main.go alignment -a matches.tsv -p ./samplefiles/pacbioreads.fasta
>chr10:66478458-66505490 1.0061776347427218
>chr10:66478458-66505490 0.9026005252839123
>chr11:66478458-66505490 0.8952021603225687
>chr11:66478458-66505490 1.0875596493175008
gauavsablok@gauravsablok ~/Desktop/codecreatede/golang/go-diamond-tags ±main⚡ » \
cat coveragestimation.txt
>chr10:66478458-66505490        1.0061776347427218
>chr10:66478458-66505490        0.9026005252839123
>chr11:66478458-66505490        0.8952021603225687
>chr11:66478458-66505490        1.0875596493175008

```
Gaurav Sablok
