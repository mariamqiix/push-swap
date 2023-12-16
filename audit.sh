#!/bin/bash
Continue() {
  echo "Press any key to continue"
  while [ true ]; do
    read -t 3 -n 1
    if [ $? = 0 ]; then
      clear
      return
    fi
  done
}
clear
echo "Try to run ./push-swap."
go run ./pushSwap/.
Continue

echo "Try to run ./push-swap '2 1 3 6 5 8' "
go run ./pushSwap/. "2 1 3 6 5 8"
echo -n "Instructions counter: " 
go run ./pushSwap/. "2 1 3 6 5 8" | wc -l
Continue

echo "Try to run ./push-swap '0 1 2 3 4 5' "
go run ./pushSwap/. "0 1 2 3 4 5"
Continue

echo "Try to run ./push-swap '0 one 2 3' "
go run ./pushSwap/. "0 one 2 3"
Continue

echo "Try to run ./push-swap '1 2 2 3' "
go run ./pushSwap/. "1 2 2 3"
Continue

echo "Try to run ./push-swap '<5 random numbers>' with 5 random numbers instead of the tag."
bash random.sh 5 0
Continue

echo "Try to run ./push-swap '<5 random numbers>' with 5 different random numbers instead of the tag."
bash random.sh 5 0
Continue

echo "Try to run ./Checker and input nothing."
go run ./checker/.
Continue

echo "Try to run ./Checker '0 one 2 3' "
go run ./checker/. "0 one 2 3"
Continue

echo "Try to run echo -e 'sa\npb\nrrr\n' | ./Checker '0 9 1 8 2 7 3 6 4 5'"
echo -e "sa\npb\nrrr\n"  | go run ./checker/. "0 9 1 8 2 7 3 6 4 5"
Continue

echo "Try to run echo -e 'pb\nra\npb\nra\nsa\nra\npa\npa\n' | ./Checker '0 9 1 8 2'"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | go run ./checker/. "0 9 1 8 2"
Continue

echo "Try to run ARG=('4 67 3 87 23'); ./push-swap ARG | ./Checker ARG"
ARG=('4 67 3 87 23'); go run ./pushSwap/. "$ARG" | go run ./checker/. "$ARG"
Continue

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG with 100 random different numbers instead of the tag."
Continue
bash random.sh 100 0
Continue

echo "Try to run ARG=('<100 random numbers>'); ./push-swap ARG | ./Checker ARG with 100 random different numbers instead of the tag."
bash random.sh 100 1
Continue