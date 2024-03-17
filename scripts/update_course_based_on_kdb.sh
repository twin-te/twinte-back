#!/bin/bash

current_year=$(date +%Y)
current_month=$(date +%m)

if [[ "$current_month" -ge "04" ]]; then
    academic_year=$current_year
else
    academic_year=$((current_year - 1))
fi

kdb_json_file_path="$PWD/kdb.json"

python twinte-parser/download_and_parse.py --year $academic_year --output-path $kdb_json_file_path 
go run . update-courses-based-on-kdb --year $academic_year --kdb-json-file-path $kdb_json_file_path 
