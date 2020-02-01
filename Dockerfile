FROM maxhedrom/photoprism:latest
# Set up project directory
WORKDIR "/go/src/github.com/mikepadge/photoprism"
COPY . .
