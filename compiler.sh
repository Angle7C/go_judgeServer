Content="$1"
Judge="$2"
Command="$3"
Work="$4"
RunId="$5"
Type="$6"
#创建编译文件夹

Path="$Work/$RunId/compiler"
mkdir -p "$Path"
touch "$Path/compiler.sh"
touch "$Path/compiler.in"
touch "$Path/compiler.err"
touch "$Path/compiler.out"
touch "$Path/compiler.log"
touch "$Path/Main.$Type"
#创建成功，将代码写入特定的 文件中
echo "$Content" >> "$Path/Main.$Type"
#写入成功，将编译指令写入compiler.sh,并给予可执行权限
echo "$Command" >> "$Path/compiler.sh"
chmod 777 "$Path/compiler.sh"
# 使用judge执行comipler.sh
"$Judge" --exe="$Path/compiler.sh" --cpu_max=20 --cpu_real=30 --input="$Path/compiler.in" --out="$Path/compiler.out" --error="$Path/compiler.err" --log="$Path/compiler.log"



