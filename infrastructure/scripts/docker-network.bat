@echo off
setlocal enabledelayedexpansion

:: Verify if network already exists
set "OUT_IP_ADDR="
for /f "delims=" %%i in ('ipconfig ^| findstr /C:"204.7.9.0"') do (
    set "OUT_IP_ADDR=%%i"
)

if not "!OUT_IP_ADDR!"=="" (
    echo Network already initialized, starting...
) else (
    docker network create --subnet 204.7.9.0/24 hermes_foods_net_dev
)

endlocal
