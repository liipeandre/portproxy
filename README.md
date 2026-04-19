# PortProxy

## Descrição

O PortProxy é uma ferramenta de rede desenvolvida em Go. Ele funciona de forma análoga ao comando netsh interface portproxy do Windows, mas com um diferencial crucial: suporte nativo a tráfego UDP, além do TCP, usando o recurso de sockets.

É a solução ideal para administradores que precisam de um redirecionamento de portas leve, sem as limitações de protocolo do utilitário nativo do sistema operacional.

## Por que usar o PortProxy?
- Além do TCP: Ao contrário do netsh, este script lida perfeitamente com pacotes UDP (essencial para serviços como SNMP, DNS e protocolos de impressão).

- Multi-Portas: Configure múltiplos túneis (TCP e UDP) em uma única linha de comando.

- Performance em Go: Aproveita o modelo de concorrência do Go para garantir que o redirecionamento não seja um gargalo na rede.

- Sem Dependências: O binário compilado contém tudo o que é necessário para rodar, facilitando a distribuição em diferentes máquinas.


## Como Compilar

go build -o ./bin/portproxy.exe ./src


## Como Usar

A sintaxe é direta. Informe o IP de destino e a lista de portas no formato abaixo:

protocolo/porta_origem:porta_destino

Exemplo: 
portproxy.exe 192.168.137.2 tcp/80:80 tcp/443:443 tcp/9100:9100 udp/161:161

