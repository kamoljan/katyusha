#!/bin/sh
CONCURRENCY=$1

echo 'Concurrency ' $CONCURRENCY

# cmd="./katyusha $CONCURRENCY > stats.csv"
# echo $cmd
echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "'$CONCURRENCY'"; \
set datafile separator ","; \
set title  "'$CONCURRENCY' concurrent connections"; \
set xlabel "Percentages served"; \
set ylabel "Response time (milliseconds)"; \
plot "stats.csv" notitle' | gnuplot -persist
