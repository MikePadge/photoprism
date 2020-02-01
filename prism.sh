docker run -d \
  --name photoprism \
  -p 2342:2342 \
  -v /Volumes/blazer/photoprism/Pictures:/home/photoprism/Pictures/Originals \
  -v /Volumes/blazer/photoprism/cache:/home/photoprism/.cache/photoprism \
  -v /Volumes/blazer/photoprism/db:/home/photoprism/.local/share/photoprism/resources/database \
  maxhedrom/photoprism:latest 

