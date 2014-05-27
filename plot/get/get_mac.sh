#!/bin/sh
REQUESTS=$1
CONCURRENCY=$2

echo 'Request ' $REQUESTS
echo 'Concurrency ' $CONCURRENCY
# set title "'$REQUESTS' requests with '$CONCURRENCY' concurrent connections"; \
echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "'$REQUESTS'_'$CONCURRENCY'.jpg"; \
set datafile separator ","; \
set xlabel "Percentages served"; \
set ylabel "Response time (milliseconds)"; \
plot "get_'$REQUESTS'_'$CONCURRENCY'" notitle' | gnuplot -persist