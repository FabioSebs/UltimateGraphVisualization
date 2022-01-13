import gzip
import shutil

with gzip.open('USA-road-d.COL.gr.gz', 'rb') as f:
    file = f.read()
    print(file)  # gives us Buffer

for x in file:
    print(x + '\n')
wad