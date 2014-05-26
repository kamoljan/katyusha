#!/bin/sh
CONCURRENCY=$1

echo 'Concurrency ' $CONCURRENCY

# cmd="./katyusha $CONCURRENCY > stats.csv"
# echo $cmd
# set title  "'$CONCURRENCY' concurrent connections"; \
echo '\
set terminal png font "/Library/Fonts/Times\ New\ Roman.ttf" 14; \
set output "'$CONCURRENCY'.jpg"; \
set datafile separator ","; \
set xlabel "Concurrent connections"; \
set ylabel "Response time (milliseconds)"; \
plot "'$CONCURRENCY'" notitle' | gnuplot -persist
