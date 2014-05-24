#!/bin/sh
CONCURRENCY=$1

echo 'Concurrency ' $CONCURRENCY

# cmd="./katyusha $CONCURRENCY > stats.csv"
# echo $cmd
# set title  "'$CONCURRENCY' concurrent connections"; \
echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "'$CONCURRENCY'"; \
set datafile separator ","; \
set xlabel "Concurrent connections"; \
set ylabel "Response time (milliseconds)"; \
plot "stats'$CONCURRENCY'.csv" notitle' | gnuplot -persist
