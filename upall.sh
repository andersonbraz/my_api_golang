#!/bin/zsh
clear
git add .
git commit -m "Up Files"
git push origin master


# DESAFIO
#
# Adapte seu código-fonte para que as credenciais do MongoDB sejam configuradas por variáveis 
# de ambiente do Docker ou argumentos da linha de comando.
#
# Utilize as bibliotecas cobra e viper para validar se as variáveis de ambiente 
# estão preenchidas antes de iniciar o serviço da web escrito em Go.
#
# Envie-me uma captura de tela do seu Dockerfile, docker-compose.yml e o código-fonte da resolução de configuração.
#
# DICA: ENV, CMD