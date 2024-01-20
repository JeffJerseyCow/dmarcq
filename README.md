# dmarcq - DMARC Query
dmarcq is a quickly written Golang tool for querying DMARC records from domains provided via stdin. It creates a csv output which is far simpler to parse for this type of exercise.

## Build
Use a combination of git and Golang to clon and build.

```bash
git clone github.com/JeffJerseyCow/dmarcq
cd dmarcq
go build
```

## Usage
The simplest way to user **dmarcq** is by running a bash command to read a file and echoing input to **dmarcq**. To print out the data to stdout directly, simply execute:

```bash
while read domain; do echo $domain | ./dmarc; done < domains/domains.list
```

Or if you'd like to write it to a csv file:

```bash
while read domain; do echo $domain | ./dmarc; done < domains/domains.list > domains/domains.dmarc.csv
```
