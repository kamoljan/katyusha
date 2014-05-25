#!/bin/sh
REQUESTS=$1
CONCURRENCY=$2
URL=$3

echo 'Request ' $REQUESTS
echo 'Concurrency ' $CONCURRENCY
echo 'URL ' "$URL"

cmd="ab -e output.csv -n $REQUESTS -c $CONCURRENCY $URL"
echo $cmd
${cmd}
sed 1d output.csv > stats.csv
echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "get_'$REQUESTS'_'$CONCURRENCY'.jpg"; \
set datafile separator ","; \
set title "'$REQUESTS' requests with '$CONCURRENCY' concurrent connections"; \
set xlabel "Percentages served"; \
set ylabel "Response time (milliseconds)"; \
plot "stats.csv" notitle' | gnuplot -persist
