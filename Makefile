ALL: bin/poc bin/generate_no_m_tests

bin/poc: $(shell find . -type f -name '*.go')
	go build -o $@ github.com/kellegous/shp/cmd/poc

bin/generate_no_m_tests: $(shell find . -type f -name '*.go')
	go build -o $@ github.com/kellegous/shp/cmd/generate_no_m_tests

test:
	go test github.com/kellegous/shapefile/shp

clean:
	rm -rf bin