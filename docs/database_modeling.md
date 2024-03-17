# Introduction
This file is about the database model chosen based on business logic provided by postech.

<img src='./img/hf-db-model.png" width="700px" height="500px">

## Banco de dados SQL Amazon RDS com Postgres

Resolvemos migrar o nosso banco de dados Postgres que estava rodando em Conteiner, usado para persistir os dados das nossas APIs (Cliente, Produto, Pedido e Voucher) para Amazon RDS, pelas seguintes vantagens:

1. Escalabilidade e desempenho: Com o Amazon RDS nós podemos escalar o nosso banco de dados verticalmente e horizontalmente, conforme necessário, facilitando futuramente o gerenciamento de desempenho da nossa aplicação a medida que a nossa base de usuários/ produtos  crescem. 

2. Disponibilidade e durabilidade dos dados: Com o amazon RDS é possível configurarmos o nosso banco de dados afim de garantir a alta disponibilidade dos nossos dados, através de backups automáticos, replicação em múltiplas zonas de disponibilidades e recuperar os nossos dados em casos de desatres.

3. Gerenciamento simplificado: Por se tratar de uma plataforma como serviço (PaaS), não precisamos nos preocupar com atualizações dos patches e sistema operacional, e acesso simplificado pelo console AWS e CLI, para configurar, monitorar e ajustar nosso banco de dados.

4. Segurança: O Amazon RDS oferece recursos de segurança robustos, como criptografia de dados em repouso e em trânsito, além de integração com outros serviços da AWS para controle de acesso e conformidade.

5. Custos: Chegamos a conclusão que caso a nossa aplicação venha a crescer, gerencia-lá pelo amazon RDS nos geraria uma economia a longo prazo, do que trabalhar dentro de um ambiente de infra-estrutura como serviço (IaaS), pois se trata de um ambiente elástico e escalável.


## Banco de dados noSQL - Amazon Dynamo DB

Resolvemos criar uma tabela no banco de dados noSQL Dynamo DB para trabahar em conjunto com o cognito no nosso sistema de autenticação, pelas seguintes vantagens:

1. Escalabilidade e desempenho: Dynamo DB é um banco noSQL altamente escalável e pode lidar com cargas de trabalho de formas mais variadas e imprevisíveis, trabalha com grande número de usuários ou picos de tráfego e tem capacidade de escalar automaticamente.

2. Gerenciamento simplificado: É um serviço totalmente gerenciado pela AWS e por esse motivo não precisamos nos preocupar com a infraestrutura subjacente, a AWS cuida da configuração, monitoramento e manutenção do banco de dados.

3. Integração com o Amazon Cognito: Uma das justificativas da nossa escolha é a fácil integração entre Dynamo DB, serviço de autenticação do cognito e as funções lambda. Usamos o cognito para gerenciar usuários e autenticá-los, enquanto o Dynamo DB foi usado para armazenar dados relacionados a autenticação.


