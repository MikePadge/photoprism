<<<<<<< HEAD
FROM maxhedrom/photoprism:latest
=======
FROM photoprism/development:20200203

>>>>>>> 5fba03844298ab501ce513a3f967b7578bc09707
# Set up project directory
WORKDIR "/go/src/github.com/mikepadge/photoprism"
COPY . .
