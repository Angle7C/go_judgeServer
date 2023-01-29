File_A="$1"
File_B="$2"

if [ ! -f "$File_A" ]; then
  echo 0
  exit
fi

if [ ! -f "$File_A" ]; then
  echo 0
  exit
fi

diff "$File_A" "$File_B"
echo "$0"
