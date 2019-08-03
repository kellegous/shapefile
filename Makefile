ALL: bin/poc bin/generate_no_m_tests bin/readshp bin/readdbf

bin/poc: $(shell find . -type f -name '*.go')
	go build -o $@ github.com/kellegous/shapefile/cmd/poc

bin/generate_no_m_tests: $(shell find . -type f -name '*.go')
	go build -o $@ github.com/kellegous/shapefile/cmd/generate_no_m_tests

bin/readshp: $(shell find cmd/readshp shp -type f -name '*.go')
	go build -o $@ github.com/kellegous/shapefile/cmd/readshp

bin/readdbf: $(shell find cmd/readdbf dbf -type f -name '*.go')
	go build -o $@ github.com/kellegous/shapefile/cmd/readdbf

test:
	go test github.com/kellegous/shapefile/shp github.com/kellegous/shapefile/dbf

clean:
	rm -rf bin