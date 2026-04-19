# Para compilar, abra o terminal e execute o comando abaixo.
go build -o ./bin/portproxy.exe ./src

# Para usar
portproxy.exe 192.168.137.2 tcp/80:80 tcp/443:443 tcp/9100:9100 udp/161:161