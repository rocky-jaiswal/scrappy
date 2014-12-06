source gvp
gpm
mkdir -p ./.godeps/src/github.com/rocky-jaiswal/scrappy/scraper
cp ./scraper/scraper.go ./.godeps/src/github.com/rocky-jaiswal/scrappy/scraper/scraper.go
go build && ./scrappy
